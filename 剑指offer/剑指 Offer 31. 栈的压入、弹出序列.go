package main

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	i := 0
	for _, num := range pushed {
		stack = append(stack, num)
		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}
	return len(stack) == 0

}
