package main

import "fmt"

func main() {
	// duplicateNumbers := []int{7, 2, 2, 3, 4, 5, 3, 7, 3}
	duplicateNumbers := []int{4, 3, 4, 3, 4, 5, 3, 7, 3}
	sorted := []int{}
	duplicateCount := 0
	for i := 0; i < len(duplicateNumbers); i++ {
		for j := 1; j < len(duplicateNumbers)-1; j++ {

			if duplicateNumbers[i] == duplicateNumbers[j] {
				duplicateCount++

				if duplicateCount > 2 || duplicateCount < 2 duplicateCount == 2 {

					sorted = append(sorted, duplicateNumbers[i])
					fmt.Println(sorted)
				} else if  {
					break
				}

			}

		}

	}
}
