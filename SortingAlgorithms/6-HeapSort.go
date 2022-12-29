//https://www.youtube.com/watch?v=Q_eia3jC9Ts
//https://www.programiz.com/dsa/heap-sort

// Time complexity -> O(n log n)

package main

import "fmt"

func HeapSort(array []int, length int) {
	//Build max heap
	for i := length/2 - 1; i >= 0; i-- {
		Heapify(array, length, i)
	}

	//Heap sort
	for i := length - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]

		//Heapify root element
		Heapify(array, i, 0)
	}

}

func Heapify(array []int, length int, i int) {
	// Find largest among root, left child and right child
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < length && array[left] > array[largest] {
		largest = left
	}

	if right < length && array[right] > array[largest] {
		largest = right
	}

	// Swap and continue heapifying if root is not largest
	if largest != i {
		array[largest], array[i] = array[i], array[largest]
		Heapify(array, length, largest)
	}
}

func main() {
	array := []int{20, 12, 10, 15, 2}
	length := len(array)

	HeapSort(array, length)
	fmt.Println(array)

}
