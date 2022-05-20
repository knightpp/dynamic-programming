package lecture7

import "example.com/dp/utils"

/*
Stairs: |-0-|-1-|-2-|-3-|
Prices: |-0-|-3-|-2-|-4-|

Description:
	paidStaircase takes index of the stair and prices of each stair. It returns
	the cheapest path to n-th stair. You are allowed to take only 1..2 steps at
	a time.

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
func paidStaircase(n int, prices []int) int {
	d := make([]int, n+1)
	for i := 0; i <= n; i++ {
		d[i] = prices[i]
		if i < 2 {
			continue
		}
		d[i] = prices[i] + utils.Min(d[i-1], d[i-2])
	}
	return d[n]
}
