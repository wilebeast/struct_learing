package quick_sort

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pivot := nums[0]
	var left, right []int
	for _, num := range nums[1:] {
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, pivot), right...)
}
