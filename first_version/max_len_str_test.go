package first_version

import "testing"

func TestMaxLenStr(t *testing.T) {
	type args struct {
		str []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				str: []byte("11"),
			},
		},
		{
			name: "abcdefgcda",
			args: args{
				str: []byte("abcdefgcda"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MaxLenStr(tt.args.str)
		})
	}
}

func TestMaxLenStr2(t *testing.T) {
	type args struct {
		str []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				str: []byte("11"),
			},
		},
		{
			name: "abcdefgcda",
			args: args{
				str: []byte("abcdefgcda"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MaxLenStrFinal(tt.args.str)
		})
	}
}

func TestMaxLenStrForce(t *testing.T) {
	type args struct {
		str []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				str: []byte("11"),
			},
		},
		{
			name: "abcdefgcda",
			args: args{
				str: []byte("abcdefgcda"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MaxLenStrForce(tt.args.str)
		})
	}
}
