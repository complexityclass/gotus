package frequency_counter

import (
	"bufio"
	"container/heap"
	"github.com/derekparker/trie"
	"log"
	"os"
	"strings"
)

func TopTen(filePath string) [10]string {
	var res [10]string
	var wp = frequencyCounter(filePath, 10)
	for i := len(wp) - 1; i >= 0; i-- {
		res[len(wp)-i-1] = wp[i].a.(string)
	}

	return res
}

func frequencyCounter(filePath string, count int) (top []Pair) {
	top = make([]Pair, 0)
	var fTrie = trie.New()
	var pq = make(PriorityQueue, 0)
	heap.Init(&pq)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var words = strings.Fields(filterPunctuation(scanner.Text()))
		for _, word := range words {
			node, ok := fTrie.Find(word)
			var wordCount int
			if ok {
				var prevCount = node.Meta().(int)
				wordCount = prevCount + 1
			} else {
				wordCount = 1
			}

			fTrie.Add(word, wordCount)

			var currInfo *WordInfo
			for _, wi := range pq {
				if wi.value == word {
					currInfo = wi
				}
			}

			if currInfo != nil {
				pq.update(currInfo, currInfo.value, currInfo.priority+1)
			} else {
				var wordInfo = &WordInfo{
					value:    word,
					priority: 1,
				}
				heap.Push(&pq, wordInfo)
				pq.update(wordInfo, wordInfo.value, wordCount)
			}

			for pq.Len() > count {
				_ = heap.Pop(&pq)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}

	for pq.Len() > 0 {
		var info = heap.Pop(&pq).(*WordInfo)
		var p = Pair{
			a: info.value,
			b: info.priority,
		}
		top = append(top, p)
	}

	return
}

func filterPunctuation(str string) string {
	var replacer = strings.NewReplacer(",", "", ".", "", ";", "")
	return replacer.Replace(str)
}
