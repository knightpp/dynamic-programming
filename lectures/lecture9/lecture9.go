package lecture9

import "example.com/dp/utils"

/*
Description:
	You have a matrix with W width and H height. Find how many ways there are
	to the cell (x, y). You are allowed to step by one cell at a time.
	You can move down or to the right.

Objective function:
	F(x,y) is a function that returns total number of ways to the cell at x, y.

Base cases:
	When w,h = 1,1
		F(0,0) = 1
	When w,h = 2,2
		F(1,0) = 1
		F(0,1) = 1
		F(1,1) = F(1,0) + F(0,1) = 2
	When w,h = 3,3
		F(2,0) = 1
		F(0,2) = 1
		F(2,2) = F(2,1) + F(1,2) = (F(2,0) + F(1,1)) + (F(1,1)+F(0,2)) =
			= (1+2)+(2+1) = 6
		|-|-|-|
		|-|-|-|
		|-|-|-|

Recurrence function:
	F(x,y) = F(x-1, y) + F(x, y-1)

Order:
	bottom-up

Solution in:
	F(x,y)
*/
func matrixN(w, h, x, y int) int {
	d := make([][]int, h)
	for i := range d {
		d[i] = make([]int, w)
	}

	d[0][0] = 1

	for hi := range d {
		for wi := range d[hi] {
			if hi == 0 && wi == 0 {
				continue
			}

			var (
				upper int
				left  int
			)

			// row
			if hi-1 >= 0 {
				upper = d[hi-1][wi]
			}
			if wi-1 >= 0 {
				// col
				left = d[hi][wi-1]
			}

			d[hi][wi] = upper + left
		}
	}

	return d[x][y]
}

/*
Description:
	You have a matrix with W width and H height. Find how many ways there are
	to the cell (x, y). You are allowed to step by one cell at a time.
	You can move down or to the right. You must not step on banned cells.

Objective function:
	F(x,y) is a function that returns total number of ways to the cell at x, y.

Recurrence function:
	F(x,y) = (F(x-1, y) + F(x, y-1))*Ban(x,y)

Order:
	bottom-up

Solution in:
	F(x,y)
*/
func matrixBan(w, h, x, y int, banned [][]bool) int {
	d := make([][]int, h)
	for i := range d {
		d[i] = make([]int, w)
	}

	d[0][0] = 1

	for hi := range d {
		for wi := range d[hi] {
			if hi == 0 && wi == 0 {
				continue
			}

			var (
				upper int
				left  int
			)

			// row
			if hi-1 >= 0 {
				upper = d[hi-1][wi]
			}
			if wi-1 >= 0 {
				// col
				left = d[hi][wi-1]
			}

			ban := 1
			if banned[hi][wi] {
				ban = 0
			}

			d[hi][wi] = (upper + left) * ban
		}
	}

	return d[x][y]
}

/*
Description:
	You have a matrix with W width and H height. Find the maximum count of coins
	you can obtain by traversing to the cell (x,y). You are allowed to step
	by one cell at a time. You can move down or to the right.

Objective function:
	F(x,y) is a function that returns the max number of coins that can be
	collected while traversing to the cell at x, y.

Recurrence function:
	F(x,y) = max(F(x-1, y),  F(x, y-1)) + C(x,y)

Order:
	bottom-up

Solution in:
	F(x,y)
*/
func matrixCoins(w, h, x, y int, coins [][]int) int {
	d := make([][]int, h)
	for i := range d {
		d[i] = make([]int, w)
	}

	for x := range d {
		for y := range d[x] {
			var (
				upper int
				left  int
			)

			if x-1 >= 0 {
				upper = d[x-1][y]
			}
			if y-1 >= 0 {
				left = d[x][y-1]
			}

			d[x][y] = utils.Max(upper, left) + coins[x][y]
		}
	}

	return d[x][y]
}
