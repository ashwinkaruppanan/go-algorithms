// https://www.youtube.com/watch?v=jlHkDBEumP0
// https://www.programiz.com/dsa/merge-sort

// Time complextiy -> O(n log n)

package main

import "fmt"

func Merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func MergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := MergeSort(items[:len(items)/2])
	second := MergeSort(items[len(items)/2:])
	return Merge(first, second)
}

func main() {
	array := []int{20, 12, 10, 15, 2}

	sorted := MergeSort(array)
	fmt.Println(sorted)

}
