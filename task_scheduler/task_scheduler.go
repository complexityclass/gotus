package task_scheduler

import (
	"sync"
)

type Task = func() (string, error)

func ScheduleTasks2(tasks []Task, maxParallelCount int, errorLimit int) <-chan string {
	limiter := make(chan struct{}, maxParallelCount)
	var result = make(chan string)
	var errorCounter ErrorCounter
	go func() {
		for _, task := range tasks {
			limiter <- struct{}{}
			go func(callable Task) {
				defer func() {
					<-limiter
				}()
				value, err := callable()
				if err == nil {
					result <- value
				} else {
					errorCounter.Inc()
					if errorCounter.Value() > errorLimit {
						close(result)
					}
				}
			}(task)
		}
	}()

	return result
}

type ErrorCounter struct {
	mtx   sync.Mutex
	value int
}

func (counter *ErrorCounter) Inc() {
	counter.mtx.Lock()
	counter.value += 1
	counter.mtx.Unlock()
}

func (counter ErrorCounter) Value() int {
	counter.mtx.Lock()
	defer counter.mtx.Unlock()
	return counter.value
}
