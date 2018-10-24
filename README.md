go-coveo
========

[![Build Status](https://travis-ci.org/coveo/go-coveo.svg)](https://travis-ci.org/coveo/go-coveo)
[![Go Report](https://goreportcard.com/badge/github.com/coveo/go-coveo)](https://goreportcard.com/report/github.com/coveo/go-coveo)

go-coveo is a Go client library for accessing the [Coveo Search API](https://docs.coveo.com/en/13/cloud-v2-api-reference/search-api) and the [Coveo Usage Analytics API](https://docs.coveo.com/en/18/cloud-v2-api-reference/usage-analytics-write-api)


# Analytics client documentation

https://godoc.org/github.com/coveo/go-coveo/analytics

## Example usage
```Go
import "github.com/coveo/go-coveo/analytics"

uaConfig := analytics.Config {
    Token: "My_Token", 
    UserAgent: "Some UserAgent", 
    IP: "Some IP", 
    Endpoint: "https://my.analytics.endpoint.com"
}
uaClient := analytics.NewClient(uaConfig)
searchEvent := analytics.NewSearchEvent()
searchEvent.SearchQueryUID = "myQueryUID"
if err := uaClient.SendSearchEvent(searchEvent); err != nil {
    // Error
}
...
```


# Search client documentation

https://godoc.org/github.com/coveo/go-coveo/search

## Example usage

```Go
import "github.com/coveo/go-coveo/search"

searchConfig := search.Config {
    Token: "My_Token", 
    UserAgent: "Some UserAgent", 
    Endpoint: "https://my.endpoint.com"
}
searchClient, err := search.NewClient(searchConfig)
if response, err = searchClient.Query(myQuery); err != nil {
    // Error
}
...
```