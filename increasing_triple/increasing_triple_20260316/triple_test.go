package increasing_triple_20260316

import "testing"

func Test_increasingTriplet(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingTriplet(tt.args.input); got != tt.want {
				t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}
