package stayntouch

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-stayntouch/utils"
)

func (c *Client) NewHotelsGet() HotelsGet {
	r := HotelsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type HotelsGet struct {
	client      *Client
	queryParams *HotelsGetQueryParams
	pathParams  *HotelsGetPathParams
	method      string
	headers     http.Header
	requestBody HotelsGetBody
}

func (r HotelsGet) NewQueryParams() *HotelsGetQueryParams {
	return &HotelsGetQueryParams{}
}

type HotelsGetQueryParams struct {
}

func (p HotelsGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *HotelsGet) QueryParams() *HotelsGetQueryParams {
	return r.queryParams
}

func (r HotelsGet) NewPathParams() *HotelsGetPathParams {
	return &HotelsGetPathParams{}
}

type HotelsGetPathParams struct {
}

func (p *HotelsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *HotelsGet) PathParams() *HotelsGetPathParams {
	return r.pathParams
}

func (r *HotelsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *HotelsGet) SetMethod(method string) {
	r.method = method
}

func (r *HotelsGet) Method() string {
	return r.method
}

func (r HotelsGet) NewRequestBody() HotelsGetBody {
	return HotelsGetBody{}
}

type HotelsGetBody struct {
}

func (r *HotelsGet) RequestBody() *HotelsGetBody {
	return nil
}

func (r *HotelsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *HotelsGet) SetRequestBody(body HotelsGetBody) {
	r.requestBody = body
}

func (r *HotelsGet) NewResponseBody() *HotelsGetResponseBody {
	return &HotelsGetResponseBody{}
}

type HotelsGetResponseBody struct {
	Results    Hotels `json:"results"`
	TotalCount int    `json:"total_count"`
}

func (r *HotelsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/hotels", r.PathParams())
	return &u
}

func (r *HotelsGet) Do() (HotelsGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
