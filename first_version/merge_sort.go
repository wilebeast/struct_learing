package first_version

// Merge function to merge two halves
func merge(arr *[]int, left int, mid int, right int) {
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

	// Merge the temporary arrays back into arr
	i, j, k := 0, 0, left
	for i < n1 && j < n2 {
		if leftArr[i] <= rightArr[j] {
			(*arr)[k] = leftArr[i]
			i++
		} else {
			(*arr)[k] = rightArr[j]
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
}

// MergeSort function to sort the array
func mergeSort(arr *[]int, left int, right int) {
	if left < right {
		// Find the middle point
		mid := left + (right-left)/2

		// Sort first and second halves
		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)

		// Merge the sorted halves
		merge(arr, left, mid, right)
	}
}
