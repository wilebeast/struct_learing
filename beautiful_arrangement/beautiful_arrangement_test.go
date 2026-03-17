package beautiful_arrangement

import (
	"reflect"
	"testing"
)

func TestCountArrangement(t *testing.T) {
	tests := []struct {
		name string
		n    int32
		want int32
	}{
		{
			name: "n=1",
			n:    1,
			want: 1,
		},
		{
			name: "n=2",
			n:    2,
			want: 2,
		},
		{
			name: "n=3",
			n:    3,
			want: 3,
		},
		{
			name: "n=4",
			n:    4,
			want: 8,
		},
		{
			name: "n=5",
			n:    5,
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountArrangement(tt.n); got != tt.want {
				t.Fatalf("CountArrangement(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestGenerateArrangements(t *testing.T) {
	want := [][]int32{
		{1, 2, 3, 4, 5},
		{1, 4, 3, 2, 5},
		{2, 1, 3, 4, 5},
		{2, 4, 3, 1, 5},
		{3, 2, 1, 4, 5},
		{3, 4, 1, 2, 5},
		{4, 1, 3, 2, 5},
		{4, 2, 3, 1, 5},
		{5, 2, 3, 4, 1},
		{5, 4, 3, 2, 1},
	}

	got := GenerateArrangements(5)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GenerateArrangements(5) = %v, want %v", got, want)
	}
}
