package repository

import (
	"arvore/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ArcGISResponse struct {
	Features []struct {
		Attributes map[string]interface{} `json:"attributes"`
		Geometry   struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"geometry"`
	} `json:"features"`
}

type TreeExternalRepository struct{}

func NewTreeExternalRepository() *TreeExternalRepository {
	return &TreeExternalRepository{}
}

func (r *TreeExternalRepository) GetTrees() ([]map[string]interface{}, error) {
	url := "https://services.arcgis.com/.../FeatureServer/0/query?where=1=1&outFields=*&f=json&resultRecordCount=50"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data ArcGISResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	var trees []map[string]interface{}

	for _, f := range data.Features {
		tree := map[string]interface{}{
			"lat":     f.Geometry.Y,
			"lng":     f.Geometry.X,
			"species": f.Attributes["species"],
			"health":  f.Attributes["condition"],
		}
		trees = append(trees, tree)
	}

	return trees, nil
}
func getString(m map[string]interface{}, key string) string {
	if val, ok := m[key]; ok && val != nil {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}
func (r *TreeExternalRepository) GetNearbyTrees(lat, lng, distance string) ([]model.TreedData, error) {

	url := "https://services.arcgis.com/ZOyb2t4B0UYuYNYH/arcgis/rest/services/SDOT_Trees_(Active)/FeatureServer/0/query"

	latF, _ := strconv.ParseFloat(lat, 64)
	lngF, _ := strconv.ParseFloat(lng, 64)

	delta := 0.01 // ajuste se quiser mais área

	minLat := latF - delta
	maxLat := latF + delta
	minLng := lngF - delta
	maxLng := lngF + delta

	params := fmt.Sprintf(
		"?where=1=1&outFields=*"+
			"&geometry=%f,%f,%f,%f"+
			"&geometryType=esriGeometryEnvelope"+
			"&inSR=4326"+
			"&spatialRel=esriSpatialRelIntersects"+
			"&f=json",
		minLng, minLat, maxLng, maxLat,
	)

	resp, err := http.Get(url + params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data ArcGISResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	var trees []model.TreedData

	for _, f := range data.Features {

		tree := model.TreedData{
			Latitude:  f.Geometry.Y,
			Longitude: f.Geometry.X,
			Species:   getString(f.Attributes, "COMMON_NAME"),
			Health:    getString(f.Attributes, "CONDITION"),
		}

		trees = append(trees, tree)
	}

	return trees, nil
}
