package data

import "time"

// TA60sAvg,TA24hMin,TA24hMax,DP60sAvg,DP24hMin,DP24hMax,RH60sAvg,RH24hMin,RH24hMax,PA60sAvg,PA24hMin,PA24hMax,QFF60sAvg,QFF24hMin,QFF24hMax,SR60sAvg,SR24hMin,SR24hMax,PR1hSum,PR6hSum,PR24hSum,,,,,,
// °C,°C,°C,°C,°C,°C,%,%,%,hPa,hPa,hPa,hPa,hPa,hPa,W/m2,W/m2,W/m2,mm,mm,mm,,,,,,
type Ptu struct {
	Site      *string    `json:"site_id,omitempty"`
	Message   string     `json:"message"`
	Timestamp *time.Time `json:"datetime_utc,omitempty"`
	Data1     DataPoint  `json:"TA60sAvg"`
	Data2     DataPoint  `json:"TA24hMin"`
	Data3     DataPoint  `json:"TA24hMax"`
	Data4     DataPoint  `json:"DP60sAvg"`
	Data5     DataPoint  `json:"DP24hMin"`
	Data6     DataPoint  `json:"DP24hMax"`
	Data7     DataPoint  `json:"RH60sAvg"`
	Data8     DataPoint  `json:"RH24hMin"`
	Data9     DataPoint  `json:"RH24hMax"`
	Data10    DataPoint  `json:"PA60sAvg"`
	Data11    DataPoint  `json:"PA24hMin"`
	Data12    DataPoint  `json:"PA24hMax"`
	Data13    DataPoint  `json:"QFF60sAvg"`
	Data14    DataPoint  `json:"QFF24hMin"`
	Data15    DataPoint  `json:"QFF24hMax"`
	Data16    DataPoint  `json:"SR60sAvg"`
	Data17    DataPoint  `json:"SR24hMin"`
	Data18    DataPoint  `json:"SR24hMax"`
	Data19    DataPoint  `json:"PR1hSum"`
	Data20    DataPoint  `json:"PR6hSum"`
	Data21    DataPoint  `json:"PR24hSum"`
}

var PtuLength = 22 // ends with empty tab

var PtuMeasurementUnits = map[int]string{
	1:  "°C",
	2:  "°C",
	3:  "°C",
	4:  "°C",
	5:  "°C",
	6:  "°C",
	7:  "%",
	8:  "%",
	9:  "%",
	10: "hPa",
	11: "hPa",
	12: "hPa",
	13: "hPa",
	14: "hPa",
	15: "hPa",
	16: "W/m2",
	17: "W/m2",
	18: "W/m2",
	19: "mm",
	20: "mm",
	21: "mm",
}

var PtuDescription = map[int]string{
	1:  "Average air temperature over 60 seconds.",
	2:  "Minimum average air temperature over 60 seconds measured in the past 24 hours.",
	3:  "Maximum average air temperature over 60 seconds measured in the past 24 hours.",
	4:  "Average dew point over 60 seconds.",
	5:  "Minimum average dew point over 60 seconds measured in the past 24 hours.",
	6:  "Maximum average dew point over 60 seconds measured in the past 24 hours.",
	7:  "Average relative humidity over 60 seconds.",
	8:  "Minimum average relative humidity over 60 seconds measured in the past 24 hours.",
	9:  "Maximum average relative humidity over 60 seconds measured in the past 24 hours.",
	10: "Average atmospheric pressure over 60 seconds.",
	11: "Minimum average atmospheric pressure over 60 seconds measured in the past 24 hours.",
	12: "Maximum average atmospheric pressure over 60 seconds measured in the past 24 hours.",
	13: "Average atmospheric pressure at sea level over 60 seconds.",
	14: "Minimum average atmospheric pressure at sea level over 60 seconds measured in the past 24 hours.",
	15: "Maximum average atmospheric pressure at sea level over 60 seconds measured in the past 24 hours.",
	16: "Average sun radiation of one (1) minute aka 60 seconds.",
	17: "Minimum average sun radiation over 60 seconds measured in the past 24 hours.",
	18: "Maximum average sun radiation over 60 seconds measured in the past 24 hours.",
	19: "Sum of precipation over last one (1) hour.",
	20: "Sum of precipation over last six (6) hours.",
	21: "Sum of precipation over last 24 hours.",
}

// eof
