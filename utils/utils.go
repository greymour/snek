package utils

type FillUnit interface {
	string
	rune
	float32
	float64
	int
	int8
	int16
	int32
	int64
	uint
	uint8
	uint16
	uint32
	uint64
}

// takes in an empty slice of type T and fills it with the provided fillUnit
func FillSlice[T any](slice []T, fillUnit T) {
	for i := range slice {
		// slice = append(slice, fillUnit)
		slice[i] = fillUnit
	}
}
