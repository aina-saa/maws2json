package data

import "time"

// TA60sAvg,DP60sAvg,RH60sAvg,PA60sAvg,QFF60sAvg,SR60sSum,PR60sSum,WD2minAvg,WS2minAvg,
// °C,°C,%,hPa,hPa,W/m2,mm,°,m/s,
type Log struct {
	Site      *string    `json:"site_id,omitempty"`
	Message   string     `json:"message"`
	Timestamp *time.Time `json:"datetime,omitempty"`
	Data1     DataPoint  `json:"TA60sAvg"`
	Data2     DataPoint  `json:"DP60sAvg"`
	Data3     DataPoint  `json:"RH60sAvg"`
	Data4     DataPoint  `json:"PA60sAvg"`
	Data5     DataPoint  `json:"QFF60sAvg"`
	Data6     DataPoint  `json:"SR60sSum"`
	Data7     DataPoint  `json:"PR60sSum"`
	Data8     DataPoint  `json:"WD2minAvg"`
	Data9     DataPoint  `json:"WS2minAvg"`
}

var LogLength = 10 // ends with empty tab

var LogMeasurementUnits = map[int]string{
	1: "°C",
	2: "°C",
	3: "%",
	4: "hPa",
	5: "hPa",
	6: "W/m2",
	7: "mm",
	8: "°",
	9: "m/s",
}

var LogDescription = map[int]string{
	1: "Average air temperature over 60 seconds.",
	2: "Average dew point over 60 seconds.",
	3: "Average relative humidity over 60 seconds.",
	4: "Average atmospheric pressure over 60 seconds.",
	5: "Average atmospheric pressure at sea level over 60 seconds.",
	6: "Average sun radiation over 60 seconds.",
	7: "Average precipation over 60 seconds.",
	8: "Average wind direction over two (2) minutes.",
	9: "Average wind speed over two (2) minutes.",
}

// eof
