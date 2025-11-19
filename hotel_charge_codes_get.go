package stayntouch

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-stayntouch/utils"
)

func (c *Client) NewHotelChargeCodesGet() HotelChargeCodesGet {
	r := HotelChargeCodesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type HotelChargeCodesGet struct {
	client      *Client
	queryParams *HotelChargeCodesGetQueryParams
	pathParams  *HotelChargeCodesGetPathParams
	method      string
	headers     http.Header
	requestBody HotelChargeCodesGetBody
}

func (r HotelChargeCodesGet) NewQueryParams() *HotelChargeCodesGetQueryParams {
	return &HotelChargeCodesGetQueryParams{}
}

type HotelChargeCodesGetQueryParams struct {
	// Specify the page that you want to retrieve
	Page int `schema:"page,omitempty"`
	// The number of results per page, maximum 50
	PerPage int `schema:"per_page,omitempty"`
}

func (p HotelChargeCodesGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *HotelChargeCodesGet) QueryParams() *HotelChargeCodesGetQueryParams {
	return r.queryParams
}

func (r HotelChargeCodesGet) NewPathParams() *HotelChargeCodesGetPathParams {
	return &HotelChargeCodesGetPathParams{}
}

type HotelChargeCodesGetPathParams struct {
	HotelID int `schema:"id"`
}

func (p *HotelChargeCodesGetPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.HotelID),
	}
}

func (r *HotelChargeCodesGet) PathParams() *HotelChargeCodesGetPathParams {
	return r.pathParams
}

func (r *HotelChargeCodesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *HotelChargeCodesGet) SetMethod(method string) {
	r.method = method
}

func (r *HotelChargeCodesGet) Method() string {
	return r.method
}

func (r HotelChargeCodesGet) NewRequestBody() HotelChargeCodesGetBody {
	return HotelChargeCodesGetBody{}
}

type HotelChargeCodesGetBody struct {
}

func (r *HotelChargeCodesGet) RequestBody() *HotelChargeCodesGetBody {
	return nil
}

func (r *HotelChargeCodesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *HotelChargeCodesGet) SetRequestBody(body HotelChargeCodesGetBody) {
	r.requestBody = body
}

func (r *HotelChargeCodesGet) NewResponseBody() *HotelChargeCodesGetResponseBody {
	return &HotelChargeCodesGetResponseBody{}
}

type HotelChargeCodesGetResponseBody struct {
	Results    HotelChargeCodes `json:"results"`
	TotalCount int              `json:"total_count"`
}

func (r *HotelChargeCodesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("hotels/{{.id}}/charge_codes", r.PathParams())
	return &u
}

func (r *HotelChargeCodesGet) Do(ctx context.Context) (HotelChargeCodesGetResponseBody, error) {
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

func (r *HotelChargeCodesGet) All(ctx context.Context) (HotelChargeCodes, error) {
	// Begin at page 1 and set per_page to 20.
	// Per page 20 is the maximum the API currently allows.
	r.queryParams.Page = 1
	r.queryParams.PerPage = 50

	// Set ledger items
	chargeCodes := HotelChargeCodes{}
	for {
		response, err := r.Do(ctx)
		if err != nil {
			return HotelChargeCodes{}, err
		}

		// Break if no results are returned
		if len(response.Results) == 0 {
			break
		}

		// Set the ledger items
		chargeCodes = append(chargeCodes, response.Results...)

		// Check if we have the total count already
		if len(chargeCodes) >= response.TotalCount {
			break
		}

		// Set the next page
		r.queryParams.Page++
	}

	return chargeCodes, nil
}
