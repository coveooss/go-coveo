package sourceapi

import (
	"testing"
)

func TestClient(t *testing.T) {
	config := Config{}
	_, err := NewClient(config)
	if err != nil {
		t.Fatalf("unexpected error.  expected %v, actual %v", nil, err)
	}
}

func TestCreateSourceReturnError(t *testing.T){
	config := Config{
		OrganizationID : "<organization_id>",
		APIKey : "<api_key>",
	}
	c, err := NewClient(config)
	source := Source{
		Name:       "test-qa-enviroment-source",
		Type:       "PUSH",
		Visibility: "SHARED",
		Enabled:    true,
	}
	_, err = c.CreateSource(source)
	if err == nil {
		t.Fatalf("Error shouldn't be empty")
	}
}

