package stayntouch

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-stayntouch/utils"
)

func (c *Client) NewReservationGet() ReservationGet {
	r := ReservationGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ReservationGet struct {
	client      *Client
	queryParams *ReservationGetQueryParams
	pathParams  *ReservationGetPathParams
	method      string
	headers     http.Header
	requestBody ReservationGetBody
}

func (r ReservationGet) NewQueryParams() *ReservationGetQueryParams {
	return &ReservationGetQueryParams{}
}

type ReservationGetQueryParams struct {
}

func (p ReservationGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ReservationGet) QueryParams() *ReservationGetQueryParams {
	return r.queryParams
}

func (r ReservationGet) NewPathParams() *ReservationGetPathParams {
	return &ReservationGetPathParams{}
}

type ReservationGetPathParams struct {
	ID int `schema:"id"`
}

func (p *ReservationGetPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *ReservationGet) PathParams() *ReservationGetPathParams {
	return r.pathParams
}

func (r *ReservationGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ReservationGet) SetMethod(method string) {
	r.method = method
}

func (r *ReservationGet) Method() string {
	return r.method
}

func (r ReservationGet) NewRequestBody() ReservationGetBody {
	return ReservationGetBody{}
}

type ReservationGetBody struct {
}

func (r *ReservationGet) RequestBody() *ReservationGetBody {
	return nil
}

func (r *ReservationGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *ReservationGet) SetRequestBody(body ReservationGetBody) {
	r.requestBody = body
}

func (r *ReservationGet) NewResponseBody() *ReservationGetResponseBody {
	return &ReservationGetResponseBody{}
}

type ReservationGetResponseBody struct {
	Reservation
}

func (r *ReservationGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/reservations/{{.id}}", r.PathParams())
	return &u
}

func (r *ReservationGet) Do(ctx context.Context) (ReservationGetResponseBody, error) {
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
