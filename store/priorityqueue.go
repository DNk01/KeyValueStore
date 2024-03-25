package store

// PriorityQueue реализует очередь с приоритетами на основе кучи.
type PriorityQueue []*Item

func (pq *PriorityQueue) Len() int { return len((*pq)) }

func (pq *PriorityQueue) Less(i, j int) bool {
	// Важно! Мы хотим, чтобы Pop() давал нам наименьший, а не наибольший элемент,
	// поэтому используем меньше здесь.
	return (*pq)[i].ExpiresAt.Before((*pq)[j].ExpiresAt)
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
