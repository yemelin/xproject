package summary

// Config represents the configuration to the Summary
type Config struct {
	// Header to group by
	GroupBy string

	// Header to sum the values
	// Values should be convertable to float64
	SumBy []string
}

// NewConfig returns a new Config
func NewConfig(groupBy string, sumBy []string) *Config {
	return &Config{
		GroupBy: groupBy,
		SumBy:   sumBy,
	}
}
