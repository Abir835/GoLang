# 🧮 Go Heap Data Structure

This project demonstrates a custom implementation of the **Min Heap** data structure in Go, including fundamental operations and **Heap Sort**.

---

## 📘 What is a Heap?

A **heap** is a special tree-based data structure that satisfies the **heap property**:

- **Min Heap**: The parent node is always **less than or equal to** its children.
- **Max Heap**: The parent node is always **greater than or equal to** its children.

This implementation uses a **Min Heap** backed by a slice.

---

## 🧮 Heap Array Representation & Formulas

A binary heap can be represented using an array.

Given a node at index `i`:

| Formula        | Description               |
|----------------|---------------------------|
| `(i - 1) / 2`  | Get parent index          |
| `2 * i + 1`    | Get left child index      |
| `2 * i + 2`    | Get right child index     |

These formulas are used in all heap operations like insertion, deletion, and heapify.

---

## 🔧 Features Implemented

✅ Basic heap operations:

| Operation         | Description                                 | Time Complexity |
|------------------|---------------------------------------------|------------------|
| `parent(i)`      | Get index of parent                         | O(1)             |
| `leftChild(i)`   | Get index of left child                     | O(1)             |
| `rightChild(i)`  | Get index of right child                    | O(1)             |
| `insert(val)`    | Insert a value and reheap up                | O(log n)         |
| `remove()`       | Remove the root (min element) and reheap down | O(log n)       |
| `upHeap(index)`  | Fix heap upwards                            | O(log n)         |
| `downHeap(index)`| Fix heap downwards                          | O(log n)         |
| `swap(i, j)`     | Swap elements                               | O(1)             |
| `Sort()`         | Perform Heap Sort (ascending order)         | O(n log n)       |

---