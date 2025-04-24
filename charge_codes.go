package stayntouch

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-stayntouch/utils"
)

func (c *Client) NewChargeCodesGet() ChargeCodesGet {
	r := ChargeCodesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ChargeCodesGet struct {
	client      *Client
	queryParams *ChargeCodesGetQueryParams
	pathParams  *ChargeCodesGetPathParams
	method      string
	headers     http.Header
	requestBody ChargeCodesGetBody
}

func (r ChargeCodesGet) NewQueryParams() *ChargeCodesGetQueryParams {
	return &ChargeCodesGetQueryParams{}
}

type ChargeCodesGetQueryParams struct {
	// Hotel Identifier
	HotelID int `schema:"hotel_id"`
	// Language code in ISO 639-1 Format
	LanguageCode string `schema:"language_code,omitempty"`
	// Specify the page that you want to retrieve
	Page int `schema:"page,omitempty"`
	// The number of results per page, maximum 50
	PerPage int `schema:"per_page,omitempty"`
}

func (p ChargeCodesGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ChargeCodesGet) QueryParams() *ChargeCodesGetQueryParams {
	return r.queryParams
}

func (r ChargeCodesGet) NewPathParams() *ChargeCodesGetPathParams {
	return &ChargeCodesGetPathParams{}
}

type ChargeCodesGetPathParams struct {
}

func (p *ChargeCodesGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ChargeCodesGet) PathParams() *ChargeCodesGetPathParams {
	return r.pathParams
}

func (r *ChargeCodesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ChargeCodesGet) SetMethod(method string) {
	r.method = method
}

func (r *ChargeCodesGet) Method() string {
	return r.method
}

func (r ChargeCodesGet) NewRequestBody() ChargeCodesGetBody {
	return ChargeCodesGetBody{}
}

type ChargeCodesGetBody struct {
}

func (r *ChargeCodesGet) RequestBody() *ChargeCodesGetBody {
	return nil
}

func (r *ChargeCodesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *ChargeCodesGet) SetRequestBody(body ChargeCodesGetBody) {
	r.requestBody = body
}

func (r *ChargeCodesGet) NewResponseBody() *ChargeCodesGetResponseBody {
	return &ChargeCodesGetResponseBody{}
}

type ChargeCodesGetResponseBody struct {
	Results    ChargeCodes `json:"results"`
	TotalCount int         `json:"total_count"`
}

func (r *ChargeCodesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("charge_codes/translations", r.PathParams())
	return &u
}

func (r *ChargeCodesGet) Do() (ChargeCodesGetResponseBody, error) {
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
