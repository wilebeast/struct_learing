package first_version

func deepSearch(triplet *[]int, tripletNum int, index int, nums *[]int) bool {
	if tripletNum == 3 {
		return true
	}

	if index >= len(*nums) {
		return false
	}

	for i := 0; i < len(*nums); i++ {
		if tripletNum == 0 {
			(*triplet)[0] = (*nums)[0]
			if deepSearch(triplet, tripletNum+1, index, nums) {
				return true
			}

		} else if (*triplet)[tripletNum-1] < (*nums)[index] {
			(*triplet)[tripletNum] = (*nums)[index]
			if deepSearch(triplet, tripletNum+1, index+1, nums) {
				return true
			}
		}
		index = index + 1
	}
	return false
}

func increasingTriplet(nums []int) bool {
	triplet := make([]int, 3)
	return deepSearch(&triplet, 0, 0, &nums)
}
