package utils

import (
	"math"
	"testing"
)

func TestDistanceBetweenPoints(t *testing.T) {
	distance := DistanceBetweenPoints([2]float64{0, 0}, [2]float64{0, 1})
	expected := 111195.0
	if math.Abs(distance-expected) > 100 {
		t.Fatalf("expected roughly %f meters, got %f", expected, distance)
	}
}

func TestDistanceOfLine(t *testing.T) {
	line := [][2]float64{{0, 0}, {0, 0.5}, {0, 1}}
	distance := DistanceOfLine(line)
	expected := DistanceBetweenPoints([2]float64{0, 0}, [2]float64{0, 0.5}) +
		DistanceBetweenPoints([2]float64{0, 0.5}, [2]float64{0, 1})

	if math.Abs(distance-expected) > 0.001 {
		t.Fatalf("expected cumulative distance %f, got %f", expected, distance)
	}
}
