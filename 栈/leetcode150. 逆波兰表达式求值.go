package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
}

func evalRPN(tokens []string) int {

	var stack []int
	for _, token := range tokens {
		switch token {
		case "+":
			stack[len(stack)-2] = stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] = stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] = stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] = stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			v, _ := strconv.ParseInt(token, 10, 64)
			stack = append(stack, int(v))
		}
		fmt.Println(stack)
	}
	return stack[0]
}
