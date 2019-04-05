package risk

// Rate is the representation of the likelihood or impact level
type Rate struct {
	Value float64
}

const (
	low    = "LOW"
	medium = "MEDIUM"
	high   = "HIGH"
)

// Color returns the color to be used for each rate level
func (r *Rate) Color() string {
	if r.Value < 3.0 {
		return "lightgreen"
	} else if r.Value < 6.0 {
		return "yellow"
	}
	return "red"
}

// Level display the rate in human comprehensive language
func (r *Rate) Level() string {
	if r.Value < 3.0 {
		return low
	} else if r.Value < 6.0 {
		return medium
	}
	return high
}

// IsLow check if level is low
func (r *Rate) IsLow() bool {
	return low == r.Level()
}

// IsMedium check if level is medium
func (r *Rate) IsMedium() bool {
	return medium == r.Level()
}

// IsHigh check if level is high
func (r *Rate) IsHigh() bool {
	return high == r.Level()
}
