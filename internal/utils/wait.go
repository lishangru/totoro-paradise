package utils

import "time"

// Wait sleeps for the provided duration. The function is extracted so it can be
// stubbed in tests when we want deterministic behaviour.
func Wait(duration time.Duration) {
	time.Sleep(duration)
}
