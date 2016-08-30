package pushapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// EndpointProduction is the push production endpoint
	EndpointProduction = "https://push.cloud.coveo.com/v1/organizations/"
)

// Client is the pushapi client to send documents or identities
type Client interface {
	PushDocument(d Document, sourceID string) (string, error)
	DeleteDocument(documentID string, sourceID string) error
	PushIdentity(i Identity, providerID string) error
	DeleteIdentity(i Identity, providerID string) error
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

type client struct {
	httpClient     *http.Client
	apikey         string
	endpoint       string
	organizationid string
}

// PushDocument will send a document to the pushapi in the specified source
func (c *client) PushDocument(d Document, sourceID string) (string, error) {
	if len(sourceID) == 0 {
		return "", errors.New("You need a sourceID")
	}

	if len(d.DocumentID) == 0 {
		return "", errors.New("You need to provide a documentID")
	}

	marshalledDocument, err := json.Marshal(d.Fields)
	if err != nil {
		return "", err
	}
	buf := bytes.NewReader(marshalledDocument)

	endpoint := fmt.Sprintf("%s%s/sources/%s/documents?documentId=%s",
		c.endpoint, c.organizationid, sourceID, d.DocumentID)

	req, err := http.NewRequest("PUT", endpoint, buf)
	resp, err := c.sendRequest(req)
	if err != nil {
		return "", err
	}

	return resp, nil
}

// DeleteDocument will send a delete request for the specified documentID in the sourceID
func (c *client) DeleteDocument(documentID, sourceID string) error {
	if len(sourceID) == 0 {
		return errors.New("You need a sourceID")
	}

	if len(documentID) == 0 {
		return errors.New("You need a documentID")
	}

	endpoint := fmt.Sprintf("%s%s/sources/%s/documents?documentId=%s",
		c.endpoint, c.organizationid, sourceID, documentID)

	req, err := http.NewRequest("DELETE", endpoint, nil)
	_, err = c.sendRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) PushIdentity(i Identity, providerID string) error {
	return nil
}

func (c *client) DeleteIdentity(i Identity, providerID string) error {
	return nil
}

func (c *client) sendRequest(req *http.Request) (string, error) {
	req.Header.Add("Authorization", "Bearer "+c.apikey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusAccepted {
		return "", errors.New(string(body))
	}

	return string(body), nil
}
