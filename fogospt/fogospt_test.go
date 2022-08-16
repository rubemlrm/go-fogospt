package fogospt_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/Rubemlrm/fogospt-golang-client/fogospt"
	"github.com/stretchr/testify/assert"
)

type dummyStruct struct{}

func TestClientCallWithSuccess(t *testing.T) {
	// act
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer server.Close()
	c := fogospt.NewClient(nil)
	c.BaseURL, _ = url.Parse(server.URL)
	// want
	dummyResponse := new(dummyStruct)
	req, err := c.NewRequest("/endpoint", dummyResponse)
	// assert
	assert.NoError(t, err)
	assert.Equal(t, 200, req.StatusCode)
}

func TestClientCallWith500Error(t *testing.T) {
	// act
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer server.Close()
	c := fogospt.NewClient(nil)
	c.BaseURL, _ = url.Parse(server.URL)
	// want
	dummyResponse := new(dummyStruct)
	_, err := c.NewRequest("/endpoint", dummyResponse)
	// assert
	assert.EqualError(t, err, "Server error")
}

func fixture(path string) string {
	b, err := os.ReadFile("../fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
