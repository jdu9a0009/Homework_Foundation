package main

import "fmt"

func main() {
	input := []string{"a12bc3", "12a", "ab3"}
	output := []string{}
	for _, v := range input {
		numbers := ""
		for _, char := range v {
			if char >= 0 && char <= 9 {
				numbers += string(char)
			}
		}

		output = append(output, numbers)
	}
	fmt.Println(output)
}
