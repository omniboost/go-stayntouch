package stayntouch

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-stayntouch/utils"
)

func (c *Client) NewRoomsGet() RoomsGet {
	r := RoomsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type RoomsGet struct {
	client      *Client
	queryParams *RoomsGetQueryParams
	pathParams  *RoomsGetPathParams
	method      string
	headers     http.Header
	requestBody RoomsGetBody
}

func (r RoomsGet) NewQueryParams() *RoomsGetQueryParams {
	return &RoomsGetQueryParams{}
}

type RoomsGetQueryParams struct {
	// Hotel Identifier
	HotelID int `schema:"hotel_id"`
	// Specify the page that you want to retrieve
	Page int `schema:"page,omitempty"`
	// The number of results per page, maximum 50
	PerPage int `schema:"per_page,omitempty"`
}

func (p RoomsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *RoomsGet) QueryParams() *RoomsGetQueryParams {
	return r.queryParams
}

func (r RoomsGet) NewPathParams() *RoomsGetPathParams {
	return &RoomsGetPathParams{}
}

type RoomsGetPathParams struct {
}

func (p *RoomsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *RoomsGet) PathParams() *RoomsGetPathParams {
	return r.pathParams
}

func (r *RoomsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *RoomsGet) SetMethod(method string) {
	r.method = method
}

func (r *RoomsGet) Method() string {
	return r.method
}

func (r RoomsGet) NewRequestBody() RoomsGetBody {
	return RoomsGetBody{}
}

type RoomsGetBody struct {
}

func (r *RoomsGet) RequestBody() *RoomsGetBody {
	return nil
}

func (r *RoomsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *RoomsGet) SetRequestBody(body RoomsGetBody) {
	r.requestBody = body
}

func (r *RoomsGet) NewResponseBody() *RoomsGetResponseBody {
	return &RoomsGetResponseBody{}
}

type RoomsGetResponseBody struct {
	Results    Rooms `json:"results"`
	TotalCount int   `json:"total_count"`
}

func (r *RoomsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/rooms", r.PathParams())
	return &u
}

func (r *RoomsGet) Do() (RoomsGetResponseBody, error) {
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
