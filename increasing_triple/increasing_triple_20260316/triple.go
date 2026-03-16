package increasing_triple_20260316

import "fmt"

func increasingTriplet(input []int) bool {
	result := make([]int, 0, 3)
	for i := 0; i < len(input); i++ {
		if deepSearch(result, i, input) {
			return true
		}
	}
	return false
}

func deepSearch(result []int, index int, input []int) bool {
	if len(result) >= 2 {
		if result[0] >= result[1] {
			return false
		}
	}
	if len(result) >= 3 {
		if result[1] >= result[2] {
			return false
		}
		fmt.Println(result)
		return true
	}

	for ; index < len(input); index++ {
		if deepSearch(append(result, input[index]), index+1, input) {
			return true
		}
		result = result[:len(result)-1]
	}
	return false
}
