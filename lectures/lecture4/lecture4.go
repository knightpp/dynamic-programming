package lecture4

/*
	Description:
	There is n stairs. You are allowed to climb one or two steps at a time.
	Return the number of distinct ways you can climb the i-th stair.

	Objective function:
	F(i) is the number of distinct ways to climb the i-th stair

	Base cases:
		F(0) = 1
		F(1) = 1
		F(2) = F(1)+F(0) = 2
		F(3) = F(2)+F(1) = 3
		F(4) = F(3)+F(2) = 5
		F(5) = F(4)+F(3) = 8

	Recurrence function:
		F(i) = F(i-1) + F(i-2)
*/
func lecture4(n int) int {
	d := make([]int, n+2)
	d[0] = 1
	d[1] = 1

	for i := 2; i <= n; i++ {
		d[i] = d[i-1] + d[i-2]
	}

	return d[n]
}
