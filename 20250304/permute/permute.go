package permute

func permute(nums []int) [][]int {
	ans := [][]int{}
	used := make([]bool, len(nums))
	backtrack(nums, []int{}, &ans, used)
	return ans
}

func sortPermute(nums []int) [][]int {
	ans := [][]int{}
	sortBacktrack(nums, []int{}, &ans, 0)
	return ans
}

func sortBacktrack(nums, path []int, ans *[][]int, index int) {
	if index == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
	}

	for i := index; i < len(nums); i++ {
		if len(path) == 0 || nums[i] > path[len(path)-1] {
			path = append(path, nums[i])
			sortBacktrack(nums, path, ans, index+1)
			path = path[:len(path)-1]
		}
	}
}

func backtrack(nums []int, path []int, ans *[][]int, used []bool) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backtrack(nums, path, ans, used)
		path = path[:len(path)-1]
		used[i] = false
	}
}
