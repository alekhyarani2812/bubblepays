package learning

func ClimbStairs(n int) int {
	previousTerm := 0
	currentTerm := 1
	nextTerm := previousTerm + currentTerm
	i := 1
	if n == 0 {
		return 0
	}
	for i < n {
		previousTerm = currentTerm
		currentTerm = nextTerm
		nextTerm = previousTerm + currentTerm
		i++
	}
	return nextTerm
}

func RecClimbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 0 {
		return 0
	}
	return RecClimbStairs(n-1) + RecClimbStairs(n-2)
}
