package first_version

type MinHeap struct {
	arr []int
}

func (h *MinHeap) Insert(value int) {
	h.arr = append(h.arr, value)
	h.popUp(len(h.arr) - 1)
}

func (h *MinHeap) popUp(index int) {
	for {
		parent := (index - 1) / 2
		if parent < 0 || h.arr[parent] < h.arr[index] {
			return
		}
		h.arr[parent], h.arr[index] = h.arr[index], h.arr[parent]
		index = parent
	}
}

func (h *MinHeap) Remove() (int, bool) {
	if len(h.arr) == 0 {
		return 0, false
	}
	value := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.popDown(0)
	return value, true
}

func (h *MinHeap) popDown(index int) {
	for {
		left := 2*index + 1
		right := 2*index + 2
		minIndex := index
		if left < len(h.arr)-1 && h.arr[left] < h.arr[minIndex] {
			minIndex = left
		}
		if right < len(h.arr)-1 && h.arr[right] < h.arr[minIndex] {
			minIndex = right
		}
		if index == minIndex {
			return
		}
		h.arr[index], h.arr[minIndex] = h.arr[minIndex], h.arr[index]
		index = minIndex
	}
}
