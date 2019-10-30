package sourceapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Source struct {
	Id string 			`json:"id"`
	Name string			`json:"name"`
	Type string 		`json:"sourceType"`
	Visibility string 	`json:"sourceVisibility"`
	Enabled bool		`json:"pushEnabled"`
}


// PushSource will send a document to the pushapi to create a source
func (c *client) CreateSource(d Source) (string, error) {
	log.Printf("[INFO] Creating a resource")
	if len(d.Name) == 0 {
		return "", errors.New("you need a name for the source")
	}

	marshalledDocument, err := json.Marshal(d)
	log.Printf("[INFO] Document to create %s", string(marshalledDocument))
	if err != nil {
		return "", err
	}
	buf := bytes.NewReader(marshalledDocument)

	endpoint := fmt.Sprintf("%s%s/sources",
		c.endpoint, c.organizationid)
	log.Printf("[INFO] Contacting endpoint %s", endpoint)

	req, err := http.NewRequest("POST", endpoint, buf)
	resp, status,  err := c.sendRequest(req)
	if err != nil {
		return "", errors.New("received status:" + status)
	}

	return resp, nil
}

// ReadSource will send a GET request for the specified sourceID to read
func (c *client) ReadSource(sourceID string) (string, error) {
	if len(sourceID) == 0 {
		return "", errors.New("you need a sourceID")
	}

	endpoint := fmt.Sprintf("%s%s/sources/%s",
		c.endpoint, c.organizationid, sourceID)

	req, err := http.NewRequest("GET", endpoint, nil)
	resp, status, err := c.sendRequest(req)
	if err != nil {
		return "", errors.New("received status:" + status)
	}

	return resp, nil
}


// UpdateSource will send a PUT request for the specified sourceID to modify
func (c *client) UpdateSource(sourceID string, d Source) (string, error) {
	if len(sourceID) == 0 {
		return "", errors.New("you need a sourceID")
	}
	marshalledDocument, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	buf := bytes.NewReader(marshalledDocument)

	endpoint := fmt.Sprintf("%s%s/sources/%s",
		c.endpoint, c.organizationid, sourceID)

	req, err := http.NewRequest("PUT", endpoint, buf)
	resp, status, err := c.sendRequest(req)
	if err != nil {
		return "", errors.New("received status:" + status)
	}

	return resp, nil
}

// DeleteSource will send a delete request for the specified sourceID
func (c *client) DeleteSource(sourceID string) error {
	if len(sourceID) == 0 {
		return errors.New("You need a sourceID")
	}

	endpoint := fmt.Sprintf("%s%s/sources/%s",
		c.endpoint, c.organizationid, sourceID)

	req, err := http.NewRequest("DELETE", endpoint, nil)
	_, status, err := c.sendRequest(req)
	if err != nil {
		return errors.New("received status:" + status)
	}

	return nil
}
