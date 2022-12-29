/* Binary search can be implemented on sorted array elements.
If the list elements are not arranged in a sorted manner, we have first to sort them. */

/*
	time complexity

Best case - O(1) -> target in first index of array
Worst case - O(logn) -> target in last index of array
*/

package main

import "fmt"

func findTarget(input []int, target int, startIndex int, endIndex int) int {
	if endIndex >= startIndex {
		mid := (startIndex + endIndex) / 2 // middle index of array

		if input[mid] == target { // target in middle index of array
			return mid + 1
		} else if input[mid] > target { // if the item to be searched is smaller than middle, then it can only be in left subarray
			return findTarget(input, target, startIndex, mid-1)
		} else { // if the item to be searched is greater than middle, then it can only be in right subarray
			return findTarget(input, target, mid+1, endIndex)
		}
	}
	return -1
}

func main() {
	target := 3
	input := []int{1, 2, 3, 4, 5}
	length := len(input)
	startIndex := 0
	endIndex := length - 1
	result := findTarget(input, target, startIndex, endIndex)
	if result == -1 {
		fmt.Println("Target not found")
	} else {
		fmt.Printf("The target found in %dth postion", result)
	}
}
