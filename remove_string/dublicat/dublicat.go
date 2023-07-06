package main

import "fmt"

func onlyTwoDuplicates(arr []int) []int {
	var resultArr []int
	for i := 0; i < len(arr); i++ {
		count := 0
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				count++
			}
		}
		if count == 2 {
			alreadyAdded := false
			for k := 0; k < len(resultArr); k++ {
				if resultArr[k] == arr[i] {
					alreadyAdded = true
					break
				}
			}
			if !alreadyAdded {
				resultArr = append(resultArr, arr[i])
			}
		}
	}
	return resultArr
}

func main() {
	fmt.Println(onlyTwoDuplicates([]int{1, 2, 1, 3, 9, 3, 4, 2, 9, 3, 1, 5, 5, 1, 2}))
}
