//-> https://www.youtube.com/watch?v=o4bAoo_gFBU&t=1984s
//-> https://www.programiz.com/dsa/bubble-sort

// time complexity
// best case -> O(n)
// worst case -> O(n^2)

package main

import "fmt"

func BubbleSort(array []int) {
	length := len(array)
	var isSorted bool
	for i := 0; i < length; i++ {
		isSorted = true
		for j := 1; j < length-i; j++ {
			if array[j] < array[j-1] { // if this condition is false till the end of 'j' loop the array is sorted
				array[j], array[j-1] = array[j-1], array[j]
				isSorted = false
			}
			//fmt.Println(array, i, j)
		}
		if isSorted {
			return
		}
	}
}

func main() {
	array := []int{2, 3, 4, 5, 6, 7, 7}

	BubbleSort(array)
	fmt.Println(array)
}
