package task

import (
	"log"
	"time"
)

var Tasks = make(map[int]*Task)

var NextID = 1

func CalculateNextExecutionTime(now time.Time, interval int) time.Time {
	minute := now.Minute()
	hour := now.Hour()
	targetMinute := ((minute / interval) + 1*interval)

	if targetMinute >= 60 {
		targetMinute -= 60
		hour++
		if hour == 24 {
			hour = 0
		}
	}

	nextTime := time.Date(now.Year(), now.Month(), now.Day(), hour, targetMinute, 0, 0, now.Location())
	if nextTime.Before(now) || nextTime.Equal(now) {
		nextTime = nextTime.Add(time.Minute * time.Duration(interval))
	}

	return nextTime

}

func StartTaskGoroutine(task *Task) {
	go func(t *Task) {
		defer close(t.DoneChan)
		for {
			select {
			case <-t.DoneChan:
				return
			default:
				nextTime := CalculateNextExecutionTime(time.Now(), t.Schedule)
				duration := nextTime.Sub(time.Now())
				timer := time.After(duration)
				select {
				case <-timer:
					t.Execute()
					log.Printf("Task %d (%s) executed at %s", t.ID, t.Name, time.Now().Format("2006-01-02 15:04:05"))
				}
			}
		}
	}(task)
}
