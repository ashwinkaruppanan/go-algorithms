// preformed with sorted and unsorted array
/*
	time complexity

Best case - O(1) -> target in first index of array
Worst case - O(n) -> target in last index of array
*/

package main

import "fmt"

func findTarget(input []int, length int, target int) int {
	for i := 0; i < length; i++ {
		if input[i] == target {
			return i + 1
		}
	}
	return -1
}

func main() {
	target := 9

	input := []int{2, 5, 0, 3, 1, 6, 9}
	length := len(input)
	result := findTarget(input, length, target)
	if result == -1 {
		fmt.Println("Target not found")
	} else {
		fmt.Printf("Target found in %dth position", result)
	}
}
