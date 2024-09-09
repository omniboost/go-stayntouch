package stayntouch

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-stayntouch/utils"
)

func (c *Client) NewRateGet() RateGet {
	r := RateGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type RateGet struct {
	client      *Client
	queryParams *RateGetQueryParams
	pathParams  *RateGetPathParams
	method      string
	headers     http.Header
	requestBody RateGetBody
}

func (r RateGet) NewQueryParams() *RateGetQueryParams {
	return &RateGetQueryParams{}
}

type RateGetQueryParams struct {
}

func (p RateGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *RateGet) QueryParams() *RateGetQueryParams {
	return r.queryParams
}

func (r RateGet) NewPathParams() *RateGetPathParams {
	return &RateGetPathParams{}
}

type RateGetPathParams struct {
	ID int `schema:"id"`
}

func (p *RateGetPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *RateGet) PathParams() *RateGetPathParams {
	return r.pathParams
}

func (r *RateGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *RateGet) SetMethod(method string) {
	r.method = method
}

func (r *RateGet) Method() string {
	return r.method
}

func (r RateGet) NewRequestBody() RateGetBody {
	return RateGetBody{}
}

type RateGetBody struct {
}

func (r *RateGet) RequestBody() *RateGetBody {
	return nil
}

func (r *RateGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *RateGet) SetRequestBody(body RateGetBody) {
	r.requestBody = body
}

func (r *RateGet) NewResponseBody() *RateGetResponseBody {
	return &RateGetResponseBody{}
}

type RateGetResponseBody Rate

func (r *RateGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/rates/{{.id}}", r.PathParams())
	return &u
}

func (r *RateGet) Do() (RateGetResponseBody, error) {
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

