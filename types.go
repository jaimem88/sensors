package sensors

// SensorData describes each item in the input file per the spec
type SensorData struct {
	ID          string  `json:"id,omitempty"`
	Timestamp   int64   `json:"timestamp,omitempty"`
	Temperature float64 `json:"temperature,omitempty"` // this could be decimal.Decimal if we care about float representation and precision. For this excercise it should suffice to use float64
}

// InputData represents all the data loaded from a file
type InputData []*SensorData

// Result represents the processed value per sensor with the required information
type Result struct {
	ID      string    `json:"id,omitempty"`
	Average float64   `json:"average,omitempty"`
	Median  float64   `json:"median,omitempty"`
	Mode    []float64 `json:"mode,omitempty"`
}
