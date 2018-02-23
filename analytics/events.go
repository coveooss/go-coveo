package analytics

// ActionEvent Basic incomplete analytics event type, contains most of the information
// sent from a search ui to the analytics
type ActionEvent struct {
	// Language The language that end user is using. Use ISO 639-1 codes. See ISO 639-1 for more information
	Language string `json:"language"`
	// UserAgent The user agent string of the end user's browser
	UserAgent string `json:"userAgent,omitempty"`
	// Device (DEPRECATED) The name of the device that the end user is using
	Device string `json:"device,omitempty"`
	// Custom data that can contain all the user defined dimensions and their values. Keys can only contain alphanumeric or underscore characters. Spaces in the key are converted to underscores. Uppercase characters in the key are converted to lowercase characters.
	CustomData map[string]interface{} `json:"customData,omitempty"`
	// Anonymous Indicates if the user is anonymous.
	Anonymous bool `json:"anonymous,omitempty"`
	// Username The unique identifier of the user performing the search event. If not provided, the username will be extracted from the token.
	Username string `json:"username,omitempty"`
	// UserDisplayName The display name of the user performing the search event.
	UserDisplayName string `json:"userDisplayName,omitempty"`
	// SplitTestRunName The name of the A/B test run, if one is active.
	SplitTestRunName string `json:"splitTestRunName,omitempty"`
	// SplitTestRunVersion The version of the A/B test run, if one is active.
	SplitTestRunVersion string `json:"splitTestRunVersion,omitempty"`
	// OriginLevel1 The origin this event took place.
	OriginLevel1 string `json:"originLevel1"`
	// OriginLevel2 The origin this event took place.
	OriginLevel2 string `json:"originLevel2"`
	// OriginLevel3 The origin this event took place
	OriginLevel3 string `json:"originLevel3,omitempty"`
	// Outcome The outcome of the event. A score you apply on the event which gives it importance
	Outcome int `json:"outcome,omitempty"`
	// OriginContext The origin of the event. Used to specify the section of the website in which the user performs the action.
	OriginContext string `json:"originContext,omitempty"`
	// Mobile (DEPRECATED) Whether the end user's device is a mobile device or not
	Mobile bool `json:"mobile,omitempty"`
}

// SearchEvent Is a structure reprensenting a search event sent to the analytics
// It incorporate an ActionEvent and adds more fields.
type SearchEvent struct {
	*ActionEvent
	// SearchQueryUID The unique identifier of the search event. Must be a UUID
	SearchQueryUID string `json:"searchQueryUid"`
	// QueryText The text that was searched
	QueryText string `json:"queryText"`
	// ActionCause The type of operation that triggered this event.
	ActionCause string `json:"actionCause"`
	// AdvancedQuery The advanced part of the query that was sent to the index
	AdvancedQuery string `json:"advancedQuery,omitempty"`
	// NumberOfResults The number of results found. Must be equal to or greater than zero
	NumberOfResults int `json:"numberOfResults,omitempty"`
	// Contextual Indicates if the query is modified by contextual filters (for instance, a query to find similar documents)
	Contextual bool `json:"contextual"`
	// ResponseTime The time, in milliseconds, between the moment the query was sent and the moment the results were received. Must be greater than zero
	ResponseTime int `json:"responseTime,omitempty"`
	// Results The results of this search event
	Results []ResultHash `json:"results,omitempty"`
	// QueryPipeline The query pipeline of the search event.
	QueryPipeline string `json:"queryPipeline,omitempty"`
	// UserGroups The groups that the end user performing the event is a member of
	UserGroups []string `json:"userGroups,omitempty"`
}

// ResultHash Is a type used by the analytics to describe a result that was
// returned by a query that is usually sent with a search event.
type ResultHash struct {
	DocumentURI     string `json:"documentUri"`
	DocumentURIHash string `json:"documentUriHash"`
}

// ClickEvent Is a structure reprensenting a click event sent to the analytics
// It incorporate an ActionEvent and adds more fields.
type ClickEvent struct {
	*ActionEvent
	// DocumentURI The @sysuri of the document that was clicked
	DocumentURI string `json:"documentUri"`
	// DocumentURIHash The @sysurihash of the document that was clicked
	DocumentURIHash string `json:"documentUriHash"`
	// SearchQueryUID The searchQueryUid of the search that returned the document that was clicked
	SearchQueryUID string `json:"searchQueryUid"`
	// CollectionName The @syscollection of the document that was clicked
	CollectionName string `json:"collectionName"`
	// SourceName The @syssource of the document that was clicked
	SourceName string `json:"sourceName"`
	// DocumentPosition The position of the document (one-based) that was clicked in the list of results
	DocumentPosition int `json:"documentPosition"`
	// ActionCause The type of operation that triggered this event. Example: 'firstSearch', 'searchEnter'
	ActionCause string `json:"actionCause"`
	// DocumentTitle The title of the document that was clicked
	DocumentTitle string `json:"documentTitle,omitempty"`
	// DocumentURL The URL of the document that was clicked.
	DocumentURL string `json:"documentUrl,omitempty"`
	// DocumentAuthor The author of the document that was clicked.
	DocumentAuthor string `json:"documentAuthor,omitempty"`
	// QueryPipeline The query pipeline of the search that returned the document that was clicked.
	QueryPipeline string `json:"queryPipeline,omitempty"`
	// RankingModifier The ranking modifier that affected the clicked document.
	RankingModifier string `json:"rankingModifier,omitempty"`
	// DocumentCategory The category of the clicked document.
	DocumentCategory string `json:"documentCategory,omitempty"`
}

// CustomEvent Is a structure reprensenting a custom event sent to the analytics
// It incorporate an ActionEvent and adds more fields.
type CustomEvent struct {
	*ActionEvent
	// EventType The type of this event
	EventType string `json:"eventType"`
	// EventValue The type of operation that triggered this event
	EventValue string `json:"eventValue"`
	// LastSearchQueryUID The searchQueryUid of the last search event that occurred before this event
	LastSearchQueryUID string `json:"lastSearchQueryUid,omitempty"`
}

// ViewEvent Is a structure reprensenting a view event sent to the analytics
// It incorporate an ActionEvent and adds more fields.
type ViewEvent struct {
	*ActionEvent
	// Location The uri of the loaded page or component that is viewed
	Location string `json:"location"`
	// Referrer The page referrer
	Referrer string `json:"referrer,omitempty"`
	// Title The title of the page
	Title string `json:"title,omitempty"`
	// ContentIDKey The content ID key to match with the value
	ContentIDKey string `json:"contentIdKey,omitempty"`
	// ContentIDValue The content ID value to match with the key
	ContentIDValue string `json:"contentIdValue,omitempty"`
	// ContentType The type of the content
	ContentType string `json:"contentType,omitempty"`
}
