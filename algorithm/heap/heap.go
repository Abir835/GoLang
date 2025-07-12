package main

import "fmt"

type MinHeap struct {
	Data []int
}

// Get Parent index
func (h *MinHeap) parent(index int) int {
	return (index - 1) / 2
}

// Get left child index
func (h *MinHeap) leftChild(index int) int {
	return 2*index + 1
}

// Get right child index
func (h *MinHeap) rightChild(index int) int {
	return 2*index + 2
}

// Swap value
func (h *MinHeap) swap(i, j int) {
	h.Data[i], h.Data[j] = h.Data[j], h.Data[i]
}

// Insert value
func (h *MinHeap) insert(val int) {
	h.Data = append(h.Data, val)
	h.upHeap(len(h.Data) - 1)
}

// Heapify Up
func (h *MinHeap) upHeap(index int) {
	for index > 0 && h.Data[h.parent(index)] > h.Data[index] {
		h.swap(index, h.parent(index))
		index = h.parent(index)
	}
}

// Heap remove
func (h *MinHeap) remove() int {
	if len(h.Data) == 0 {
		panic("heap is empty")
	}

	root := h.Data[0]
	last := len(h.Data) - 1
	h.Data[0] = h.Data[last]
	h.Data = h.Data[:last]
	h.downHeap(0)
	return root
}

// downHeap
func (h *MinHeap) downHeap(index int) {
	size := len(h.Data)

	for {
		smallest := index
		left := h.leftChild(index)
		right := h.rightChild(index)

		if left < size && h.Data[left] < h.Data[smallest] {
			smallest = left
		}
		if right < size && h.Data[right] < h.Data[smallest] {
			smallest = right
		}

		if smallest == index {
			break
		}

		h.swap(index, smallest)
		index = smallest
	}
}

// sort
func (h *MinHeap) sort() []int {
	result := []int{}
	original := make([]int, len(h.Data))
	copy(original, h.Data)

	for len(h.Data) > 0 {
		result = append(result, h.remove())
	}
	h.Data = original
	return result
}

// Print Heap
func (h *MinHeap) printHeap() {
	fmt.Println(h.Data)
}

func main() {
	fmt.Println("Learning Heap")
	fmt.Println("==============")

	h := &MinHeap{}
	values := []int{10, 4, 15, 20, 0, 8}

	for _, val := range values {
		h.insert(val)
	}

	h.printHeap()

	fmt.Println("Removed:", h.remove())
	h.printHeap()

	// Insert again for sort test
	for _, val := range []int{10, 4, 15, 20, 0, 8} {
		h.insert(val)
	}

	sorted := h.sort()
	fmt.Println("Sorted:", sorted)
}
