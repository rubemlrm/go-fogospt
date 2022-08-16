package fogospt_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Rubemlrm/fogospt-golang-client/fogospt"
	"github.com/stretchr/testify/assert"
)

func TestIncidentService_GetActiveFires(t *testing.T) {
	// act
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, fixture("active_fires.json"))
	}))
	defer server.Close()
	c := fogospt.NewClient(nil)
	c.BaseURL, _ = url.Parse(server.URL)
	// want
	response, err := c.Incident.GetActiveFires()
	var data []fogospt.Incident
	data = append(data, fogospt.Incident{
		Coords:         false,
		OccurrenceCode: 3101,
		Important:      false,
		Active:         true,
	})
	expected := &fogospt.IncidentResponse{Success: true, Data: data}
	// assert
	assert.NoError(t, err)
	assert.Equal(t, response.Success, expected.Success)
	assert.Equal(t, response.Data[0].Active, expected.Data[0].Active)
}

func TestIncidentService_GetGeoJsonActiveFires(t *testing.T) {
	// act
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, fixture("active_fires_geojson.json"))
	}))
	defer server.Close()
	c := fogospt.NewClient(nil)
	c.BaseURL, _ = url.Parse(server.URL)
	// want
	response, err := c.Incident.GetGeoJSONActiveFires()
	var data []fogospt.Feature
	data = append(data, fogospt.Feature{
		Type: "Feature",
		Geometry: fogospt.GeometryGeoJSON{
			Type: "Point",
		},
	})
	expected := &fogospt.IncidentGeoJSONResponse{Type: "FeatureCollection", Features: data}
	// assert
	assert.NoError(t, err)
	assert.Equal(t, response.Type, expected.Type)
	assert.Equal(t, response.Features[0].Geometry.Type, expected.Features[0].Geometry.Type)
}

func TestIncidentService_GetActiveFiresError(t *testing.T) {
	// act
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer server.Close()
	c := fogospt.NewClient(nil)
	c.BaseURL, _ = url.Parse(server.URL)
	// want
	_, err := c.Incident.GetActiveFires()
	// assert
	assert.EqualError(t, err, "Server error")
}

func TestIncidentService_TestIncidentService_GetGeoJsonActiveFiresError(t *testing.T) {
	// act
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer server.Close()
	c := fogospt.NewClient(nil)
	c.BaseURL, _ = url.Parse(server.URL)
	// want
	_, err := c.Incident.GetGeoJSONActiveFires()
	// assert
	assert.EqualError(t, err, "Server error")
}
