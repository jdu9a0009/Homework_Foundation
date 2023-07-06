package main

import "fmt"

func findDuplicates(arr []int) []int {
	duplicates := make([]int, 0)
	visited := make(map[int]bool)

	for _, num := range arr {
		if visited[num] {
			duplicates = append(duplicates, num)
		} else {
			visited[num] = true
		}
	}

	return duplicates
}

func main() {
	numRay := []int{0, 4, 3, 2, 7, 8, 2, 3, 1, 1, 1}
	duplicates := findDuplicates(numRay)

	fmt.Println("Duplicates:", duplicates)
}
