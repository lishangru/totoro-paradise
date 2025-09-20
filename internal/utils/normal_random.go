package utils

import (
	"math"
	"math/rand"
)

// NormalRandom returns a value sampled from a normal distribution with the
// provided mean and standard deviation. Values beyond three standard deviations
// are rejected to keep the output stable just like the original TypeScript
// implementation.
func NormalRandom(mean, std float64) float64 {
	for {
		u := rand.Float64()*2 - 1.0
		v := rand.Float64()*2 - 1.0
		w := u*u + v*v
		if w == 0 || w >= 1 {
			continue
		}
		c := math.Sqrt((-2 * math.Log(w)) / w)
		result := mean + u*c*std
		if result < mean-3*std || result > mean+3*std {
			continue
		}
		return result
	}
}
