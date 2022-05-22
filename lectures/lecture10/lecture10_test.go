package lecture10

import "testing"

func Test_fences(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple #1",
			args: args{
				n: 3,
			}, // 3!
			want: 6,
		},
		{
			name: "simple #2",
			args: args{
				n: 4,
			},
			want: 10,
		},
		{
			name: "simple #3",
			args: args{
				n: 5,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fences(tt.args.n); got != tt.want {
				t.Errorf("fences() = %v, want %v", got, tt.want)
			}
		})
	}
}
