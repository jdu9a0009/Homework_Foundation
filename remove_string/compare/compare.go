package main

import "fmt"

func main() {
	firstArray := []int{1, 2, 4}
	secondArray := []int{1, 2, 3}

	equal := true
	for i := 0; i < len(firstArray); i++ {
		if firstArray[i] != secondArray[i] {
			equal = false
			break
		}
	}
	if equal {

		fmt.Println(firstArray)

	} else {
		firstArray = append(firstArray, secondArray...)
		fmt.Println(firstArray)
	}

}
