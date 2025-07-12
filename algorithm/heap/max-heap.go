package main

import "fmt"

type MaxHeap struct {
	data []int
}

// 1. Get Parent Index
func (h *MaxHeap) parent(index int) int {
	return (index - 1) / 2
}

// 2. Get Left Child Index
func (h *MaxHeap) leftChild(index int) int {
	return 2*index + 1
}

// 3. Get Right Child Index
func (h *MaxHeap) rightChild(index int) int {
	return 2*index + 2
}

// 4. Insert a value
func (h *MaxHeap) insert(val int) {
	h.data = append(h.data, val)
	h.upHeap(len(h.data) - 1)
}

// 5. Swap
func (h *MaxHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// 6. Upheap (heapify up)
func (h *MaxHeap) upHeap(index int) {
	for index > 0 && h.data[h.parent(index)] < h.data[index] {
		h.swap(index, h.parent(index))
		index = h.parent(index)
	}
}

// 7. Remove max element (root)
func (h *MaxHeap) remove() int {
	if len(h.data) == 0 {
		panic("Heap is empty")
	}
	root := h.data[0]
	last := len(h.data) - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	h.downHeap(0)
	return root
}

// 8. Downheap (heapify down)
func (h *MaxHeap) downHeap(index int) {
	size := len(h.data)
	for {
		largest := index
		left := h.leftChild(index)
		right := h.rightChild(index)

		if left < size && h.data[left] > h.data[largest] {
			largest = left
		}
		if right < size && h.data[right] > h.data[largest] {
			largest = right
		}
		if largest == index {
			break
		}
		h.swap(index, largest)
		index = largest
	}
}

// Sort 9. Heap Sort (Descending Order)
func (h *MaxHeap) Sort() []int {
	result := []int{}
	original := make([]int, len(h.data))
	copy(original, h.data)

	for len(h.data) > 0 {
		result = append(result, h.remove())
	}

	h.data = original // restore original heap
	return result
}

// Helper: Print heap
func (h *MaxHeap) printHeap() {
	fmt.Println("Heap:", h.data)
}

func main() {
	h := &MaxHeap{}
	values := []int{10, 4, 15, 20, 0, 8}

	for _, v := range values {
		h.insert(v)
	}

	h.printHeap()

	fmt.Println("Removed max:", h.remove())
	h.printHeap()

	// Insert again for sort test
	for _, v := range []int{10, 4, 15, 20, 0, 8} {
		h.insert(v)
	}

	sorted := h.Sort()
	fmt.Println("Heap Sorted (Descending):", sorted)
}
