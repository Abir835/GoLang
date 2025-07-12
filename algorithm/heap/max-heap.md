# 📦 Go Max Heap Data Structure

This project demonstrates a custom implementation of a **Max Heap** in Go, including all key operations and **Heap Sort** functionality.

---

## 📘 What is a Max Heap?

A **Max Heap** is a complete binary tree where every parent node is **greater than or equal to** its children. The largest element is always at the root.

Used in:
- Priority queues
- Heap Sort
- Scheduling algorithms

---

## 🧮 Array Representation & Formulas

In a heap implemented with an array:

| Formula         | Description               |
|------------------|---------------------------|
| `(i - 1) / 2`    | Get parent index          |
| `2 * i + 1`      | Get left child index      |
| `2 * i + 2`      | Get right child index     |

These formulas help navigate the heap structure efficiently in constant time.

---

## ✅ Features Implemented

| Function            | Description                                | Time Complexity |
|---------------------|--------------------------------------------|------------------|
| `parent(i)`         | Get parent index                           | O(1)             |
| `leftChild(i)`      | Get left child index                       | O(1)             |
| `rightChild(i)`     | Get right child index                      | O(1)             |
| `insert(val)`       | Insert a value and heapify up              | O(log n)         |
| `remove()`          | Remove max value (root) and heapify down   | O(log n)         |
| `upHeap(index)`     | Maintain heap from bottom to top           | O(log n)         |
| `downHeap(index)`   | Maintain heap from top to bottom           | O(log n)         |
| `swap(i, j)`        | Swap two elements                          | O(1)             |
| `Sort()`            | Heap Sort (returns descending order)       | O(n log n)       |

---