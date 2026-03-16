package main

import "fmt"

func main() {
	result := make([]int, 0, 10)
	handleNext(0, result)
}

func handleNext(index int, result []int) {
	if index == 9 {
		sum := 0
		for i := 0; i < len(result); i++ {
			sum += result[i]
		}
		if sum == 100 {
			fmt.Println(result)
		}
		return
	}

	// -
	result = append(result, -1*(index+1))
	handleNext(index+1, result)
	result = result[:len(result)-1]
	// +
	result = append(result, index+1)
	handleNext(index+1, result)
	result = result[:len(result)-1]

	if index > 0 {
		// append
		result[len(result)-1] = result[len(result)-1]*10 + (index + 1)
		handleNext(index+1, result)
		result[len(result)-1] = result[len(result)-1] / 10
	}
}
