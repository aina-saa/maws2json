package data

import "time"

// WS,WD
// m/s, °
type Wind struct {
	Site      *string    `json:"site_id,omitempty"`
	Message   string     `json:"message"`
	Timestamp *time.Time `json:"datetime,omitempty"`
	Data1     DataPoint  `json:"WSCur"`
	Data2     DataPoint  `json:"WDCur"`
}

var WindLength = 3 // ends with empty tab

var WindMeasurementUnits = map[int]string{
	1: "m/s",
	2: "°",
}

var WindDescription = map[int]string{
	1: "Current wind speed",
	2: "Current wind direction",
}

// eof
