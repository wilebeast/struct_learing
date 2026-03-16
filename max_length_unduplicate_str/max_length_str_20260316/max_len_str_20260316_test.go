package max_length_str_20260316

import "testing"

func Test_maxLengthStr(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				input: "abcdefgd",
			},
			want: "abcdefg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLengthStr(tt.args.input); got != tt.want {
				t.Errorf("maxLengthStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
