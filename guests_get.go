package stayntouch

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-stayntouch/utils"
)

func (c *Client) NewGuestsGet() GuestsGet {
	r := GuestsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GuestsGet struct {
	client      *Client
	queryParams *GuestsGetQueryParams
	pathParams  *GuestsGetPathParams
	method      string
	headers     http.Header
	requestBody GuestsGetBody
}

func (r GuestsGet) NewQueryParams() *GuestsGetQueryParams {
	return &GuestsGetQueryParams{}
}

type GuestsGetQueryParams struct {
	// First Name of Guest
	FirstName string `schema:"first_name,omitempty"`
	// Last Name of Guest
	LastName string `schema:"last_name,omitempty"`
	// City of Guest
	City string `schema:"city,omitempty"`
	// Email of Guest
	Email string `schema:"email,omitempty"`
	// Membership card number of Guest
	MembershipNo string `schema:"membership_no,omitempty"`
	// Search text
	Query string `schema:"query,omitempty"`
	// Specify the page that you want to retrieve
	Page int `schema:"page,omitempty"`
	// The number of results per page, maximum 50
	PerPage int `schema:"per_page,omitempty"`
}

func (p GuestsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GuestsGet) QueryParams() *GuestsGetQueryParams {
	return r.queryParams
}

func (r GuestsGet) NewPathParams() *GuestsGetPathParams {
	return &GuestsGetPathParams{}
}

type GuestsGetPathParams struct {
}

func (p *GuestsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GuestsGet) PathParams() *GuestsGetPathParams {
	return r.pathParams
}

func (r *GuestsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GuestsGet) SetMethod(method string) {
	r.method = method
}

func (r *GuestsGet) Method() string {
	return r.method
}

func (r GuestsGet) NewRequestBody() GuestsGetBody {
	return GuestsGetBody{}
}

type GuestsGetBody struct {
}

func (r *GuestsGet) RequestBody() *GuestsGetBody {
	return nil
}

func (r *GuestsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *GuestsGet) SetRequestBody(body GuestsGetBody) {
	r.requestBody = body
}

func (r *GuestsGet) NewResponseBody() *GuestsGetResponseBody {
	return &GuestsGetResponseBody{}
}

type GuestsGetResponseBody struct {
	Results    Guests `json:"results"`
	TotalCount int    `json:"total_count"`
}

func (r *GuestsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/guests", r.PathParams())
	return &u
}

func (r *GuestsGet) Do() (GuestsGetResponseBody, error) {
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
