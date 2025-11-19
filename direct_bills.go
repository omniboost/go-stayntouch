package stayntouch

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-stayntouch/utils"
)

func (c *Client) NewDirectBillsGet() DirectBillsGet {
	r := DirectBillsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type DirectBillsGet struct {
	client      *Client
	queryParams *DirectBillsGetQueryParams
	pathParams  *DirectBillsGetPathParams
	method      string
	headers     http.Header
	requestBody DirectBillsGetBody
}

func (r DirectBillsGet) NewQueryParams() *DirectBillsGetQueryParams {
	return &DirectBillsGetQueryParams{}
}

type DirectBillsGetQueryParams struct {
	// Hotel Identifier
	HotelID int `schema:"hotel_id"`
	// Business date when direct bill was settled
	Date Date `schema:"date"`
	// Specify the page that you want to retrieve
	Page int `schema:"page,omitempty"`
	// The number of results per page. A maximum of 10 records will be displayed.
	PerPage int `schema:"per_page,omitempty"`
}

func (p DirectBillsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *DirectBillsGet) QueryParams() *DirectBillsGetQueryParams {
	return r.queryParams
}

func (r DirectBillsGet) NewPathParams() *DirectBillsGetPathParams {
	return &DirectBillsGetPathParams{}
}

type DirectBillsGetPathParams struct {
}

func (p *DirectBillsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DirectBillsGet) PathParams() *DirectBillsGetPathParams {
	return r.pathParams
}

func (r *DirectBillsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DirectBillsGet) SetMethod(method string) {
	r.method = method
}

func (r *DirectBillsGet) Method() string {
	return r.method
}

func (r DirectBillsGet) NewRequestBody() DirectBillsGetBody {
	return DirectBillsGetBody{}
}

type DirectBillsGetBody struct {
}

func (r *DirectBillsGet) RequestBody() *DirectBillsGetBody {
	return nil
}

func (r *DirectBillsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *DirectBillsGet) SetRequestBody(body DirectBillsGetBody) {
	r.requestBody = body
}

func (r *DirectBillsGet) NewResponseBody() *DirectBillsGetResponseBody {
	return &DirectBillsGetResponseBody{}
}

type DirectBillsGetResponseBody struct {
	Results    DirectBills `json:"results"`
	TotalCount int         `json:"total_count"`
}

func (r *DirectBillsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("bills/direct_bills", r.PathParams())
	return &u
}

func (r *DirectBillsGet) Do(ctx context.Context) (DirectBillsGetResponseBody, error) {
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

func (r *DirectBillsGet) All(ctx context.Context) (DirectBills, error) {
	// Begin at page 1 and set per_page to 20.
	// Per page 20 is the maximum the API currently allows.
	r.queryParams.Page = 1
	r.queryParams.PerPage = 50

	// Set ledger items
	bills := DirectBills{}
	for {
		response, err := r.Do(ctx)
		if err != nil {
			return DirectBills{}, err
		}

		// Break if no results are returned
		if len(response.Results) == 0 {
			break
		}

		// Set the ledger items
		bills = append(bills, response.Results...)

		// Check if we have the total count already
		if len(bills) >= response.TotalCount {
			break
		}

		// Set the next page
		r.queryParams.Page++
	}

	return bills, nil
}
