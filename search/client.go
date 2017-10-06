package search

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"io"
)

const (
	// EndpointProduction is the Search production endpoint
	EndpointProduction = "https://cloudplatform.coveo.com/rest/search/"
	// EndpointStaging is the Search staging endpoint
	EndpointStaging = "https://cloudplatformstaging.coveo.com/rest/search/"
	// EndpointDevelopment is the Search development endpoint
	EndpointDevelopment = "https://cloudplatformdev.coveo.com/rest/search/"
)

// Client is the search client to make search requests
type Client interface {
	Query(q Query) (*Response, error)
	ListFacetValues(field string, maximumNumberOfValues int) (*FacetValues, error)
}

// Config is used to configure a new client
type Config struct {
	Token          string
	UserAgent      string
	OrganizationId string
	// Endpoint is used if you want to use custom endpoints (dev,staging,testing)
	Endpoint string
}

// NewClient returns a configured http search client using default http client
func NewClient(c Config) (Client, error) {
	if len(c.Endpoint) == 0 {
		c.Endpoint = EndpointProduction
	}

	return &client{
		token:      c.Token,
		endpoint:   c.Endpoint,
		orgId:      c.OrganizationId,
		httpClient: http.DefaultClient,
		useragent:  c.UserAgent,
	}, nil
}

type client struct {
	httpClient *http.Client
	token      string
	endpoint   string
	orgId      string
	useragent  string
}

func (c *client) Query(q Query) (*Response, error) {
	marshalledQuery, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewReader(marshalledQuery)

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accepts":      "application/json",
	}

	req, err := createRequest(c, "POST", map[string]string{}, headers, c.endpoint, buf)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	queryResponse := &Response{}
	err = json.NewDecoder(resp.Body).Decode(queryResponse)
	return queryResponse, err
}

func (c *client) ListFacetValues(field string, maximumNumberOfValues int) (*FacetValues, error) {
	url, err := url.Parse(c.endpoint + "/values")
	if err != nil {
		return nil, err
	}

	queryParams := map[string]string{
		"field":                 field,
		"maximumNumberOfValues": strconv.Itoa(maximumNumberOfValues),
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accepts":      "application/json",
	}

	req, err := createRequest(c, "GET", queryParams, headers, url.String(), nil)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	facetValues := &FacetValues{}
	err = json.NewDecoder(resp.Body).Decode(facetValues)
	return facetValues, nil
}

func createRequest(client *client, reqType string, queryStringParams map[string]string, customHeaders map[string]string, endpoint string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(reqType, endpoint, body)
	if err != nil {
		return nil, err
	}

	customHeaders["Authorization"] = "Bearer " + client.token
	customHeaders["User-Agent"] = client.useragent
	for key, value := range customHeaders {
		req.Header.Add(key, value)
	}

	queryStringParams["organizationId"] = client.orgId
	urlQuery := req.URL.Query()
	for key, value := range queryStringParams {
		urlQuery.Add(key, value)
	}
	req.URL.RawQuery = urlQuery.Encode()

	return req, err
}
