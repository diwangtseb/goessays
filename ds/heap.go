package ds

type Heap struct {
	Data []int
}

func (h *Heap) InsertMax(data int) {
	if len(h.Data) == 0 {
		h.Data = append(h.Data, data)
		return
	}
	h.Data = append(h.Data, data)
	for mid := len(h.Data)/2 - 1; mid >= 0; mid-- {
		h.MaxHeapify(mid)
	}
}

func (h *Heap) MaxHeapify(index int) {
	size := len(h.Data)
	if size <= 1 {
		return
	}
	largest := index
	left := 2*index + 1
	right := 2*index + 2
	if left < size && h.Data[left] > h.Data[largest] {
		largest = left
	}
	if right < size && h.Data[right] > h.Data[largest] {
		largest = right
	}
	if largest != index {
		h.Data[index], h.Data[largest] = h.Data[largest], h.Data[index]
		h.MaxHeapify(largest)
	}
}
