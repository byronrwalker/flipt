// Code generated by protoc-gen-go-gateway-client. DO NOT EDIT.

package flipt

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	status "google.golang.org/genproto/googleapis/rpc/status"
	status1 "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
	http "net/http"
	url "net/url"
)

type FliptHTTPClient struct {
	client *http.Client
	addr   string
}

func NewFliptHTTPClient(addr string) *FliptHTTPClient {
	return &FliptHTTPClient{client: http.DefaultClient, addr: addr}
}

func (x *FliptHTTPClient) Evaluate(ctx context.Context, v *EvaluationRequest) (*EvaluationResponse, error) {
	var body io.Reader
	var values url.Values
	reqData, err := protojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(reqData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, x.addr+"/api/v1/evaluate", body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output EvaluationResponse
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) BatchEvaluate(ctx context.Context, v *BatchEvaluationRequest) (*BatchEvaluationResponse, error) {
	var body io.Reader
	var values url.Values
	reqData, err := protojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(reqData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, x.addr+"/api/v1/batch-evaluate", body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output BatchEvaluationResponse
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) GetFlag(ctx context.Context, v *GetFlagRequest) (*Flag, error) {
	var body io.Reader
	var values url.Values
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, x.addr+fmt.Sprintf("/api/v1/flags/%v", v.Key), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output Flag
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) ListFlags(ctx context.Context, v *ListFlagRequest) (*FlagList, error) {
	var body io.Reader
	values := url.Values{}
	values.Set("limit", fmt.Sprintf("%v", v.Limit))
	values.Set("offset", fmt.Sprintf("%v", v.Offset))
	values.Set("pageToken", v.PageToken)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, x.addr+"/api/v1/flags", body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output FlagList
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) CreateFlag(ctx context.Context, v *CreateFlagRequest) (*Flag, error) {
	var body io.Reader
	var values url.Values
	reqData, err := protojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(reqData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, x.addr+"/api/v1/flags", body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output Flag
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) UpdateFlag(ctx context.Context, v *UpdateFlagRequest) (*Flag, error) {
	var body io.Reader
	var values url.Values
	reqData, err := protojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(reqData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, x.addr+fmt.Sprintf("/api/v1/flags/%v", v.Key), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output Flag
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) DeleteFlag(ctx context.Context, v *DeleteFlagRequest) (*emptypb.Empty, error) {
	var body io.Reader
	var values url.Values
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, x.addr+fmt.Sprintf("/api/v1/flags/%v", v.Key), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output emptypb.Empty
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) CreateVariant(ctx context.Context, v *CreateVariantRequest) (*Variant, error) {
	var body io.Reader
	var values url.Values
	reqData, err := protojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(reqData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, x.addr+"/api/v1/flags/{flag_key}/variants", body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output Variant
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) UpdateVariant(ctx context.Context, v *UpdateVariantRequest) (*Variant, error) {
	var body io.Reader
	var values url.Values
	reqData, err := protojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(reqData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, x.addr+fmt.Sprintf("/api/v1/flags/{flag_key}/variants/%v", v.Id), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output Variant
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) DeleteVariant(ctx context.Context, v *DeleteVariantRequest) (*emptypb.Empty, error) {
	var body io.Reader
	values := url.Values{}
	values.Set("flagKey", v.FlagKey)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, x.addr+fmt.Sprintf("/api/v1/flags/{flag_key}/variants/%v", v.Id), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output emptypb.Empty
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) GetRule(ctx context.Context, v *GetRuleRequest) (*Rule, error) {
	var body io.Reader
	values := url.Values{}
	values.Set("flagKey", v.FlagKey)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, x.addr+fmt.Sprintf("/api/v1/flags/{flag_key}/rules/%v", v.Id), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output Rule
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) ListRules(ctx context.Context, v *ListRuleRequest) (*RuleList, error) {
	var body io.Reader
	values := url.Values{}
	values.Set("limit", fmt.Sprintf("%v", v.Limit))
	values.Set("offset", fmt.Sprintf("%v", v.Offset))
	values.Set("flagKey", v.FlagKey)
	values.Set("pageToken", v.PageToken)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, x.addr+"/api/v1/flags/{flag_key}/rules", body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output RuleList
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) CreateRule(ctx context.Context, v *CreateRuleRequest) (*Rule, error) {
	var body io.Reader
	var values url.Values
	reqData, err := protojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(reqData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, x.addr+"/api/v1/flags/{flag_key}/rules", body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output Rule
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) UpdateRule(ctx context.Context, v *UpdateRuleRequest) (*Rule, error) {
	var body io.Reader
	var values url.Values
	reqData, err := protojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(reqData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, x.addr+fmt.Sprintf("/api/v1/flags/{flag_key}/rules/%v", v.Id), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output Rule
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) OrderRules(ctx context.Context, v *OrderRulesRequest) (*emptypb.Empty, error) {
	var body io.Reader
	var values url.Values
	reqData, err := protojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(reqData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, x.addr+"/api/v1/flags/{flag_key}/rules/order", body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output emptypb.Empty
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) DeleteRule(ctx context.Context, v *DeleteRuleRequest) (*emptypb.Empty, error) {
	var body io.Reader
	values := url.Values{}
	values.Set("flagKey", v.FlagKey)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, x.addr+fmt.Sprintf("/api/v1/flags/{flag_key}/rules/%v", v.Id), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output emptypb.Empty
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) CreateDistribution(ctx context.Context, v *CreateDistributionRequest) (*Distribution, error) {
	var body io.Reader
	var values url.Values
	reqData, err := protojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(reqData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, x.addr+"/api/v1/flags/{flag_key}/rules/{rule_id}/distributions", body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output Distribution
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) UpdateDistribution(ctx context.Context, v *UpdateDistributionRequest) (*Distribution, error) {
	var body io.Reader
	var values url.Values
	reqData, err := protojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(reqData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, x.addr+fmt.Sprintf("/api/v1/flags/{flag_key}/rules/{rule_id}/distributions/%v", v.Id), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output Distribution
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) DeleteDistribution(ctx context.Context, v *DeleteDistributionRequest) (*emptypb.Empty, error) {
	var body io.Reader
	values := url.Values{}
	values.Set("flagKey", v.FlagKey)
	values.Set("ruleId", v.RuleId)
	values.Set("variantId", v.VariantId)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, x.addr+fmt.Sprintf("/api/v1/flags/{flag_key}/rules/{rule_id}/distributions/%v", v.Id), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output emptypb.Empty
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) GetSegment(ctx context.Context, v *GetSegmentRequest) (*Segment, error) {
	var body io.Reader
	var values url.Values
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, x.addr+fmt.Sprintf("/api/v1/segments/%v", v.Key), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output Segment
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) ListSegments(ctx context.Context, v *ListSegmentRequest) (*SegmentList, error) {
	var body io.Reader
	values := url.Values{}
	values.Set("limit", fmt.Sprintf("%v", v.Limit))
	values.Set("offset", fmt.Sprintf("%v", v.Offset))
	values.Set("pageToken", v.PageToken)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, x.addr+"/api/v1/segments", body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output SegmentList
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) CreateSegment(ctx context.Context, v *CreateSegmentRequest) (*Segment, error) {
	var body io.Reader
	var values url.Values
	reqData, err := protojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(reqData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, x.addr+"/api/v1/segments", body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output Segment
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) UpdateSegment(ctx context.Context, v *UpdateSegmentRequest) (*Segment, error) {
	var body io.Reader
	var values url.Values
	reqData, err := protojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(reqData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, x.addr+fmt.Sprintf("/api/v1/segments/%v", v.Key), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output Segment
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) DeleteSegment(ctx context.Context, v *DeleteSegmentRequest) (*emptypb.Empty, error) {
	var body io.Reader
	var values url.Values
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, x.addr+fmt.Sprintf("/api/v1/segments/%v", v.Key), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output emptypb.Empty
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) CreateConstraint(ctx context.Context, v *CreateConstraintRequest) (*Constraint, error) {
	var body io.Reader
	var values url.Values
	reqData, err := protojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(reqData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, x.addr+"/api/v1/segments/{segment_key}/constraints", body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output Constraint
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) UpdateConstraint(ctx context.Context, v *UpdateConstraintRequest) (*Constraint, error) {
	var body io.Reader
	var values url.Values
	reqData, err := protojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(reqData)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, x.addr+fmt.Sprintf("/api/v1/segments/{segment_key}/constraints/%v", v.Id), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output Constraint
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (x *FliptHTTPClient) DeleteConstraint(ctx context.Context, v *DeleteConstraintRequest) (*emptypb.Empty, error) {
	var body io.Reader
	values := url.Values{}
	values.Set("segmentKey", v.SegmentKey)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, x.addr+fmt.Sprintf("/api/v1/segments/{segment_key}/constraints/%v", v.Id), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	resp, err := x.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output emptypb.Empty
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp, respData); err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(respData, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func checkResponse(resp *http.Response, v []byte) error {
	if resp.StatusCode != http.StatusOK {
		var status status.Status
		if err := protojson.Unmarshal(v, &status); err != nil {
			return err
		}
		return status1.ErrorProto(&status)
	}

	return nil
}
