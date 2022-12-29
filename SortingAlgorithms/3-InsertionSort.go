// https://www.youtube.com/watch?v=yCxV0kBpA6M
// https://www.programiz.com/dsa/insertion-sort

// time complexity
// best case O(n)
// worst case O(n^2)

package main

import (
	"fmt"
)

func InsertionSort(array []int) {
	for i := 0; i < len(array); i++ {
		temp := array[i]
		j := i - 1
		for j >= 0 && array[j] > temp {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = temp
	}
}

func main() {
	array := []int{20, 12, 10, 15, 2}

	InsertionSort(array)
	fmt.Println(array)
}
