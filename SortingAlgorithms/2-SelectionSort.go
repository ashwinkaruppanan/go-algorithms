//https://www.youtube.com/watch?v=9oWd4VJOwr0
//https://www.programiz.com/dsa/selection-sort

// time complexity O(n^2) in both best and worst case

package main

import "fmt"

func SelectionSort(array []int) {
	length := len(array)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if array[min] > array[j] {
				min = j
			}
		}
		if min != i {
			array[min], array[i] = array[i], array[min]
		}
	}
}

func main() {
	array := []int{20, 12, 10, 15, 2}

	SelectionSort(array)
	fmt.Println(array)
}
