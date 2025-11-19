package stayntouch

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-stayntouch/utils"
)

func (c *Client) NewInvoiceGet() InvoiceGet {
	r := InvoiceGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type InvoiceGet struct {
	client      *Client
	queryParams *InvoiceGetQueryParams
	pathParams  *InvoiceGetPathParams
	method      string
	headers     http.Header
	requestBody InvoiceGetBody
}

func (r InvoiceGet) NewQueryParams() *InvoiceGetQueryParams {
	return &InvoiceGetQueryParams{}
}

type InvoiceGetQueryParams struct {
	// Hotel Identifier
	HotelID int `schema:"hotel_id"`
}

func (p InvoiceGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *InvoiceGet) QueryParams() *InvoiceGetQueryParams {
	return r.queryParams
}

func (r InvoiceGet) NewPathParams() *InvoiceGetPathParams {
	return &InvoiceGetPathParams{}
}

type InvoiceGetPathParams struct {
	ID string `schema:"id"`
}

func (p *InvoiceGetPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *InvoiceGet) PathParams() *InvoiceGetPathParams {
	return r.pathParams
}

func (r *InvoiceGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *InvoiceGet) SetMethod(method string) {
	r.method = method
}

func (r *InvoiceGet) Method() string {
	return r.method
}

func (r InvoiceGet) NewRequestBody() InvoiceGetBody {
	return InvoiceGetBody{}
}

type InvoiceGetBody struct {
}

func (r *InvoiceGet) RequestBody() *InvoiceGetBody {
	return nil
}

func (r *InvoiceGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *InvoiceGet) SetRequestBody(body InvoiceGetBody) {
	r.requestBody = body
}

func (r *InvoiceGet) NewResponseBody() *InvoiceGetResponseBody {
	return &InvoiceGetResponseBody{}
}

type InvoiceGetResponseBody DirectBill

func (r *InvoiceGet) URL() *url.URL {
	u := r.client.GetEndpointURL("bills/{{.id}}/invoice", r.PathParams())
	return &u
}

func (r *InvoiceGet) Do(ctx context.Context) (InvoiceGetResponseBody, error) {
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
