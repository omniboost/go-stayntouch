package stayntouch

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/omniboost/go-stayntouch/utils"
)

func (c *Client) NewAccountsGet() AccountsGet {
	r := AccountsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AccountsGet struct {
	client      *Client
	queryParams *AccountsGetQueryParams
	pathParams  *AccountsGetPathParams
	method      string
	headers     http.Header
	requestBody AccountsGetBody
}

func (r AccountsGet) NewQueryParams() *AccountsGetQueryParams {
	return &AccountsGetQueryParams{}
}

type AccountsGetQueryParams struct {
	// Hotel Identifier
	HotelID int `schema:"hotel_id"`
	// Possible values: COMPANY, TRAVELAGENT
	Type string `schema:"type,omitempty"`
	// COMPANY or TRAVELAGENT Contract Access Code, valid on current business date only.
	AccessCode string `schema:"access_code,omitempty"`
	// Minimum first 3 characters of the COMPANY or TRAVEL AGENT name
	SearchKeyword string `schema:"search_keyword,omitempty"`
	// Specify the page that you want to retrieve
	Page int `schema:"page,omitempty"`
	// The number of results per page, maximum 50
	PerPage int `schema:"per_page,omitempty"`
}

func (p AccountsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *AccountsGet) QueryParams() *AccountsGetQueryParams {
	return r.queryParams
}

func (r AccountsGet) NewPathParams() *AccountsGetPathParams {
	return &AccountsGetPathParams{}
}

type AccountsGetPathParams struct {
	ID string `schema:"id,omitempty"`
}

func (p *AccountsGetPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *AccountsGet) PathParams() *AccountsGetPathParams {
	return r.pathParams
}

func (r *AccountsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AccountsGet) SetMethod(method string) {
	r.method = method
}

func (r *AccountsGet) Method() string {
	return r.method
}

func (r AccountsGet) NewRequestBody() AccountsGetBody {
	return AccountsGetBody{}
}

type AccountsGetBody struct {
}

func (r *AccountsGet) RequestBody() *AccountsGetBody {
	return nil
}

func (r *AccountsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *AccountsGet) SetRequestBody(body AccountsGetBody) {
	r.requestBody = body
}

func (r *AccountsGet) NewResponseBody() *AccountsGetResponseBody {
	return &AccountsGetResponseBody{}
}

type AccountsGetResponseBody Accounts

func (v *AccountsGetResponseBody) UnmarshalJSON(b []byte) (err error) {
	vc, vcs := Account{}, Accounts{}
	if err = json.Unmarshal(b, &vcs); err == nil && vcs.Results != nil {
		*v = AccountsGetResponseBody(Accounts{
			Results:    vcs.Results,
			TotalCount: vcs.TotalCount,
		})
		return
	}
	if err = json.Unmarshal(b, &vc); err == nil {
		*v = AccountsGetResponseBody(Accounts{
			Results:    []Account{vc},
			TotalCount: 1,
		})
		return
	}
	return
}

func (r *AccountsGet) URL() *url.URL {
	if r.PathParams().ID != "" {
		u := r.client.GetEndpointURL("/accounts/{{.id}}", r.PathParams())
		return &u
	}
	u := r.client.GetEndpointURL("/accounts", r.PathParams())
	return &u
}

func (r *AccountsGet) Do() (AccountsGetResponseBody, error) {
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
