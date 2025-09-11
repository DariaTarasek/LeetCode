package other

func ClimbStairs(n int) int {
	f1, f2 := 1, 1
	for i := 0; i < n; i++ {
		f1 = f1 + f2
		f1, f2 = f2, f1
	}
	return f1
}
