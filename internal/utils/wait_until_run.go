package utils

import (
	"fmt"
	"time"
)

// TimeNow is a helper so the current time can be mocked in tests.
var TimeNow = time.Now

// Waiter abstracts the waiting behaviour. It allows unit tests to stub the
// delay without sleeping for real.
type Waiter func(time.Duration)

// WaitUntilRun prints a friendly message and blocks the execution until the
// provided target time is reached. When the target is in the past the function
// returns immediately.
func WaitUntilRun(target time.Time, wait Waiter) {
	if wait == nil {
		wait = Wait
	}

	fmt.Printf("请等到%02d:%02d:%02d\n", target.Hour(), target.Minute(), target.Second())
	now := TimeNow()
	if target.Before(now) {
		return
	}

	wait(target.Sub(now))
}
