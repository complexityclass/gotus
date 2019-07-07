package task_scheduler

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"time"
)

func TestScheduleTasksPositive(t *testing.T) {
	//given
	var tasks []func() (string, error)
	tasks = append(tasks, func() (string, error) {
		time.Sleep(1 * time.Second)
		return "ok 0", nil
	})

	tasks = append(tasks, func() (string, error) {
		time.Sleep(2 * time.Second)
		return "ok 1", nil
	})

	tasks = append(tasks, func() (string, error) {
		time.Sleep(1 * time.Second)
		return "", errors.New("err")
	})

	tasks = append(tasks, func() (string, error) {
		time.Sleep(1 * time.Second)
		return "ok 3", nil
	})

	timeout := time.After(10 * time.Second)
	var expected = []string{"ok 0", "ok 1", "ok 3"}

	//when
	var resultsChan = ScheduleTasks2(tasks, 3, 1)
	var actual []string
OUTER:
	for {
		select {
		case result := <-resultsChan:
			actual = append(actual, result)
			if len(actual) == len(expected) {
				//then
				sort.Strings(actual)
				sort.Strings(expected)
				assert.Equal(t, expected, actual)
				break OUTER
			}
		case <-timeout:
			t.Fail()
			break OUTER
		}
	}
}

func TestScheduleTasksNegative(t *testing.T) {
	//given
	var tasks []func() (string, error)
	tasks = append(tasks, func() (string, error) {
		time.Sleep(1 * time.Second)
		return "", errors.New("err1")
	})

	tasks = append(tasks, func() (string, error) {
		time.Sleep(1 * time.Second)
		return "", errors.New("err2")
	})

	tasks = append(tasks, func() (string, error) {
		return "ok 1", nil
	})

	tasks = append(tasks, func() (string, error) {
		time.Sleep(1 * time.Second)
		return "", errors.New("err3")
	})

	tasks = append(tasks, func() (string, error) {
		time.Sleep(6 * time.Second)
		return "ok2", nil
	})

	timeout := time.After(10 * time.Second)
	var expected = []string{"ok 1"}

	//when
	var resultsChan = ScheduleTasks2(tasks, 3, 2)
	var actual []string
OUTER:
	for {
		select {
		case result, ok := <-resultsChan:
			if !ok {
				//channel closed
				assert.Equal(t, expected, actual)
				break OUTER
			} else {
				actual = append(actual, result)
			}
		case <-timeout:
			t.Fail()
			break OUTER
		}
	}
}
