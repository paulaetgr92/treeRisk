package model

type TreedData struct {
	ID        int64   `json:"id,omitempty"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Species   string  `json:"species"`
	Health    string  `json:"health"`
}
