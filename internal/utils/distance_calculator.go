package utils

import "math"

// DistanceBetweenPoints returns the approximate distance in metres between two
// longitude/latitude points. The implementation mirrors the JavaScript
// counterpart by using a simplified spherical law of cosines formula.
func DistanceBetweenPoints(pointA, pointB [2]float64) float64 {
	const rad = 0.0174532925194329
	d2, d3 := pointA[0], pointA[1]
	d4, d5 := pointB[0], pointB[1]
	d2 *= rad
	d3 *= rad
	d4 *= rad
	d5 *= rad
	d6 := math.Sin(d2)
	d7 := math.Sin(d3)
	d8 := math.Cos(d2)
	d9 := math.Cos(d3)
	d10 := math.Sin(d4)
	d11 := math.Sin(d5)
	d12 := math.Cos(d4)
	d13 := math.Cos(d5)
	s11 := d9 * d8
	s12 := d9 * d6
	s13 := d7
	s21 := d13 * d12
	s22 := d13 * d10
	s23 := d11

	d14 := math.Sqrt(
		math.Pow(s11-s21, 2) +
			math.Pow(s12-s22, 2) +
			math.Pow(s13-s23, 2),
	)

	return math.Asin(d14/2.0) * 1.2740015798544e7
}

// DistanceOfLine accumulates the distance of a polyline of geo points.
func DistanceOfLine(line [][2]float64) float64 {
	distance := 0.0
	for i := 0; i < len(line)-1; i++ {
		distance += DistanceBetweenPoints(line[i], line[i+1])
	}
	return distance
}
