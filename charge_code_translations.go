package stayntouch

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-stayntouch/utils"
)

func (c *Client) NewChargeCodeTranslationsGet() ChargeCodeTranslationsGet {
	r := ChargeCodeTranslationsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ChargeCodeTranslationsGet struct {
	client      *Client
	queryParams *ChargeCodeTranslationsGetQueryParams
	pathParams  *ChargeCodeTranslationsGetPathParams
	method      string
	headers     http.Header
	requestBody ChargeCodeTranslationsGetBody
}

func (r ChargeCodeTranslationsGet) NewQueryParams() *ChargeCodeTranslationsGetQueryParams {
	return &ChargeCodeTranslationsGetQueryParams{}
}

type ChargeCodeTranslationsGetQueryParams struct {
	// Hotel Identifier
	HotelID int `schema:"hotel_id"`
	// Language code in ISO 639-1 Format
	LanguageCode string `schema:"language_code,omitempty"`
	// Specify the page that you want to retrieve
	Page int `schema:"page,omitempty"`
	// The number of results per page, maximum 50
	PerPage int `schema:"per_page,omitempty"`
}

func (p ChargeCodeTranslationsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ChargeCodeTranslationsGet) QueryParams() *ChargeCodeTranslationsGetQueryParams {
	return r.queryParams
}

func (r ChargeCodeTranslationsGet) NewPathParams() *ChargeCodeTranslationsGetPathParams {
	return &ChargeCodeTranslationsGetPathParams{}
}

type ChargeCodeTranslationsGetPathParams struct {
}

func (p *ChargeCodeTranslationsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ChargeCodeTranslationsGet) PathParams() *ChargeCodeTranslationsGetPathParams {
	return r.pathParams
}

func (r *ChargeCodeTranslationsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ChargeCodeTranslationsGet) SetMethod(method string) {
	r.method = method
}

func (r *ChargeCodeTranslationsGet) Method() string {
	return r.method
}

func (r ChargeCodeTranslationsGet) NewRequestBody() ChargeCodeTranslationsGetBody {
	return ChargeCodeTranslationsGetBody{}
}

type ChargeCodeTranslationsGetBody struct {
}

func (r *ChargeCodeTranslationsGet) RequestBody() *ChargeCodeTranslationsGetBody {
	return nil
}

func (r *ChargeCodeTranslationsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *ChargeCodeTranslationsGet) SetRequestBody(body ChargeCodeTranslationsGetBody) {
	r.requestBody = body
}

func (r *ChargeCodeTranslationsGet) NewResponseBody() *ChargeCodeTranslationsGetResponseBody {
	return &ChargeCodeTranslationsGetResponseBody{}
}

type ChargeCodeTranslationsGetResponseBody struct {
	Results    ChargeCodeTranslations `json:"results"`
	TotalCount int                    `json:"total_count"`
}

func (r *ChargeCodeTranslationsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("charge_codes/translations", r.PathParams())
	return &u
}

func (r *ChargeCodeTranslationsGet) Do() (ChargeCodeTranslationsGetResponseBody, error) {
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
