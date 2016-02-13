package godashery

import (
	"time"
)

var checkInterval = 10 * time.Millisecond;

type callable func()

func RateLimit(interval time.Duration, callback callable) {
	start := time.Now()

	for {
		elapsed := time.Since(start)

		if elapsed > interval {
			start = time.Now()
			callback()
		} else {
			time.Sleep(checkInterval)
		}
	}
}
