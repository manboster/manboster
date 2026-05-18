package llm

import "math"

// CalculateCompactTokens returns when the tokens above which, it should be compacted and open a new conversation
func CalculateCompactTokens(m Model) uint64 {
	if m.Context == 0 {
		return 0
	}
	if m.MaxOutputTokens == 0 || (float64(m.Context)-float64(m.MaxOutputTokens)) < 0 {
		return uint64(math.Floor(float64(m.Context) * 0.6))
	}
	return uint64(math.Floor((float64(m.Context) - float64(m.MaxOutputTokens)) * 0.8))
}

// MergeCapabilityFields merges fields into one single field
func MergeCapabilityFields(fields []CapabilityType) CapabilityType {
	var ct CapabilityType
	for _, f := range fields {
		ct |= f
	}
	return ct
}
