package lecture7

import (
	"reflect"
	"testing"
)

func Test_paidStaircase(t *testing.T) {
	type args struct {
		n      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "edge case #1",
			args: args{
				n:      2,
				prices: []int{0, 3, 2},
			},
			want: []int{0, 2},
		},
		{
			name: "when n = 8",
			args: args{
				n:      8,
				prices: []int{0, 3, 2, 4, 6, 1, 1, 5, 3},
			},
			want: []int{0, 2, 3, 5, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paidStaircaseCheapestPath(tt.args.n, tt.args.prices); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("paidStaircase() = %v, want %v", got, tt.want)
			}
		})
	}
}
