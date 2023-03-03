package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/genproto/googleapis/api/serviceconfig"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/pluginpb"
)

type rule struct {
	method string
	path   string
	body   bool
}

type mappings map[string]rule

func main() {
	var flags flag.FlagSet
	grpcAPIConfig := flags.String("grpc_api_configuration", "", "path to GRPC API configuration")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		if *grpcAPIConfig == "" {
			fmt.Fprintln(os.Stderr, "path to grpc API configuration required")
			os.Exit(1)
		}

		data, err := os.ReadFile(*grpcAPIConfig)
		if err != nil {
			panic(err)
		}

		json, err := yaml.YAMLToJSON(data)
		if err != nil {
			panic(err)
		}

		var config serviceconfig.Service
		if err := (protojson.UnmarshalOptions{DiscardUnknown: true}).Unmarshal(json, &config); err != nil {
			panic(err)
		}

		m := mappings{}
		for _, r := range config.Http.Rules {
			rule := rule{
				body: r.Body == "*",
			}

			switch r.Pattern.(type) {
			case *annotations.HttpRule_Get:
				rule.method = "MethodGet"
				rule.path = r.GetGet()
			case *annotations.HttpRule_Post:
				rule.method = "MethodPost"
				rule.path = r.GetPost()
			case *annotations.HttpRule_Put:
				rule.method = "MethodPut"
				rule.path = r.GetPut()
			case *annotations.HttpRule_Delete:
				rule.method = "MethodDelete"
				rule.path = r.GetDelete()
			default:
				fmt.Fprintf(os.Stderr, "unsupported pattern: %T\n", r.Pattern)
			}

			m[r.Selector] = rule
		}
		// We have some use of the optional feature in our proto3 definitions.
		// This broadcasts that our plugin supports it and hides the generated
		// warning.
		gen.SupportedFeatures |= uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(m, gen, f)
		}
		return nil
	})
}

// generateFile generates a _ascii.pb.go file containing gRPC service definitions.
func generateFile(m mappings, gen *protogen.Plugin, file *protogen.File) {
	filename := file.GeneratedFilenamePrefix + "_client.pb.gw.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-gateway-client. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	ident := func(pkg string) func(name string) string {
		return func(name string) string {
			return g.QualifiedGoIdent(protogen.GoIdent{
				GoImportPath: protogen.GoImportPath(pkg),
				GoName:       name,
			})
		}
	}

	pkgfmt := ident("fmt")
	bytes := ident("bytes")
	context := ident("context")
	io := ident("io")

	netURL := ident("net/url")
	netHTTP := ident("net/http")

	protojson := ident("google.golang.org/protobuf/encoding/protojson")
	pbstatus := ident("google.golang.org/genproto/googleapis/rpc/status")
	status := ident("google.golang.org/grpc/status")

	for _, srv := range file.Services {
		typ := srv.GoName + "HTTPClient"
		g.P("type ", typ, " struct {")
		g.P("client *", netHTTP("Client"))
		g.P("addr string")
		g.P("}\n")

		g.P("func New", typ, "(addr string) (*", typ, ") {")
		g.P("return &", typ, "{ client: ", netHTTP("DefaultClient"), ", addr: addr }")
		g.P("}\n")

		for _, method := range srv.Methods {
			rule, ok := m[string(method.Desc.FullName())]
			if !ok {
				continue
			}

			g.P("func (x *", typ, ") ", method.GoName, "(ctx ", context("Context"), ", v *", method.Input.GoIdent, ") (*", method.Output.GoIdent, ", error) {")

			g.P("var body ", io("Reader"))
			path, inPath := renderPath(pkgfmt, rule, method.Input)
			if rule.body {
				g.P("var values ", netURL("Values"))
				g.P("reqData, err := ", protojson("Marshal"), "(v)")
				g.P("if err != nil { return nil, err }")
				g.P("body = ", bytes("NewReader"), "(reqData)")
			} else {
				var setValues []string
				for _, field := range method.Input.Fields {
					if _, ok := inPath[field]; !ok {
						var val string
						if field.Desc.Kind() == protoreflect.StringKind {
							val = "v." + field.GoName
						} else {
							val = fmt.Sprintf(`%s("%%v", v.%s)`, pkgfmt("Sprintf"), field.GoName)
						}

						setValues = append(setValues, fmt.Sprintf(`values.Set("%s", %s)`, field.Desc.JSONName(), val))
					}
				}

				// only allocate if we have any values to set on the query
				if len(setValues) == 0 {
					g.P("var values ", netURL("Values"))
				} else {
					g.P("values := ", netURL("Values"), "{}")
					for _, val := range setValues {
						g.P(val)
					}
				}
			}

			g.P("req, err := ", netHTTP("NewRequestWithContext"), "(ctx, ", netHTTP(rule.method), ", x.addr+", path, ", body)")
			g.P("if err != nil { return nil, err }")

			g.P("req.URL.RawQuery = values.Encode()")

			g.P("resp, err := x.client.Do(req)")
			g.P("if err != nil { return nil, err }")
			g.P("defer resp.Body.Close()")

			g.P("var output ", method.Output.GoIdent)
			g.P("respData, err := ", io("ReadAll"), "(resp.Body)")
			g.P("if err != nil { return nil, err }")

			g.P("if err := checkResponse(resp, respData); err != nil {return nil, err}")

			g.P("if err := ", protojson("Unmarshal"), "(respData, &output); err != nil { return nil, err }")
			g.P("return &output, nil")
			g.P("}\n")
		}
	}

	g.P("func checkResponse(resp *", netHTTP("Response"), ", v []byte) error {")
	g.P("if resp.StatusCode != ", netHTTP("StatusOK"), "{")
	g.P("var status ", pbstatus("Status"))
	g.P("if err := ", protojson("Unmarshal"), "(v, &status); err != nil { return err }")
	g.P("return ", status("ErrorProto"), "(&status)")
	g.P("}\n")
	g.P("return nil")
	g.P("}\n")
}

func renderPath(pkgfmt func(string) string, rule rule, msg *protogen.Message) (string, map[*protogen.Field]struct{}) {
	var (
		args   []string
		inPath = map[*protogen.Field]struct{}{}
	)

	v := `"/`
	parts := strings.Split(rule.path, "/")
	for i, p := range parts {
		if p == "" {
			continue
		}

		if p[0] == '{' && p[len(p)-1] == '}' {
			for _, field := range msg.Fields {
				if p[1:len(p)-1] == field.Desc.JSONName() {
					p = "%v"
					args = append(args, "v."+field.GoName)
					inPath[field] = struct{}{}
				}
			}
		}
		v += p

		if i < len(parts)-1 {
			v += "/"
		}
	}

	if len(args) > 0 {
		return fmt.Sprintf("%s(%s, %s)", pkgfmt("Sprintf"), v+`"`, strings.Join(args, ",")), inPath
	}

	return v + `"`, inPath
}
