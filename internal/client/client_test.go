package client_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Rubemlrm/fogospt-golang-client/internal/client"
	"github.com/stretchr/testify/assert"
)

func TestClientCallWithSuccess(t *testing.T) {
	expected := "dummy data"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer server.Close()

	req, err := client.Do(server.URL)
	assert.NoError(t, err)
	assert.Equal(t, 200, req.StatusCode)
}

func TestClientCallWith500Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer server.Close()
	_, err := client.Do(server.URL)
	assert.EqualError(t, err, "Server error")
}
