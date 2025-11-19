package stayntouch

import (
	"context"
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
	HotelID string        `schema:"hotel_id,omitempty"`
	Date    Date          `schema:"date,omitempty"`
	Ledger  RevenueLedger `schema:"ledger,omitempty"`
	Page    int           `schema:"page,omitempty"`
	PerPage int           `schema:"per_page,omitempty"` // Don't use this, as the API ignores it at the moment.
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

type BillsRevenueGetPathParams struct{}

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
	Results    LedgerItems `json:"results"`
	TotalCount int         `json:"total_count"`
}

func (r *BillsRevenueGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/bills/revenue", r.PathParams())
	return &u
}

func (r *BillsRevenueGet) Do(ctx context.Context) (BillsRevenueGetResponseBody, error) {
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

func (r *BillsRevenueGet) All(ctx context.Context) (LedgerItems, error) {
	// Begin at page 1 and set per_page to 20.
	// Per page 20 is the maximum the API currently allows.
	r.queryParams.Page = 1
	r.queryParams.PerPage = 50

	// Set ledger items
	ledgerItems := LedgerItems{}
	for {
		response, err := r.Do(ctx)
		if err != nil {
			return LedgerItems{}, err
		}

		// Break if no results are returned
		if len(response.Results) == 0 {
			break
		}

		// Set the ledger items
		ledgerItems = append(ledgerItems, response.Results...)

		// Check if we have the total count already
		if len(ledgerItems) >= response.TotalCount {
			break
		}

		// Set the next page
		r.queryParams.Page++
	}

	return ledgerItems, nil
}
