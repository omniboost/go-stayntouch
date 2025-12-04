package stayntouch

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-stayntouch/utils"
)

func (c *Client) NewInvoicesPerDayGet() InvoicesPerDayGet {
	r := InvoicesPerDayGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type InvoicesPerDayGet struct {
	client      *Client
	queryParams *InvoicesPerDayGetQueryParams
	pathParams  *InvoicesPerDayGetPathParams
	method      string
	headers     http.Header
	requestBody InvoicesPerDayGetBody
}

func (r InvoicesPerDayGet) NewQueryParams() *InvoicesPerDayGetQueryParams {
	return &InvoicesPerDayGetQueryParams{}
}

type InvoicesPerDayGetQueryParams struct {
	// Hotel Identifier
	HotelID int `schema:"hotel_id"`
	// Business date when direct bill was settled
	Date Date `schema:"date"`
	// Specify the page that you want to retrieve
	Page int `schema:"page,omitempty"`
	// The number of results per page. A maximum of 10 records will be displayed.
	PerPage int `schema:"per_page,omitempty"`
}

func (p InvoicesPerDayGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *InvoicesPerDayGet) QueryParams() *InvoicesPerDayGetQueryParams {
	return r.queryParams
}

func (r InvoicesPerDayGet) NewPathParams() *InvoicesPerDayGetPathParams {
	return &InvoicesPerDayGetPathParams{}
}

type InvoicesPerDayGetPathParams struct {
}

func (p *InvoicesPerDayGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *InvoicesPerDayGet) PathParams() *InvoicesPerDayGetPathParams {
	return r.pathParams
}

func (r *InvoicesPerDayGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *InvoicesPerDayGet) SetMethod(method string) {
	r.method = method
}

func (r *InvoicesPerDayGet) Method() string {
	return r.method
}

func (r InvoicesPerDayGet) NewRequestBody() InvoicesPerDayGetBody {
	return InvoicesPerDayGetBody{}
}

type InvoicesPerDayGetBody struct {
}

func (r *InvoicesPerDayGet) RequestBody() *InvoicesPerDayGetBody {
	return nil
}

func (r *InvoicesPerDayGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *InvoicesPerDayGet) SetRequestBody(body InvoicesPerDayGetBody) {
	r.requestBody = body
}

func (r *InvoicesPerDayGet) NewResponseBody() *InvoicesPerDayGetResponseBody {
	return &InvoicesPerDayGetResponseBody{}
}

type InvoicesPerDayGetResponseBody struct {
	Results    InvoicesPerDay `json:"results"`
	TotalCount int            `json:"total_count"`
}

func (r *InvoicesPerDayGet) URL() *url.URL {
	u := r.client.GetEndpointURL("bills/invoices_per_day", r.PathParams())
	return &u
}

func (r *InvoicesPerDayGet) Do(ctx context.Context) (InvoicesPerDayGetResponseBody, error) {
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

func (r *InvoicesPerDayGet) All(ctx context.Context) (InvoicesPerDay, error) {
	// Begin at page 1 and set per_page to 20.
	// Per page 20 is the maximum the API currently allows.
	r.queryParams.Page = 1
	r.queryParams.PerPage = 20

	// Set ledger items
	invoicesPerDay := InvoicesPerDay{}
	for {
		response, err := r.Do(ctx)
		if err != nil {
			return InvoicesPerDay{}, err
		}

		// Break if no results are returned
		if len(response.Results) == 0 {
			break
		}

		// Set the ledger items
		invoicesPerDay = append(invoicesPerDay, response.Results...)

		// Check if we have the total count already
		if len(invoicesPerDay) >= response.TotalCount {
			break
		}

		// Set the next page
		r.queryParams.Page++
	}

	return invoicesPerDay, nil
}
