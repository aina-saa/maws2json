package data

type DataPoint struct {
	Description string  `json:"description,omitempty"`
	Value       float64 `json:"value"`
	Unit        string  `json:"unit_of_measurement,omitempty"`
}

// eof
