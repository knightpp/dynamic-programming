package lecture10

/*
Description:
	You have to paint a fence with N posts. You have to paint posts such that
	no more than two posts are adjacent. You have 2 colors.

Objective function:
	F(n, c) is a function that returns how many ways there are to paint up to
	n'th fence and the n'th fence must be in c color.

Recurrence function:
	F(n, c) = F(n-1, 1-c) + F(n-2, 1-c)
*/
func fences(n int) int {
	d := make([][]int, n+1)
	for i := range d {
		d[i] = make([]int, 2)
	}

	d[1][blue] = 1
	d[1][green] = 1
	d[2][blue] = 2
	d[2][green] = 2

	for i := 3; i <= n; i++ {
		// iterate all colors
		for j := 0; j <= 1; j++ {
			// TODO: repeat
			d[i][j] = d[i-1][1-j] + d[i-2][1-j]
		}
	}

	return d[n][blue] + d[n][green]
}

type Color int

const (
	blue Color = iota
	green
)
