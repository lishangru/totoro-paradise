package utils

import "math"

// Vector represents a mathematical vector with its norm and unit vector.
type Vector struct {
	Values     []float64
	Norm       float64
	UnitVector []float64
}

func NewVector(values []float64) Vector {
	norm := 0.0
	for _, v := range values {
		norm = math.Hypot(norm, v)
	}
	unit := make([]float64, len(values))
	if norm > 0 {
		for i, v := range values {
			unit[i] = v / norm
		}
	}
	return Vector{
		Values:     values,
		Norm:       norm,
		UnitVector: unit,
	}
}
