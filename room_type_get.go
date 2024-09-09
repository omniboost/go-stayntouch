package stayntouch

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-stayntouch/utils"
)

func (c *Client) NewRoomTypeGet() RoomTypeGet {
	r := RoomTypeGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type RoomTypeGet struct {
	client      *Client
	queryParams *RoomTypeGetQueryParams
	pathParams  *RoomTypeGetPathParams
	method      string
	headers     http.Header
	requestBody RoomTypeGetBody
}

func (r RoomTypeGet) NewQueryParams() *RoomTypeGetQueryParams {
	return &RoomTypeGetQueryParams{}
}

type RoomTypeGetQueryParams struct {
}

func (p RoomTypeGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *RoomTypeGet) QueryParams() *RoomTypeGetQueryParams {
	return r.queryParams
}

func (r RoomTypeGet) NewPathParams() *RoomTypeGetPathParams {
	return &RoomTypeGetPathParams{}
}

type RoomTypeGetPathParams struct {
	ID int `schema:"id"`
}

func (p *RoomTypeGetPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *RoomTypeGet) PathParams() *RoomTypeGetPathParams {
	return r.pathParams
}

func (r *RoomTypeGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *RoomTypeGet) SetMethod(method string) {
	r.method = method
}

func (r *RoomTypeGet) Method() string {
	return r.method
}

func (r RoomTypeGet) NewRequestBody() RoomTypeGetBody {
	return RoomTypeGetBody{}
}

type RoomTypeGetBody struct {
}

func (r *RoomTypeGet) RequestBody() *RoomTypeGetBody {
	return nil
}

func (r *RoomTypeGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *RoomTypeGet) SetRequestBody(body RoomTypeGetBody) {
	r.requestBody = body
}

func (r *RoomTypeGet) NewResponseBody() *RoomTypeGetResponseBody {
	return &RoomTypeGetResponseBody{}
}

type RoomTypeGetResponseBody RoomType

func (r *RoomTypeGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/room_types/{{.id}}", r.PathParams())
	return &u
}

func (r *RoomTypeGet) Do() (RoomTypeGetResponseBody, error) {
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
