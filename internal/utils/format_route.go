package utils

import "github.com/totoro-paradise/goapp/internal/types"

// FormatPointToAMap converts a point with string coordinates into numeric
// longitude/latitude pair.
func FormatPointToAMap(point types.Point) [2]float64 {
	return [2]float64{
		parseFloat(point.Longitude),
		parseFloat(point.Latitude),
	}
}

// FormatRouteToAMap maps a slice of points to their numeric representation.
func FormatRouteToAMap(route []types.Point) [][2]float64 {
	formatted := make([][2]float64, len(route))
	for i, point := range route {
		formatted[i] = FormatPointToAMap(point)
	}
	return formatted
}
