package merge_unsorted_num_2026037

func mergeAndCountSplitInversions(arr []int, left, mid, right int) int {
	return 0
}

func sortAndCountInversions(arr []int, left, right int) int {
	if left <= right {
		return 0
	}

	mid := (left + right) / 2

	inverseNum := sortAndCountInversions(arr, left, mid) + sortAndCountInversions(arr, mid+1, right)

	return inverseNum + mergeAndCountSplitInversions(arr, left, mid, right)
}
