package task_scheduler

import (
	"sync"
)

type TaskResult = func(processed []string)
type Task = func() (string, error)

func ScheduleTasks(tasks []Task, maxParallelCount int, errorLimit int, success TaskResult, failure TaskResult) {
	var run = func(task Task, group *sync.WaitGroup, results chan<- string, errors chan<- error) {
		var res, err = task()
		if err != nil {
			errors <- err
		} else {
			results <- res
		}
		group.Done()
	}

	var group sync.WaitGroup
	var nextIdx = make(chan int)
	var successes = make(chan string)
	var failures = make(chan error)
	var resultStore ResultStore
	var errorCounter ErrorCounter
	group.Add(len(tasks))

	for i := 0; i < maxParallelCount; i++ {
		go func() {
		OUTER:
			for {
				select {
				case taskIdx := <-nextIdx:
					go run(tasks[taskIdx], &group, successes, failures)
				case result := <-successes:
					resultStore.Add(result)
					if len(resultStore.Values()) == len(tasks) {
						break OUTER
					}
				case <-failures:
					errorCounter.Inc()
					if errorCounter.Value() > errorLimit {
						var done = resultStore.Values()
						var left = len(tasks) - len(done)
						for i := 0; i < left; i++ {
							group.Done()
						}
						failure(resultStore.Values())
						break OUTER
					}
				}
			}
		}()
	}

	go func() {
		for idx := range tasks {
			nextIdx <- idx
		}
	}()

	group.Wait()
	//close(successes)
	//close(failures)
	//close(nextIdx)
	success(resultStore.Values())
}

type ResultStore struct {
	mtx     sync.Mutex
	results []string
}

func (store *ResultStore) Add(result string) {
	store.mtx.Lock()
	store.results = append(store.results, result)
	store.mtx.Unlock()
}

func (store ResultStore) Values() []string {
	store.mtx.Lock()
	defer store.mtx.Unlock()
	return store.results
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

func ScheduleTasks2(tasks []Task, maxParallelCount int, errorLimit int) <-chan string {
	limiter := make(chan struct{}, maxParallelCount)
	var result = make(chan string)
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
				}
			}(task)
		}
	}()

	return result
}
