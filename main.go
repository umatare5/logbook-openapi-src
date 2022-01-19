// Package logbook provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package logbook

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// GetDivelogsParams defines parameters for GetDivelogs.
type GetDivelogsParams struct {
	// The numbers of items to return
	Limit *int `json:"limit,omitempty"`

	// A unique identifier for a specific record
	Cursor *string `json:"cursor,omitempty"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetAdminHealth request
	GetAdminHealth(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetAdminTokens request
	GetAdminTokens(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetDivelog request
	GetDivelog(ctx context.Context, divelogId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetDivelogs request
	GetDivelogs(ctx context.Context, params *GetDivelogsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetAdminHealth(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAdminHealthRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetAdminTokens(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAdminTokensRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetDivelog(ctx context.Context, divelogId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetDivelogRequest(c.Server, divelogId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetDivelogs(ctx context.Context, params *GetDivelogsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetDivelogsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetAdminHealthRequest generates requests for GetAdminHealth
func NewGetAdminHealthRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/admin/health")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetAdminTokensRequest generates requests for GetAdminTokens
func NewGetAdminTokensRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/admin/tokens")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetDivelogRequest generates requests for GetDivelog
func NewGetDivelogRequest(server string, divelogId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "divelogId", runtime.ParamLocationPath, divelogId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/divelog/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetDivelogsRequest generates requests for GetDivelogs
func NewGetDivelogsRequest(server string, params *GetDivelogsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/divelogs")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Limit != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Cursor != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "cursor", runtime.ParamLocationQuery, *params.Cursor); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetAdminHealth request
	GetAdminHealthWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAdminHealthResponse, error)

	// GetAdminTokens request
	GetAdminTokensWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAdminTokensResponse, error)

	// GetDivelog request
	GetDivelogWithResponse(ctx context.Context, divelogId string, reqEditors ...RequestEditorFn) (*GetDivelogResponse, error)

	// GetDivelogs request
	GetDivelogsWithResponse(ctx context.Context, params *GetDivelogsParams, reqEditors ...RequestEditorFn) (*GetDivelogsResponse, error)
}

type GetAdminHealthResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Code    *string `json:"code,omitempty"`
		Message *string `json:"message,omitempty"`
	}
	JSONDefault *struct {
		// Error code
		Code *string `json:"code,omitempty"`

		// Error message
		Message *string `json:"message,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r GetAdminHealthResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAdminHealthResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAdminTokensResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Code    *string `json:"code,omitempty"`
		Message *string `json:"message,omitempty"`
		Tokens  *[]struct {
			AccessToken *string `json:"access_token,omitempty"`
			Expired     *bool   `json:"expired,omitempty"`
		} `json:"tokens,omitempty"`
	}
	JSONDefault *struct {
		// Error code
		Code *string `json:"code,omitempty"`

		// Error message
		Message *string `json:"message,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r GetAdminTokensResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAdminTokensResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetDivelogResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Code     *string `json:"code,omitempty"`
		Message  *string `json:"message,omitempty"`
		Response *struct {
			Divelog *struct {
				AirIn            *int     `json:"air_in,omitempty"`
				AirInText        *string  `json:"air_in_text,omitempty"`
				AirOut           *int     `json:"air_out,omitempty"`
				AirOutText       *string  `json:"air_out_text,omitempty"`
				AirTemperature   *string  `json:"air_temperature,omitempty"`
				Altitude         *string  `json:"altitude,omitempty"`
				AltitudeType     *int     `json:"altitude_type,omitempty"`
				AverageDepth     *string  `json:"average_depth,omitempty"`
				CommentCount     *int     `json:"comment_count,omitempty"`
				CreateDatetime   *float32 `json:"create_datetime,omitempty"`
				CurrentType      *int     `json:"current_type,omitempty"`
				CylinderSize     *int     `json:"cylinder_size,omitempty"`
				CylinderSizeText *string  `json:"cylinder_size_text,omitempty"`
				DeleteDatetime   *float32 `json:"delete_datetime,omitempty"`
				DeleteFlag       *bool    `json:"delete_flag,omitempty"`
				DeviceLogId      *float32 `json:"device_log_id,omitempty"`
				DiveBuddies      *string  `json:"dive_buddies,omitempty"`
				DiveDatetime     *int     `json:"dive_datetime,omitempty"`
				DiveDuration     *int     `json:"dive_duration,omitempty"`
				DiveProfiles     *[]struct {
					Direction   *float32 `json:"direction,omitempty"`
					Ndl         *float32 `json:"ndl,omitempty"`
					Pressure    *string  `json:"pressure,omitempty"`
					ProfileTime *float32 `json:"profile_time,omitempty"`
					Temperature *string  `json:"temperature,omitempty"`
				} `json:"dive_profiles,omitempty"`
				DiveType          *string `json:"dive_type,omitempty"`
				DivecomputerBrand *string `json:"divecomputer_brand,omitempty"`
				DivecomputerModel *string `json:"divecomputer_model,omitempty"`
				DivelogId         *string `json:"divelog_id,omitempty"`
				EntryPosition     *struct {
					Latitude  *float32 `json:"latitude,omitempty"`
					Longitude *float32 `json:"longitude,omitempty"`
				} `json:"entry_position,omitempty"`
				EventCount *float32 `json:"event_count,omitempty"`
				Events     *[]struct {
					EventCode  *float32 `json:"event_code,omitempty"`
					EventTime  *float32 `json:"event_time,omitempty"`
					EventType  *float32 `json:"event_type,omitempty"`
					EventValue *float32 `json:"event_value,omitempty"`
				} `json:"events,omitempty"`
				ExitPosition *struct {
					Latitude  *float32 `json:"latitude,omitempty"`
					Longitude *float32 `json:"longitude,omitempty"`
				} `json:"exit_position,omitempty"`
				FirmwareVersion  *string  `json:"firmware_version,omitempty"`
				FreeDepthAlarm1  *int     `json:"free_depth_alarm1,omitempty"`
				FreeDepthAlarm2  *int     `json:"free_depth_alarm2,omitempty"`
				FreeDepthAlarm3  *int     `json:"free_depth_alarm3,omitempty"`
				FreeDepthAlarm4  *int     `json:"free_depth_alarm4,omitempty"`
				FreeDepthAlarm5  *int     `json:"free_depth_alarm5,omitempty"`
				FreeDescentCount *int     `json:"free_descent_count,omitempty"`
				FreeTimeAlarm1   *int     `json:"free_time_alarm1,omitempty"`
				FreeTimeAlarm2   *int     `json:"free_time_alarm2,omitempty"`
				FreeTimeAlarm3   *int     `json:"free_time_alarm3,omitempty"`
				FreeTimeAlarm4   *int     `json:"free_time_alarm4,omitempty"`
				FreeTimeAlarm5   *int     `json:"free_time_alarm5,omitempty"`
				GaugeDepthAlarm  *int     `json:"gauge_depth_alarm,omitempty"`
				GaugeTimeAlarm   *int     `json:"gauge_time_alarm,omitempty"`
				GfSetting        *string  `json:"gf_setting,omitempty"`
				HardwareSerialNo *string  `json:"hardware_serial_no,omitempty"`
				Instructors      *string  `json:"instructors,omitempty"`
				IsLiked          *bool    `json:"is_liked,omitempty"`
				LikeCount        *int     `json:"like_count,omitempty"`
				LogInterval      *int     `json:"log_interval,omitempty"`
				LogNumber        *float32 `json:"log_number,omitempty"`
				MaxDepth         *string  `json:"max_depth,omitempty"`
				MaxPressure      *string  `json:"max_pressure,omitempty"`
				MediaCount       *int     `json:"media_count,omitempty"`
				Medias           *[]struct {
					MediaId   *string  `json:"media_id,omitempty"`
					MediaPath *string  `json:"media_path,omitempty"`
					MediaType *string  `json:"media_type,omitempty"`
					Sequence  *float32 `json:"sequence,omitempty"`
				} `json:"medias,omitempty"`
				MinTemperature *string  `json:"min_temperature,omitempty"`
				ModifyDatetime *float32 `json:"modify_datetime,omitempty"`
				NoFlightTime   *int     `json:"no_flight_time,omitempty"`
				Notes          *string  `json:"notes,omitempty"`
				Po2            *int     `json:"po2,omitempty"`
				Poi            *struct {
					CountryCode *string `json:"country_code,omitempty"`
					GpsLocation *struct {
						Latitude  *float32 `json:"latitude,omitempty"`
						Longitude *float32 `json:"longitude,omitempty"`
					} `json:"gps_location,omitempty"`
					LogCount  *float32 `json:"log_count,omitempty"`
					PoiId     *string  `json:"poi_id,omitempty"`
					PoiName   *string  `json:"poi_name,omitempty"`
					PoiRegion *string  `json:"poi_region,omitempty"`
					PoiType   *string  `json:"poi_type,omitempty"`
				} `json:"poi,omitempty"`
				Ppo2            *string `json:"ppo2,omitempty"`
				Privacy         *string `json:"privacy,omitempty"`
				PublishStatus   *string `json:"publish_status,omitempty"`
				ScubaDepthAlarm *int    `json:"scuba_depth_alarm,omitempty"`
				ScubaTimeAlarm  *int    `json:"scuba_time_alarm,omitempty"`
				ShareCount      *int    `json:"share_count,omitempty"`
				SuitThickness   *string `json:"suit_thickness,omitempty"`
				SurfaceInterval *int    `json:"surface_interval,omitempty"`
				Timezone        *int    `json:"timezone,omitempty"`
				Unit            *int    `json:"unit,omitempty"`
				User            *struct {
					CertificationId  *string `json:"certification_id,omitempty"`
					ProfileImagePath *string `json:"profile_image_path,omitempty"`
					UserId           *string `json:"user_id,omitempty"`
					UserName         *string `json:"user_name,omitempty"`
				} `json:"user,omitempty"`
				VisibilityType *int    `json:"visibility_type,omitempty"`
				WaterType      *int    `json:"water_type,omitempty"`
				WaveType       *int    `json:"wave_type,omitempty"`
				WeatherType    *int    `json:"weather_type,omitempty"`
				Weight         *string `json:"weight,omitempty"`
			} `json:"divelog,omitempty"`
		} `json:"response,omitempty"`
	}
	JSONDefault *struct {
		// Error code
		Code *string `json:"code,omitempty"`

		// Error message
		Message *string `json:"message,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r GetDivelogResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetDivelogResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetDivelogsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Code     *string `json:"code,omitempty"`
		Message  *string `json:"message,omitempty"`
		Response *struct {
			Divelogs *[]struct {
				DiveDatetime *int    `json:"dive_datetime,omitempty"`
				DiveDuration *int    `json:"dive_duration,omitempty"`
				DiveType     *string `json:"dive_type,omitempty"`
				DivelogId    *string `json:"divelog_id,omitempty"`
				MaxDepth     *string `json:"max_depth,omitempty"`
				MediaCount   *int    `json:"media_count,omitempty"`
				Poi          *struct {
					CountryCode *string `json:"country_code,omitempty"`
					GpsLocation *struct {
						Latitude  *float32 `json:"latitude,omitempty"`
						Longitude *float32 `json:"longitude,omitempty"`
					} `json:"gps_location,omitempty"`
					LogCount  *float32 `json:"log_count,omitempty"`
					PoiId     *string  `json:"poi_id,omitempty"`
					PoiName   *string  `json:"poi_name,omitempty"`
					PoiRegion *string  `json:"poi_region,omitempty"`
					PoiType   *string  `json:"poi_type,omitempty"`
				} `json:"poi,omitempty"`
				Privacy       *string `json:"privacy,omitempty"`
				PublishStatus *string `json:"publish_status,omitempty"`
				Timezone      *int    `json:"timezone,omitempty"`
				User          *struct {
					CertificationId  *string `json:"certification_id,omitempty"`
					ProfileImagePath *string `json:"profile_image_path,omitempty"`
					UserId           *string `json:"user_id,omitempty"`
					UserName         *string `json:"user_name,omitempty"`
				} `json:"user,omitempty"`
			} `json:"divelogs,omitempty"`
			PageInfo *struct {
				EndCursor   *string `json:"end_cursor,omitempty"`
				HasNextPage *bool   `json:"has_next_page,omitempty"`
			} `json:"page_info,omitempty"`
		} `json:"response,omitempty"`
	}
	JSONDefault *struct {
		// Error code
		Code *string `json:"code,omitempty"`

		// Error message
		Message *string `json:"message,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r GetDivelogsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetDivelogsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetAdminHealthWithResponse request returning *GetAdminHealthResponse
func (c *ClientWithResponses) GetAdminHealthWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAdminHealthResponse, error) {
	rsp, err := c.GetAdminHealth(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAdminHealthResponse(rsp)
}

// GetAdminTokensWithResponse request returning *GetAdminTokensResponse
func (c *ClientWithResponses) GetAdminTokensWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAdminTokensResponse, error) {
	rsp, err := c.GetAdminTokens(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAdminTokensResponse(rsp)
}

// GetDivelogWithResponse request returning *GetDivelogResponse
func (c *ClientWithResponses) GetDivelogWithResponse(ctx context.Context, divelogId string, reqEditors ...RequestEditorFn) (*GetDivelogResponse, error) {
	rsp, err := c.GetDivelog(ctx, divelogId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetDivelogResponse(rsp)
}

// GetDivelogsWithResponse request returning *GetDivelogsResponse
func (c *ClientWithResponses) GetDivelogsWithResponse(ctx context.Context, params *GetDivelogsParams, reqEditors ...RequestEditorFn) (*GetDivelogsResponse, error) {
	rsp, err := c.GetDivelogs(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetDivelogsResponse(rsp)
}

// ParseGetAdminHealthResponse parses an HTTP response from a GetAdminHealthWithResponse call
func ParseGetAdminHealthResponse(rsp *http.Response) (*GetAdminHealthResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAdminHealthResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest struct {
			// Error code
			Code *string `json:"code,omitempty"`

			// Error message
			Message *string `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseGetAdminTokensResponse parses an HTTP response from a GetAdminTokensWithResponse call
func ParseGetAdminTokensResponse(rsp *http.Response) (*GetAdminTokensResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAdminTokensResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
			Tokens  *[]struct {
				AccessToken *string `json:"access_token,omitempty"`
				Expired     *bool   `json:"expired,omitempty"`
			} `json:"tokens,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest struct {
			// Error code
			Code *string `json:"code,omitempty"`

			// Error message
			Message *string `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseGetDivelogResponse parses an HTTP response from a GetDivelogWithResponse call
func ParseGetDivelogResponse(rsp *http.Response) (*GetDivelogResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetDivelogResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Code     *string `json:"code,omitempty"`
			Message  *string `json:"message,omitempty"`
			Response *struct {
				Divelog *struct {
					AirIn            *int     `json:"air_in,omitempty"`
					AirInText        *string  `json:"air_in_text,omitempty"`
					AirOut           *int     `json:"air_out,omitempty"`
					AirOutText       *string  `json:"air_out_text,omitempty"`
					AirTemperature   *string  `json:"air_temperature,omitempty"`
					Altitude         *string  `json:"altitude,omitempty"`
					AltitudeType     *int     `json:"altitude_type,omitempty"`
					AverageDepth     *string  `json:"average_depth,omitempty"`
					CommentCount     *int     `json:"comment_count,omitempty"`
					CreateDatetime   *float32 `json:"create_datetime,omitempty"`
					CurrentType      *int     `json:"current_type,omitempty"`
					CylinderSize     *int     `json:"cylinder_size,omitempty"`
					CylinderSizeText *string  `json:"cylinder_size_text,omitempty"`
					DeleteDatetime   *float32 `json:"delete_datetime,omitempty"`
					DeleteFlag       *bool    `json:"delete_flag,omitempty"`
					DeviceLogId      *float32 `json:"device_log_id,omitempty"`
					DiveBuddies      *string  `json:"dive_buddies,omitempty"`
					DiveDatetime     *int     `json:"dive_datetime,omitempty"`
					DiveDuration     *int     `json:"dive_duration,omitempty"`
					DiveProfiles     *[]struct {
						Direction   *float32 `json:"direction,omitempty"`
						Ndl         *float32 `json:"ndl,omitempty"`
						Pressure    *string  `json:"pressure,omitempty"`
						ProfileTime *float32 `json:"profile_time,omitempty"`
						Temperature *string  `json:"temperature,omitempty"`
					} `json:"dive_profiles,omitempty"`
					DiveType          *string `json:"dive_type,omitempty"`
					DivecomputerBrand *string `json:"divecomputer_brand,omitempty"`
					DivecomputerModel *string `json:"divecomputer_model,omitempty"`
					DivelogId         *string `json:"divelog_id,omitempty"`
					EntryPosition     *struct {
						Latitude  *float32 `json:"latitude,omitempty"`
						Longitude *float32 `json:"longitude,omitempty"`
					} `json:"entry_position,omitempty"`
					EventCount *float32 `json:"event_count,omitempty"`
					Events     *[]struct {
						EventCode  *float32 `json:"event_code,omitempty"`
						EventTime  *float32 `json:"event_time,omitempty"`
						EventType  *float32 `json:"event_type,omitempty"`
						EventValue *float32 `json:"event_value,omitempty"`
					} `json:"events,omitempty"`
					ExitPosition *struct {
						Latitude  *float32 `json:"latitude,omitempty"`
						Longitude *float32 `json:"longitude,omitempty"`
					} `json:"exit_position,omitempty"`
					FirmwareVersion  *string  `json:"firmware_version,omitempty"`
					FreeDepthAlarm1  *int     `json:"free_depth_alarm1,omitempty"`
					FreeDepthAlarm2  *int     `json:"free_depth_alarm2,omitempty"`
					FreeDepthAlarm3  *int     `json:"free_depth_alarm3,omitempty"`
					FreeDepthAlarm4  *int     `json:"free_depth_alarm4,omitempty"`
					FreeDepthAlarm5  *int     `json:"free_depth_alarm5,omitempty"`
					FreeDescentCount *int     `json:"free_descent_count,omitempty"`
					FreeTimeAlarm1   *int     `json:"free_time_alarm1,omitempty"`
					FreeTimeAlarm2   *int     `json:"free_time_alarm2,omitempty"`
					FreeTimeAlarm3   *int     `json:"free_time_alarm3,omitempty"`
					FreeTimeAlarm4   *int     `json:"free_time_alarm4,omitempty"`
					FreeTimeAlarm5   *int     `json:"free_time_alarm5,omitempty"`
					GaugeDepthAlarm  *int     `json:"gauge_depth_alarm,omitempty"`
					GaugeTimeAlarm   *int     `json:"gauge_time_alarm,omitempty"`
					GfSetting        *string  `json:"gf_setting,omitempty"`
					HardwareSerialNo *string  `json:"hardware_serial_no,omitempty"`
					Instructors      *string  `json:"instructors,omitempty"`
					IsLiked          *bool    `json:"is_liked,omitempty"`
					LikeCount        *int     `json:"like_count,omitempty"`
					LogInterval      *int     `json:"log_interval,omitempty"`
					LogNumber        *float32 `json:"log_number,omitempty"`
					MaxDepth         *string  `json:"max_depth,omitempty"`
					MaxPressure      *string  `json:"max_pressure,omitempty"`
					MediaCount       *int     `json:"media_count,omitempty"`
					Medias           *[]struct {
						MediaId   *string  `json:"media_id,omitempty"`
						MediaPath *string  `json:"media_path,omitempty"`
						MediaType *string  `json:"media_type,omitempty"`
						Sequence  *float32 `json:"sequence,omitempty"`
					} `json:"medias,omitempty"`
					MinTemperature *string  `json:"min_temperature,omitempty"`
					ModifyDatetime *float32 `json:"modify_datetime,omitempty"`
					NoFlightTime   *int     `json:"no_flight_time,omitempty"`
					Notes          *string  `json:"notes,omitempty"`
					Po2            *int     `json:"po2,omitempty"`
					Poi            *struct {
						CountryCode *string `json:"country_code,omitempty"`
						GpsLocation *struct {
							Latitude  *float32 `json:"latitude,omitempty"`
							Longitude *float32 `json:"longitude,omitempty"`
						} `json:"gps_location,omitempty"`
						LogCount  *float32 `json:"log_count,omitempty"`
						PoiId     *string  `json:"poi_id,omitempty"`
						PoiName   *string  `json:"poi_name,omitempty"`
						PoiRegion *string  `json:"poi_region,omitempty"`
						PoiType   *string  `json:"poi_type,omitempty"`
					} `json:"poi,omitempty"`
					Ppo2            *string `json:"ppo2,omitempty"`
					Privacy         *string `json:"privacy,omitempty"`
					PublishStatus   *string `json:"publish_status,omitempty"`
					ScubaDepthAlarm *int    `json:"scuba_depth_alarm,omitempty"`
					ScubaTimeAlarm  *int    `json:"scuba_time_alarm,omitempty"`
					ShareCount      *int    `json:"share_count,omitempty"`
					SuitThickness   *string `json:"suit_thickness,omitempty"`
					SurfaceInterval *int    `json:"surface_interval,omitempty"`
					Timezone        *int    `json:"timezone,omitempty"`
					Unit            *int    `json:"unit,omitempty"`
					User            *struct {
						CertificationId  *string `json:"certification_id,omitempty"`
						ProfileImagePath *string `json:"profile_image_path,omitempty"`
						UserId           *string `json:"user_id,omitempty"`
						UserName         *string `json:"user_name,omitempty"`
					} `json:"user,omitempty"`
					VisibilityType *int    `json:"visibility_type,omitempty"`
					WaterType      *int    `json:"water_type,omitempty"`
					WaveType       *int    `json:"wave_type,omitempty"`
					WeatherType    *int    `json:"weather_type,omitempty"`
					Weight         *string `json:"weight,omitempty"`
				} `json:"divelog,omitempty"`
			} `json:"response,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest struct {
			// Error code
			Code *string `json:"code,omitempty"`

			// Error message
			Message *string `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseGetDivelogsResponse parses an HTTP response from a GetDivelogsWithResponse call
func ParseGetDivelogsResponse(rsp *http.Response) (*GetDivelogsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetDivelogsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Code     *string `json:"code,omitempty"`
			Message  *string `json:"message,omitempty"`
			Response *struct {
				Divelogs *[]struct {
					DiveDatetime *int    `json:"dive_datetime,omitempty"`
					DiveDuration *int    `json:"dive_duration,omitempty"`
					DiveType     *string `json:"dive_type,omitempty"`
					DivelogId    *string `json:"divelog_id,omitempty"`
					MaxDepth     *string `json:"max_depth,omitempty"`
					MediaCount   *int    `json:"media_count,omitempty"`
					Poi          *struct {
						CountryCode *string `json:"country_code,omitempty"`
						GpsLocation *struct {
							Latitude  *float32 `json:"latitude,omitempty"`
							Longitude *float32 `json:"longitude,omitempty"`
						} `json:"gps_location,omitempty"`
						LogCount  *float32 `json:"log_count,omitempty"`
						PoiId     *string  `json:"poi_id,omitempty"`
						PoiName   *string  `json:"poi_name,omitempty"`
						PoiRegion *string  `json:"poi_region,omitempty"`
						PoiType   *string  `json:"poi_type,omitempty"`
					} `json:"poi,omitempty"`
					Privacy       *string `json:"privacy,omitempty"`
					PublishStatus *string `json:"publish_status,omitempty"`
					Timezone      *int    `json:"timezone,omitempty"`
					User          *struct {
						CertificationId  *string `json:"certification_id,omitempty"`
						ProfileImagePath *string `json:"profile_image_path,omitempty"`
						UserId           *string `json:"user_id,omitempty"`
						UserName         *string `json:"user_name,omitempty"`
					} `json:"user,omitempty"`
				} `json:"divelogs,omitempty"`
				PageInfo *struct {
					EndCursor   *string `json:"end_cursor,omitempty"`
					HasNextPage *bool   `json:"has_next_page,omitempty"`
				} `json:"page_info,omitempty"`
			} `json:"response,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest struct {
			// Error code
			Code *string `json:"code,omitempty"`

			// Error message
			Message *string `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /admin/health)
	GetAdminHealth(ctx echo.Context) error

	// (GET /admin/tokens)
	GetAdminTokens(ctx echo.Context) error

	// (GET /divelog/{divelogId})
	GetDivelog(ctx echo.Context, divelogId string) error

	// (GET /divelogs)
	GetDivelogs(ctx echo.Context, params GetDivelogsParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetAdminHealth converts echo context to params.
func (w *ServerInterfaceWrapper) GetAdminHealth(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAdminHealth(ctx)
	return err
}

// GetAdminTokens converts echo context to params.
func (w *ServerInterfaceWrapper) GetAdminTokens(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAdminTokens(ctx)
	return err
}

// GetDivelog converts echo context to params.
func (w *ServerInterfaceWrapper) GetDivelog(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "divelogId" -------------
	var divelogId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "divelogId", runtime.ParamLocationPath, ctx.Param("divelogId"), &divelogId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter divelogId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDivelog(ctx, divelogId)
	return err
}

// GetDivelogs converts echo context to params.
func (w *ServerInterfaceWrapper) GetDivelogs(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetDivelogsParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "cursor" -------------

	err = runtime.BindQueryParameter("form", true, false, "cursor", ctx.QueryParams(), &params.Cursor)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter cursor: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDivelogs(ctx, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/admin/health", wrapper.GetAdminHealth)
	router.GET(baseURL+"/admin/tokens", wrapper.GetAdminTokens)
	router.GET(baseURL+"/divelog/:divelogId", wrapper.GetDivelog)
	router.GET(baseURL+"/divelogs", wrapper.GetDivelogs)

}
