package ds

import "fmt"

func main() {
	heap := Heap[int]{}
	heap.mininsert(1)
	heap.mininsert(3)
	heap.mininsert(4)
	heap.mininsert(2)
	heap.mininsert(5)
	fmt.Printf("%v", heap)
}

type Heap[T float64 | int] struct {
	data []T
}

func (heap *Heap[T]) maxinsert(t T) {
	if len(heap.data) == 0 {
		heap.data = append(heap.data, t)
		return
	}
	heap.data = append(heap.data, t)
	for i := len(heap.data)/2 - 1; i >= 0; i-- {
		heap.maxheapify(i)
	}
}

func (heap *Heap[T]) maxheapify(index int) {
	size := len(heap.data)
	if size <= 1 {
		return
	}
	largest := index
	left := 2*index + 1
	right := 2*index + 2
	if left < size && heap.data[left] > heap.data[largest] {
		largest = left
	}
	if right < size && heap.data[right] > heap.data[largest] {
		largest = right
	}
	if largest != index {
		heap.data[index], heap.data[largest] = heap.data[largest], heap.data[index]
		heap.maxheapify(largest)
	}
}

func (heap *Heap[T]) mininsert(t T) {
	if len(heap.data) == 0 {
		heap.data = append(heap.data, t)
		return
	}
	heap.data = append(heap.data, t)
	for i := len(heap.data)/2 - 1; i >= 0; i-- {
		heap.minheapify(i)
	}
}

func (heap *Heap[T]) minheapify(index int) {
	size := len(heap.data)
	if size <= 1 {
		return
	}
	largest := index
	left := 2*index + 1
	right := 2*index + 2
	if left < size && heap.data[left] < heap.data[largest] {
		largest = left
	}
	if right < size && heap.data[right] < heap.data[largest] {
		largest = right
	}
	if largest != index {
		heap.data[index], heap.data[largest] = heap.data[largest], heap.data[index]
		heap.minheapify(largest)
	}
}
