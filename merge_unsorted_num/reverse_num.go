package merge_unsorted_num

// mergeAndCountSplitInversions merges two sorted halves and counts cross-half inversions.
func mergeAndCountSplitInversions(arr *[]int, left int, mid int, right int) int {
	// Create temporary arrays for left and right
	n1 := mid - left + 1
	n2 := right - mid
	leftArr := make([]int, n1)
	rightArr := make([]int, n2)

	// Copy data to temporary arrays
	for i := 0; i < n1; i++ {
		leftArr[i] = (*arr)[left+i]
	}
	for j := 0; j < n2; j++ {
		rightArr[j] = (*arr)[mid+1+j]
	}

	// Merge the temporary arrays back into arr and count inversions
	i, j, k := 0, 0, left
	inversionCount := 0
	for i < n1 && j < n2 {
		if leftArr[i] <= rightArr[j] {
			(*arr)[k] = leftArr[i]
			i++
		} else {
			(*arr)[k] = rightArr[j]
			inversionCount += n1 - i // Count inversions
			j++
		}
		k++
	}

	// Copy remaining elements of leftArr, if any
	for i < n1 {
		(*arr)[k] = leftArr[i]
		i++
		k++
	}

	// Copy remaining elements of rightArr, if any
	for j < n2 {
		(*arr)[k] = rightArr[j]
		j++
		k++
	}
	return inversionCount
}

// sortAndCountInversions sorts the range and returns its inversion count.
func sortAndCountInversions(arr *[]int, left int, right int) int {
	if left >= right {
		return 0
	}
	mid := left + (right-left)/2

	// Count inversions in left half, right half, and during merge
	inversions := sortAndCountInversions(arr, left, mid) +
		sortAndCountInversions(arr, mid+1, right)

	// Merge the two halves and count inversions during merge
	inversions = inversions + mergeAndCountSplitInversions(arr, left, mid, right)

	return inversions
}
