package lecture7

import "testing"

func Test_paidStaircase(t *testing.T) {
	type args struct {
		n      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "edge case #1",
			args: args{
				n:      0,
				prices: []int{0},
			},
			want: 0,
		},
		{
			name: "edge case #2",
			args: args{
				n:      1,
				prices: []int{0, 3},
			},
			want: 3,
		},
		{
			name: "edge case #3",
			args: args{
				n:      2,
				prices: []int{0, 3, 2},
			},
			want: 2,
		},
		{
			name: "when n = 3",
			args: args{
				n:      3,
				prices: []int{0, 3, 2, 4},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paidStaircase(tt.args.n, tt.args.prices); got != tt.want {
				t.Errorf("paidStaircase() = %v, want %v", got, tt.want)
			}
		})
	}
}
