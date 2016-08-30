package pushapi

// Document represents the structure of a document
type Document struct {
	DocumentID string                 `json:"documentId"`
	Fields     map[string]interface{} `json:"body"`
}
