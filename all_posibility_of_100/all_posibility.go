package main

import "fmt"

var results [][]int

var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {

	result := make([]int, 0, 10)

	deepSearch(&result, 0)

	for _, result := range results {

		fmt.Println(result)

	}

}

func deepSearch(result *[]int, index int) {

	if index >= len(nums) {

		//if len(*result) >= 7 && (*result)[0] == 1 && (*result)[1] == 2 && (*result)[2] == 34 && (*result)[3] == -5 &&
		//	(*result)[4] == 67 && (*result)[5] == -8 && (*result)[6] == 9 {
		//	fmt.Println(*result)
		//}

		// cal the result and check if 100

		sum := 0

		for _, num := range *result {

			sum = sum + num

		}

		if sum == 100 {
			temp := make([]int, len(*result))
			copy(temp, *result)
			results = append(results, temp)
		}

		// (*result) = make([]int, 10)

		return

	}

	// +

	*result = append(*result, nums[index])

	deepSearch(result, index+1)

	*result = (*result)[0 : len(*result)-1]

	// -

	*result = append(*result, -nums[index])

	deepSearch(result, index+1)

	*result = (*result)[0 : len(*result)-1]

	// append

	if len(*result) > 0 {

		lastNum := (*result)[len(*result)-1]

		if lastNum > 0 {

			(*result)[len(*result)-1] = (*result)[len(*result)-1]*10 + nums[index]

		} else {

			(*result)[len(*result)-1] = (*result)[len(*result)-1]*10 - nums[index]

		}

		deepSearch(result, index+1)

		(*result)[len(*result)-1] = lastNum

	}

}
