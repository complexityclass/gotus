package frequency_counter

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordHeap(t *testing.T) {
	//given
	words := map[string]int{
		"word0": 4,
		"word1": 5,
		"word2": 3,
		"word3": 99,
	}
	expected := []string{
		"word2", "word0", "word1",
	}

	//when
	var pq = make(PriorityQueue, len(words))
	var i = 0
	for value, priority := range words {
		pq[i] = &WordInfo{
			value:    value,
			priority: priority,
			index:    i,
		}
		i += 1
	}
	heap.Init(&pq)
	var result []string

	for pq.Len() > 0 {
		var item = heap.Pop(&pq).(*WordInfo)
		result = append(result, item.value)
	}
	result = result[:3]

	//then
	assert.Equal(t, expected, result)
}

func TestWordFrequency(t *testing.T) {
	//given
	var filePath = "sample.txt"
	expected := [10]string{
		"the",
		"of",
		"Ipsum",
		"Lorem",
		"and",
		"text",
		"a",
		"has",
		"type",
		"dummy",
	}

	//when
	var topTen = TopTen(filePath)

	//then
	assert.Equal(t, expected, topTen)
}
