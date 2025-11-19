package stayntouch

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-stayntouch/utils"
)

func (c *Client) NewChargeCodesBasicTaxesGet() ChargeCodesBasicTaxesGet {
	r := ChargeCodesBasicTaxesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ChargeCodesBasicTaxesGet struct {
	client      *Client
	queryParams *ChargeCodesBasicTaxesGetQueryParams
	pathParams  *ChargeCodesBasicTaxesGetPathParams
	method      string
	headers     http.Header
	requestBody ChargeCodesBasicTaxesGetBody
}

func (r ChargeCodesBasicTaxesGet) NewQueryParams() *ChargeCodesBasicTaxesGetQueryParams {
	return &ChargeCodesBasicTaxesGetQueryParams{}
}

type ChargeCodesBasicTaxesGetQueryParams struct {
	// Hotel Identifier
	HotelID int `schema:"hotel_id"`
	// Charge Code Identifier
	ID string `schema:"id"`
}

func (p ChargeCodesBasicTaxesGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ChargeCodesBasicTaxesGet) QueryParams() *ChargeCodesBasicTaxesGetQueryParams {
	return r.queryParams
}

func (r ChargeCodesBasicTaxesGet) NewPathParams() *ChargeCodesBasicTaxesGetPathParams {
	return &ChargeCodesBasicTaxesGetPathParams{}
}

type ChargeCodesBasicTaxesGetPathParams struct {
	ID int `schema:"id"`
}

func (p *ChargeCodesBasicTaxesGetPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *ChargeCodesBasicTaxesGet) PathParams() *ChargeCodesBasicTaxesGetPathParams {
	return r.pathParams
}

func (r *ChargeCodesBasicTaxesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ChargeCodesBasicTaxesGet) SetMethod(method string) {
	r.method = method
}

func (r *ChargeCodesBasicTaxesGet) Method() string {
	return r.method
}

func (r ChargeCodesBasicTaxesGet) NewRequestBody() ChargeCodesBasicTaxesGetBody {
	return ChargeCodesBasicTaxesGetBody{}
}

type ChargeCodesBasicTaxesGetBody struct {
}

func (r *ChargeCodesBasicTaxesGet) RequestBody() *ChargeCodesBasicTaxesGetBody {
	return nil
}

func (r *ChargeCodesBasicTaxesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *ChargeCodesBasicTaxesGet) SetRequestBody(body ChargeCodesBasicTaxesGetBody) {
	r.requestBody = body
}

func (r *ChargeCodesBasicTaxesGet) NewResponseBody() *ChargeCodesBasicTaxesGetResponseBody {
	return &ChargeCodesBasicTaxesGetResponseBody{}
}

type ChargeCodesBasicTaxesGetResponseBody ChargeCodesBasicTaxes

func (r *ChargeCodesBasicTaxesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("charge_codes/{{.id}}/basic_taxes", r.PathParams())
	return &u
}

func (r *ChargeCodesBasicTaxesGet) Do(ctx context.Context) (ChargeCodesBasicTaxesGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(ctx, r)
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
