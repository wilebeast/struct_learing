package merge_unsorted_num

import (
	"slices"
	"testing"
)

func TestSortAndCountInversions(t *testing.T) {
	tests := []struct {
		name       string
		input      []int
		wantCount  int
		wantSorted []int
	}{
		{
			name:       "empty slice",
			input:      []int{},
			wantCount:  0,
			wantSorted: []int{},
		},
		{
			name:       "single element",
			input:      []int{42},
			wantCount:  0,
			wantSorted: []int{42},
		},
		{
			name:       "already sorted",
			input:      []int{1, 2, 3, 4, 5},
			wantCount:  0,
			wantSorted: []int{1, 2, 3, 4, 5},
		},
		{
			name:       "reverse sorted",
			input:      []int{5, 4, 3, 2, 1},
			wantCount:  10,
			wantSorted: []int{1, 2, 3, 4, 5},
		},
		{
			name:       "mixed values",
			input:      []int{2, 4, 1, 3, 5},
			wantCount:  3,
			wantSorted: []int{1, 2, 3, 4, 5},
		},
		{
			name:       "with duplicates",
			input:      []int{2, 3, 2, 3, 1},
			wantCount:  5,
			wantSorted: []int{1, 2, 2, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := append([]int(nil), tt.input...)
			got := sortAndCountInversions(&arr, 0, len(arr)-1)

			if got != tt.wantCount {
				t.Fatalf("sortAndCountInversions() = %d, want %d", got, tt.wantCount)
			}
			if !slices.Equal(arr, tt.wantSorted) {
				t.Fatalf("sortAndCountInversions() sorted array = %v, want %v", arr, tt.wantSorted)
			}
		})
	}
}
