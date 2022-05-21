package lecture9

import (
	"reflect"
	"testing"
)

func Test_matrixN(t *testing.T) {
	type args struct {
		w int
		h int
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "edge case #1",
		// 	args: args{
		// 		w: 1,
		// 		h: 1,
		// 		x: 0,
		// 		y: 0,
		// 	},
		// 	want: 1,
		// },
		{
			name: "matrix 3x3",
			args: args{
				w: 3,
				h: 3,
				x: 2,
				y: 2,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixN(tt.args.w, tt.args.h, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("matrixN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matrixBan(t *testing.T) {
	type args struct {
		w      int
		h      int
		x      int
		y      int
		banned [][]bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3x3 #1",
			args: args{
				w: 3,
				h: 3,
				x: 2,
				y: 2,
				banned: [][]bool{
					{false, false, false},
					{false, true, false},
					{false, true, false},
				},
			},
			want: 1,
		},
		{
			name: "3x3 #2",
			args: args{
				w: 3,
				h: 3,
				x: 2,
				y: 2,
				banned: [][]bool{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixBan(tt.args.w, tt.args.h, tt.args.x, tt.args.y, tt.args.banned); got != tt.want {
				t.Errorf("matrixBan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matrixCoins(t *testing.T) {
	type args struct {
		w     int
		h     int
		x     int
		y     int
		coins [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3x3 #1",
			args: args{
				w: 3,
				h: 3,
				x: 2,
				y: 2,
				coins: [][]int{
					{1, 1, 1},
					{1, 5, 1},
					{1, 1, 0},
				},
			},
			want: 8,
		},
		{
			name: "3x3 #2",
			args: args{
				w: 3,
				h: 3,
				x: 2,
				y: 2,
				coins: [][]int{
					{1, 1, 5},
					{2, 1, 1},
					{2, 2, 0},
				},
			},
			want: 8,
		},
		{
			name: "3x3 #2",
			args: args{
				w: 3,
				h: 3,
				x: 2,
				y: 2,
				coins: [][]int{
					{1, 1, 5},
					{4, 1, 1},
					{2, 2, 0},
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixCoins(tt.args.w, tt.args.h, tt.args.x, tt.args.y, tt.args.coins); got != tt.want {
				t.Errorf("matrixCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matrixCoinsWithPath(t *testing.T) {
	type args struct {
		w     int
		h     int
		x     int
		y     int
		coins [][]int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []Point
	}{
		{
			name: "simple #1",
			args: args{
				w: 3,
				h: 3,
				x: 2,
				y: 2,
				coins: [][]int{
					{1, 1, 5}, // v
					{4, 1, 1}, // v
					{2, 2, 0}, // > > x
				},
			},
			want:  9,
			want1: []Point{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := matrixCoinsWithPath(tt.args.w, tt.args.h, tt.args.x, tt.args.y, tt.args.coins)
			if got != tt.want {
				t.Errorf("matrixCoinsWithPath() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("matrixCoinsWithPath() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
