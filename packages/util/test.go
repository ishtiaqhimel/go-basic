package util

// variadic function usage: takes two or more arguments & finds the maximum among them
func Max(a int, b int, args ...int) int {
	mx := a
	if b > mx {
		mx = b
	}
	for _, val := range args {
		if val > mx {
			mx = val
		}
	}
	return mx
}

// variadic function usage: takes two or more arguments & finds the minimum among them
func Min(a int, b int, args ...int) int {
	mn := a
	if b < mn {
		mn = b
	}
	for _, val := range args {
		if val < mn {
			mn = val
		}
	}
	return mn
}
