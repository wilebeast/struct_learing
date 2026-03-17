package beautiful_arrangement

import "fmt"

// CountArrangement returns the number of beautiful arrangements for 1..n.
// A permutation is beautiful when for every 1-based position pos:
// value%pos == 0 || pos%value == 0.
func CountArrangement(n int32) int32 {
	return int32(len(GenerateArrangements(n)))
}

// GenerateArrangements returns all beautiful arrangements for 1..n.
func GenerateArrangements(n int32) [][]int32 {
	if n <= 0 {
		return nil
	}

	used := make([]bool, int(n)+1)
	path := make([]int32, 0, int(n))
	result := make([][]int32, 0)
	backtrack(1, n, used, path, &result)
	return result
}

// PrintArrangements prints every beautiful arrangement for 1..n.
func PrintArrangements(n int32) {
	arrangements := GenerateArrangements(n)
	for _, arrangement := range arrangements {
		fmt.Println(arrangement)
	}
}

func backtrack(pos, n int32, used []bool, path []int32, result *[][]int32) {
	if pos > n {
		temp := make([]int32, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	for num := int32(1); num <= n; num++ {
		if used[int(num)] {
			continue
		}
		if num%pos != 0 && pos%num != 0 {
			continue
		}

		used[int(num)] = true
		path = append(path, num)
		backtrack(pos+1, n, used, path, result)
		path = path[:len(path)-1]
		used[int(num)] = false
	}
}
