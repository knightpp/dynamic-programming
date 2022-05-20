package lecture4

import "testing"

func Test_lecture4(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "edge case #1",
			args: args{
				n: 0,
			},
			want: 1,
		},
		{
			name: "edge case #2",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "when n = 3",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "when n = 4",
			args: args{
				n: 4,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lecture4(tt.args.n); got != tt.want {
				t.Errorf("lecture4(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
