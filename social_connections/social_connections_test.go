package social_connections

import (
	"reflect"
	"testing"
)

func TestGetVisibleProfilesCount(t *testing.T) {
	tests := []struct {
		name    string
		nodes   int32
		u       []int32
		v       []int32
		queries []int32
		want    []int32
	}{
		{
			name:    "example from prompt",
			nodes:   7,
			u:       []int32{1, 2, 3, 5},
			v:       []int32{2, 3, 4, 6},
			queries: []int32{1, 3, 5, 7},
			want:    []int32{4, 4, 2, 1},
		},
		{
			name:    "sample case 0",
			nodes:   5,
			u:       []int32{2, 2, 1, 1},
			v:       []int32{1, 3, 3, 4},
			queries: []int32{4, 2, 5},
			want:    []int32{4, 4, 1},
		},
		{
			name:    "isolated nodes only",
			nodes:   4,
			u:       nil,
			v:       nil,
			queries: []int32{1, 2, 4},
			want:    []int32{1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetVisibleProfilesCount(tt.nodes, tt.u, tt.v, tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("GetVisibleProfilesCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
