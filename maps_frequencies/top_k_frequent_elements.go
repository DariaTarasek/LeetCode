package maps_frequencies

import (
	"container/heap"
)

type NumWithFreq struct {
	Num  int
	Freq int
}

type PriorityQueue []NumWithFreq

func TopKFrequent(nums []int, k int) []int {
	numsFreq := make(map[int]int)

	for _, v := range nums {
		numsFreq[v] += 1
	}
	pq := &PriorityQueue{}
	heap.Init(pq)
	for num, freq := range numsFreq {
		heap.Push(pq, NumWithFreq{num, freq})
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}
	res := make([]int, len(*pq))
	for i := 0; i < len(*pq); i++ {
		res[i] = pq.get(i).Num
	}
	return res
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Freq < pq[j].Freq
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(NumWithFreq))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq PriorityQueue) get(i int) NumWithFreq {
	return pq[i]
}
