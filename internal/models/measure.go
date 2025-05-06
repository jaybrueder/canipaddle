package models

type Measure struct {
	StationName string  `json:"station_name"`
	StationID   int     `json:"station_id"`
	SensorID    int     `json:"sensor_id"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Dt          string  `json:"dt"` // Keeping as string for now, can be parsed to time.Time if needed
	Value       float64 `json:"value"`
}
