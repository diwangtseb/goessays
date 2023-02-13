package main

func main() {
	maxHeap := []int{}
	heap := NewHeap[int](func(a, b int) bool {
		return a > b
	})
	heap.pop()
}

type heap[T any] struct {
	elms []T
	comp func(a, b T) bool
}

func NewHeap[T any](comp func(a, b T) bool) *heap[T] {
	return &heap[T]{
		elms: make([]T, 1),
		comp: comp,
	}
}

func (h *heap[T]) len() int {
	return len(h.elms)
}

func (h *heap[T]) swap(i int, j int) {
	h.elms[i], h.elms[j] = h.elms[j], h.elms[i]
}

func (h *heap[T]) push(t T) {
	h.elms = append(h.elms, t)
	h.heapifyUp(len(h.elms) - 1)
}

func (h *heap[T]) pop() {
}

func (h *heap[T]) heapifyUp(i int) {

}

func parentIndex(n int) int {
	return (n - 1) / 2
}

func leftChildenIndex(n int) int {
	return 2*n + 1
}

func rightChildenIndex(n int) int {
	return 2*n + 2
}
