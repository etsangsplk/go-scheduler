package scheduler

import (
	"time"
)

// RunTaskAtInterval - run funcion every specified time interval, optionaly you can add delay at start
func RunTaskAtInterval(funcToRun func(), interval time.Duration, startDelay time.Duration) {
	go func(funcToRun func(), interval time.Duration, startDelay time.Duration) {
		time.Sleep(startDelay)
		ticker := time.NewTicker(interval)
		for range ticker.C {
			funcToRun()
		}
	}(funcToRun, interval, startDelay)
}
