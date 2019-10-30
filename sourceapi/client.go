package sourceapi

import (
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	// EndpointProduction is the push production endpoint
	EndpointProduction = "https://platform.cloud.coveo.com/rest/organizations/"
)

// Client is the pushapi client to send documents or identities
type Client interface {
	CreateSource(d Source) (string, error) 
	UpdateSource(sourceID string, d Source) (string, error)
	DeleteSource(sourceID string) error
	ReadSource(sourceID string) (string, error)
}

type client struct {
	httpClient     *http.Client
	apikey         string
	endpoint       string
	organizationid string
}

// Config is used to configure a new client
type Config struct {
	// Endpoint is used if you want to use custom endpoints (dev,staging,testing)
	Endpoint string
	// The Coveo organization ID
	OrganizationID string
	// APIKey is the key used to push content to Coveo
	APIKey string
}

// NewClient initializes a new pushapi client with the config param
func NewClient(c Config) (Client, error) {
	if len(c.Endpoint) == 0 {
		c.Endpoint = EndpointProduction
	}

	return &client{
		apikey:         c.APIKey,
		endpoint:       c.Endpoint,
		organizationid: c.OrganizationID,
		httpClient:     http.DefaultClient,
	}, nil
}

func (c *client) sendRequest(req *http.Request) (string, string,  error) {
	req.Header.Add("Authorization", "Bearer "+c.apikey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "",  err
	}

	if resp.StatusCode >= 300  {
		return "", resp.Status, errors.New(string(body))
	}

	return string(body), resp.Status, nil
}
