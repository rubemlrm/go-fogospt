package fogospt

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const defaultBaseURL = "https://api.fogos.pt"

// Client act's as the entry object for sdk
type Client struct {
	BaseURL  *url.URL
	Incident *IncidentService
	client   *http.Client
	common   service
}

type service struct {
	client *Client
}

// NewClient creates a new http client instance in case the provided one is nil
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	client := &Client{BaseURL: baseURL, client: httpClient}
	client.common.client = client
	client.Incident = (*IncidentService)(&client.common)
	return client
}

// NewRequest handles all the requests to fogospt api and returns the response
// object with the given interface
func (c *Client) NewRequest(endpoint string, parseResponse interface{}) (*http.Response, error) {
	composedURL := fmt.Sprintf("%v%v", c.BaseURL, endpoint)
	request, err := http.NewRequest(http.MethodGet, composedURL, nil)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 500 {
		return nil, errors.New("Server error")
	}

	decondingResponse := json.NewDecoder(response.Body).Decode(parseResponse)
	if decondingResponse == io.EOF {
		decondingResponse = nil
	}
	if decondingResponse != nil {
		return nil, decondingResponse
	}

	return response, nil
}
