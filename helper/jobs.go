package helper

import (
	"context"
	"os"
	"os/signal"
	"time"
)

type JobTime struct {
	H int
	M int
	S int
}

func (t JobTime) After(t2 time.Time) bool {
	if t.H > t2.Hour() {
		return true
	} else if t.H == t2.Hour() && t.M > t2.Minute() {
		return true
	} else if t.M == t2.Minute() && t.S > t2.Second() {
		return true
	}
	return false
}

func Job(start JobTime, end JobTime, duration time.Duration, callable func(...interface{}), values ...interface{}) {
	ticker := time.NewTicker(duration)
	done := make(chan bool)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				now := time.Now().In(GetIranTimeZone())
				if !start.After(now) && end.After(now) {
					callable(values...)
				}
			}
		}
	}()

	<-ctx.Done()
	stop()
	done <- true
}
