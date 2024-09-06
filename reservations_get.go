package stayntouch

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-stayntouch/utils"
)

func (c *Client) NewReservationsGet() ReservationsGet {
	r := ReservationsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ReservationsGet struct {
	client      *Client
	queryParams *ReservationsGetQueryParams
	pathParams  *ReservationsGetPathParams
	method      string
	headers     http.Header
	requestBody ReservationsGetBody
}

func (r ReservationsGet) NewQueryParams() *ReservationsGetQueryParams {
	return &ReservationsGetQueryParams{}
}

type ReservationsGetQueryParams struct {
	HotelID               int    `schema:"hotel_id"`
	ConfirmationNumber    string `schema:"confirmation_number,omitempty"`
	AltConfirmationNumber string `schema:"alt_confirmation_number,omitempty"`
	LastName              string `schema:"last_name,omitempty"`
	DepartureDate         Date   `schema:"departure_date,omitempty"`
	Date                  Date   `schema:"date,omitempty"`
	FromDate              Date   `schema:"from_date,omitempty"`
	ToDate                Date   `schema:"to_date,omitempty"`
	DateFilter            string `schema:"date_filter,omitempty"`
	DateOperator          string `schema:"date_operator,omitempty"`
	GroupID               string `schema:"group_id,omitempty"`
	CreditCardLast4       string `schema:"credit_card_last_4,omitempty"`
	ReservationID         string `schema:"reservation_id,omitempty"`
	Email                 string `schema:"email,omitempty"`
	NoOfNights            int    `schema:"no_of_nights,omitempty"`
	ExcludeCanceled       bool   `schema:"exclude_canceled,omitempty"`
	CheckedIn             bool   `schema:"checked_in,omitempty"`
	Room                  string `schema:"room,omitempty"`
	Page                  int    `schema:"page,omitempty"`
	PerPage               int    `schema:"per_page,omitempty"`
}

func (p ReservationsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ReservationsGet) QueryParams() *ReservationsGetQueryParams {
	return r.queryParams
}

func (r ReservationsGet) NewPathParams() *ReservationsGetPathParams {
	return &ReservationsGetPathParams{}
}

type ReservationsGetPathParams struct {
}

func (p *ReservationsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ReservationsGet) PathParams() *ReservationsGetPathParams {
	return r.pathParams
}

func (r *ReservationsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ReservationsGet) SetMethod(method string) {
	r.method = method
}

func (r *ReservationsGet) Method() string {
	return r.method
}

func (r ReservationsGet) NewRequestBody() ReservationsGetBody {
	return ReservationsGetBody{}
}

type ReservationsGetBody struct {
}

func (r *ReservationsGet) RequestBody() *ReservationsGetBody {
	return nil
}

func (r *ReservationsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *ReservationsGet) SetRequestBody(body ReservationsGetBody) {
	r.requestBody = body
}

func (r *ReservationsGet) NewResponseBody() *ReservationsGetResponseBody {
	return &ReservationsGetResponseBody{}
}

type ReservationsGetResponseBody struct {
	Results    Reservations `json:"results"`
	TotalCount int          `json:"total_count"`
}

func (r *ReservationsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/reservations", r.PathParams())
	return &u
}

func (r *ReservationsGet) Do() (ReservationsGetResponseBody, error) {
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
