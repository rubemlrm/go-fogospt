package fogospt

type created struct {
	Sec int `json:"sec,omitempty"`
}

type updated struct {
	Sec int `json:"sec,omitempty"`
}

type datetime struct {
	Sec int `json:"sec,omitempty"`
}

// IncidentService handles all communication with incident endpoints
type IncidentService service

// IncidentResponse represents the endpoint response
type IncidentResponse struct {
	Success bool       `json:"success,omitempty"`
	Data    []Incident `json:"data,omitempty"`
}

// IncidentGeoJSONResponse represents endpoint on geojson format
type IncidentGeoJSONResponse struct {
	Type     string    `json:"type,omitempty"`
	Features []Feature `json:"features,omitempty"`
}

// Incident Represents incident details
type Incident struct {
	ID             string   `json:"id,omitempty"`
	Coords         bool     `json:"coords,omitempty"`
	Date           string   `json:"date,omitempty"`
	Hour           string   `json:"hour,omitempty"`
	Location       string   `json:"location,omitempty"`
	Aerial         int      `json:"aerial,omitempty"`
	Terrain        int      `json:"terrain,omitempty"`
	Man            int      `json:"man,omitempty"`
	District       string   `json:"district,omitempty"`
	County         string   `json:"concelho,omitempty"`
	Dico           string   `json:"dico,omitempty"`
	Parish         string   `json:"freguesia,omitempty"`
	Lat            float64  `json:"lat,omitempty"`
	Lng            float64  `json:"lng,omitempty"`
	OccurrenceCode int      `json:"naturezaCode,omitempty"`
	Nature         string   `json:"natureza,omitempty"`
	StatusCode     int      `json:"statusCode,omitempty"`
	StatusColor    string   `json:"statusColor,omitempty"`
	Status         string   `json:"status,omitempty"`
	Important      bool     `json:"important,omitempty"`
	Active         bool     `json:"active,omitempty"`
	SadoID         string   `json:"sadoId,omitempty"`
	SharepointID   int      `json:"sharepointId,omitempty"`
	Extra          string   `json:"extra,omitempty"`
	Disappear      bool     `json:"disappear,omitempty"`
	DateTime       datetime `json:"datetime,omitempty"`
	Created        created  `json:"created,omitempty"`
	Updated        updated  `json:"updated,omitempty"`
}

// Feature represents the geojson fire details
type Feature struct {
	Type       string            `json:"type,omitempty"`
	Geometry   GeometryGeoJSON   `json:"geometry,omitempty"`
	Properties FeatureProperties `json:"properties,omitempty"`
}

// FeatureProperties represents the geojson fire details properties
type FeatureProperties struct {
	ID                string   `json:"id,omitempty"`
	Coords            bool     `json:"coords,omitempty"`
	Date              string   `json:"date,omitempty"`
	Hour              string   `json:"hour,omitempty"`
	Location          string   `json:"location,omitempty"`
	Aerial            int      `json:"aerial,omitempty"`
	Terrain           int      `json:"terrain,omitempty"`
	Man               int      `json:"man,omitempty"`
	District          string   `json:"district,omitempty"`
	County            string   `json:"concelho,omitempty"`
	Dico              string   `json:"dico,omitempty"`
	Parish            string   `json:"freguesia,omitempty"`
	Lat               float64  `json:"lat,omitempty"`
	Lng               float64  `json:"lng,omitempty"`
	OccurrenceCode    string   `json:"naturezaCode,omitempty"`
	Nature            string   `json:"natureza,omitempty"`
	StatusCode        int      `json:"statusCode,omitempty"`
	EspecieName       string   `json:"especieName,omitempty"`
	FamilyName        string   `json:"familyName,omitempty"`
	StatusColor       string   `json:"statusColor,omitempty"`
	Status            string   `json:"status,omitempty"`
	Important         bool     `json:"important,omitempty"`
	Active            bool     `json:"active,omitempty"`
	SadoID            string   `json:"sadoId,omitempty"`
	SharepointID      int      `json:"sharepointId,omitempty"`
	Extra             string   `json:"extra,omitempty"`
	Disappear         bool     `json:"disappear,omitempty"`
	DetailLocation    string   `json:"detailLocation,omitempty"`
	KML               string   `json:"kml,omitempty"`
	PCO               string   `json:"pco,omitempty"`
	COS               string   `json:"cos,omitempty"`
	HeliFight         int      `json:"heliFight,omitempty"`
	HeliCoord         int      `json:"heliCoord,omitempty"`
	PlaneFight        int      `json:"planeFight,omitempty"`
	AnepcDirectUpdate bool     `json:"anepcDirectUpdate,omitempty"`
	ICNF              ICNF     `json:"icnf,omitempty"`
	DateTime          datetime `json:"datetime,omitempty"`
	Created           created  `json:"created,omitempty"`
	Updated           updated  `json:"updated,omitempty"`
}

// IDGeoJSON represents database id
type IDGeoJSON struct {
	ID IDInnerGeoJSON `json:"_id,omitempty"`
}

// IDInnerGeoJSON id property for IDGeoJSON
type IDInnerGeoJSON struct {
	ID string `json:"$id,omitempty"`
}

// GeometryGeoJSON represents the type of coordinates and location
type GeometryGeoJSON struct {
	Type        string    `json:"type,omitempty"`
	Coordinates []float32 `json:"coordinates,omitempty"`
}

// BurnArea represents the stats of burned area by type
type BurnArea struct {
	Povoamento float32 `json:"povoamento,omitempty"`
	Rural      float32 `json:"agricola,omitempty"`
	Bush       float32 `json:"mato,omitempty"`
	Total      float32 `json:"total,omitempty"`
}

// ICNF represents the stats of burned area and alert source
type ICNF struct {
	BurnArea    BurnArea `json:"burnArea,omitempty"`
	Altitude    float32  `json:"altitude,omitempty"`
	Fire        bool     `json:"incendio,omitempty"`
	AlertSource string   `json:"fontealerta,omitempty"`
}

// GetActiveFires get all of active fires on real time
func (s *IncidentService) GetActiveFires() (*IncidentResponse, error) {
	response := &IncidentResponse{}
	_, err := s.client.NewRequest("/v2/incidents/active", response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// GetGeoJSONActiveFires get all of active fires on real time using a geojson format
func (s *IncidentService) GetGeoJSONActiveFires() (*IncidentGeoJSONResponse, error) {
	response := &IncidentGeoJSONResponse{}
	_, err := s.client.NewRequest("/v2/incidents/active?geojson=false", response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
