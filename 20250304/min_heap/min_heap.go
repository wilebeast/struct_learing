package min_heap

type MinHeap struct {
	heap []int
}

func (h *MinHeap) Insert(value int) {
	h.heap = append(h.heap, value)
	h.popUp(len(h.heap) - 1)
}

func (h *MinHeap) popUp(index int) {
	for {
		parent := (index - 1) / 2
		if parent < 0 || h.heap[parent] < h.heap[index] {
			return
		}
		h.heap[parent], h.heap[index] = h.heap[index], h.heap[parent]
		index = parent
	}
}

func (h *MinHeap) Remove() (int, bool) {
	if len(h.heap) == 0 {
		return 0, false
	}
	value := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.popDown(0)
	return value, true
}

func (h *MinHeap) popDown(index int) {
	for {
		left := 2*index + 1
		right := 2*index + 2
		minIndex := index
		if left < len(h.heap)-1 && h.heap[left] < h.heap[minIndex] {
			minIndex = left
		}
		if right < len(h.heap)-1 && h.heap[right] < h.heap[minIndex] {
			minIndex = right
		}
		if index == minIndex {
			return
		}
		h.heap[index], h.heap[minIndex] = h.heap[minIndex], h.heap[index]
		index = minIndex
	}
}
