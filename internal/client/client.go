package client

import (
	"errors"
	"net/http"
)

var Client http.Client

func Do(url string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	response, err := Client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 500 {
		return nil, errors.New("Server error")
	}
	return response, nil
}
