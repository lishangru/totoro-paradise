package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/totoro-paradise/goapp/internal/types"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const deviationStd = 1.0 / 50000.0

// GeneratedRoute represents the result of the GenerateRoute function.
type GeneratedRoute struct {
	MockRoute []types.Point `json:"mockRoute"`
	Distance  float64       `json:"distance"`
}

// GenerateRoute produces a mock running route with a similar behaviour to the
// original Node implementation. It builds a dense route by interpolating points
// from a template route and keeps sampling random points until the desired
// distance in kilometres is reached.
func GenerateRoute(distanceKM float64, task types.RunPoint) (GeneratedRoute, error) {
	if len(task.PointList) == 0 {
		return GeneratedRoute{}, errors.New("task is empty")
	}

	addDeviation := func(point [2]float64) [2]float64 {
		return [2]float64{
			NormalRandom(point[0], deviationStd),
			NormalRandom(point[1], deviationStd),
		}
	}

	// addPoints returns pointA followed by additional points between A and B.
	// The number of interpolated points depends on the distance between the two
	// coordinates and the configured step size. The function deliberately omits
	// pointB so that callers can decide when to append the final endpoint.
	addPoints := func(pointA, pointB [2]float64) [][2]float64 {
		const stepLength = 0.0001
		pointVector := NewVector([]float64{pointB[0] - pointA[0], pointB[1] - pointA[1]})
		numberOfPoints := int(pointVector.Norm / stepLength)
		points := make([][2]float64, 0, numberOfPoints)
		points = append(points, pointA)
		for i := 1; i < numberOfPoints; i++ {
			pointX := pointA[0] + float64(i)*stepLength*pointVector.UnitVector[0]
			pointY := pointA[1] + float64(i)*stepLength*pointVector.UnitVector[1]
			points = append(points, [2]float64{pointX, pointY})
		}
		return points
	}

	combinePoints := func() [][2]float64 {
		route := FormatRouteToAMap(task.PointList)
		combined := make([][2]float64, 0, len(route)*10)
		for i := 0; i < len(route); i++ {
			if i == len(route)-1 {
				combined = append(combined, route[i])
				break
			}
			for _, p := range addPoints(route[i], route[i+1]) {
				combined = append(combined, p)
			}
		}
		return combined
	}

	trimRoute := func(route [][2]float64) (points [][2]float64, distance float64) {
		distanceMeters := distanceKM * 1000
		start := rand.Intn(len(route))
		i := start
		points = append(points, addDeviation(route[start]))
		for distance < distanceMeters {
			point := addDeviation(route[i])
			points = append(points, point)
			distance = DistanceOfLine(points)
			i++
			if i >= len(route)-2 {
				i = 0
			}
		}
		return points, distance
	}

	route := combinePoints()
	points, dist := trimRoute(route)

	mockRoute := make([]types.Point, len(points))
	for i, point := range points {
		mockRoute[i] = types.Point{
			Longitude: fmt.Sprintf("%.6f", point[0]),
			Latitude:  fmt.Sprintf("%.6f", point[1]),
		}
	}

	return GeneratedRoute{
		MockRoute: mockRoute,
		Distance:  dist / 1000,
	}, nil
}
