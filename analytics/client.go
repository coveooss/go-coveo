package analytics

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client interface {
	SendInterfaceLoad(InterfaceLoadEvent) ([]*http.Cookie, error)
	SendSearchEvent(SearchEvent) (*SearchEventResponse, error)
	SendSearchEventWithCookie(event SearchEvent, cookie *http.Cookie) (string, error)
	SendSearchesEvent([]SearchEvent) (*SearchEventsResponse, error)
	SendClickEvent(ClickEvent) (*ClickEventResponse, error)
	SendCustomEvent(CustomEvent) (*CustomEventResponse, error)
	GetVisit() (*VisitResponse, error)
	GetStatus() (*StatusResponse, error)
	DeleteVisit() (bool, error)
}

type Config struct {
	Token     string
	UserAgent string
}

func NewClient(c Config) (Client, error) {
	return &client{
		token:      c.Token,
		endpoint:   "https://usageanalytics.coveo.com/rest/v14/analytics/",
		httpClient: http.DefaultClient,
		useragent:  c.UserAgent,
	}, nil
}

type client struct {
	httpClient *http.Client
	token      string
	endpoint   string
	useragent  string
}

type InterfaceLoadResponse struct{}
type StatusResponse struct{}
type SearchEventResponse struct {
	SearchUID string `json:"searchEventUid"`
	VisitID   string `json:"visitId"`
}
type SearchEventsResponse struct{}
type ClickEventResponse struct{}
type CustomEventResponse struct{}
type VisitResponse struct{}

/*
*	Send the Interface load event to the usage analytics endpoint.
*	This needs to be called before any other event sent to UA.
*
*	Will return the setCookie and the potential errors
 */
func (c *client) SendInterfaceLoad(event InterfaceLoadEvent) ([]*http.Cookie, error) {
	marshalledEvent, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewReader(marshalledEvent)

	req, err := http.NewRequest("POST", c.endpoint+"search/", buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accepts", "application/json")
	req.Header.Set("User-Agent", c.useragent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cookies := resp.Cookies()

	return cookies, err
}

func (c *client) SendSearchEvent(event SearchEvent) (*SearchEventResponse, error) {
	marshalledEvent, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewReader(marshalledEvent)

	req, err := http.NewRequest("POST", c.endpoint+"search/", buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accepts", "application/json")
	req.Header.Set("User-Agent", c.useragent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	queryResponse := &SearchEventResponse{}
	err = json.NewDecoder(resp.Body).Decode(queryResponse)
	return queryResponse, err
}
func (c *client) SendSearchEventWithCookie(event SearchEvent, cookie *http.Cookie) (string, error) {
	marshalledEvent, err := json.Marshal(event)
	if err != nil {
		return "", err
	}
	buf := bytes.NewReader(marshalledEvent)

	req, err := http.NewRequest("POST", c.endpoint+"search/", buf)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accepts", "application/json")
	req.Header.Set("User-Agent", c.useragent)

	req.AddCookie(cookie)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)

	return string(contents), err
}
func (c *client) SendSearchesEvent(event []SearchEvent) (*SearchEventsResponse, error) {
	return nil, nil
}
func (c *client) SendClickEvent(event ClickEvent) (*ClickEventResponse, error) {
	return nil, nil
}
func (c *client) SendCustomEvent(event CustomEvent) (*CustomEventResponse, error) {
	return nil, nil
}
func (c *client) GetVisit() (*VisitResponse, error) {
	return nil, nil
}
func (c *client) DeleteVisit() (bool, error) {
	return false, nil
}
func (c *client) GetStatus() (*StatusResponse, error) {
	return nil, nil
}