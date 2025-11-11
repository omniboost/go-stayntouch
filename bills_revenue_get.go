package stayntouch

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-stayntouch/utils"
)

func (c *Client) NewBillsRevenueGet() BillsRevenueGet {
	r := BillsRevenueGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BillsRevenueGet struct {
	client      *Client
	queryParams *BillsRevenueGetQueryParams
	pathParams  *BillsRevenueGetPathParams
	method      string
	headers     http.Header
	requestBody BillsRevenueGetBody
}

func (r BillsRevenueGet) NewQueryParams() *BillsRevenueGetQueryParams {
	return &BillsRevenueGetQueryParams{}
}

type BillsRevenueGetQueryParams struct {
	// Hotel Identifier
	HotelID int `schema:"hotel_id"`
	// Business date for revenue data
	Date Date `schema:"date"`
	// Specify the page that you want to retrieve
	Page int `schema:"page,omitempty"`
	// The number of results per page
	PerPage int `schema:"per_page,omitempty"`
}

func (p BillsRevenueGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *BillsRevenueGet) QueryParams() *BillsRevenueGetQueryParams {
	return r.queryParams
}

func (r BillsRevenueGet) NewPathParams() *BillsRevenueGetPathParams {
	return &BillsRevenueGetPathParams{}
}

type BillsRevenueGetPathParams struct {
}

func (p *BillsRevenueGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *BillsRevenueGet) PathParams() *BillsRevenueGetPathParams {
	return r.pathParams
}

func (r *BillsRevenueGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BillsRevenueGet) SetMethod(method string) {
	r.method = method
}

func (r *BillsRevenueGet) Method() string {
	return r.method
}

func (r BillsRevenueGet) NewRequestBody() BillsRevenueGetBody {
	return BillsRevenueGetBody{}
}

type BillsRevenueGetBody struct {
}

func (r *BillsRevenueGet) RequestBody() *BillsRevenueGetBody {
	return nil
}

func (r *BillsRevenueGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *BillsRevenueGet) SetRequestBody(body BillsRevenueGetBody) {
	r.requestBody = body
}

func (r *BillsRevenueGet) NewResponseBody() *BillsRevenueGetResponseBody {
	return &BillsRevenueGetResponseBody{}
}

type BillsRevenueGetResponseBody struct {
	Results    RevenueItems `json:"results"`
	TotalCount int          `json:"total_count"`
}

func (r *BillsRevenueGet) URL() *url.URL {
	u := r.client.GetEndpointURL("bills/revenue", r.PathParams())
	return &u
}

func (r *BillsRevenueGet) Do() (BillsRevenueGetResponseBody, error) {
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
