package lecture7

import "golang.org/x/exp/constraints"

/*
Stairs: |-0-|-1-|-2-|-3-|
Prices: |-0-|-3-|-2-|-4-|

Description:
	paidStaircaseCheapestPath takes index of the stair and prices of each stair.
	It returns the cheapest path to n-th stair. You are allowed to take only
	1..2 steps at a time.

Objective function:
	F(i) is the cheapest path to i-th stair.

Base cases:
	F(0) = 0
	F(1) = 3
	F(2) = 2
	F(3) = 6

Recurrence function:
	F(n) = P(n) + min(F(n-1), F(n-2))

Order: bottom-up

Solution in: F(n)
*/
func paidStaircaseCheapestPath(n int, prices []int) []int {
	from := make([]int, n+1)
	d := make([]int, n+1)
	d[0] = prices[0]
	d[1] = prices[1]
	from[1] = 0

	for i := 2; i <= n; i++ {
		d[i] = prices[i] + min(d[i-1], d[i-2])
		if d[i-1] < d[i-2] {
			from[i] = i - 1
		} else {
			from[i] = i - 2
		}
	}

	var path []int
	for curr := n; curr > 0; curr = from[curr] {
		path = append(path, curr)
	}
	path = append(path, 0)

	reverse(path)

	return path
}

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
