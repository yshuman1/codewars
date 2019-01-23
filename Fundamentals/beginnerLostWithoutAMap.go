/*
Given an array of integers, return a new array with each value doubled.

For example:

[1, 2, 3] --> [2, 4, 6]

For the beginner, try to use the map method - it comes in very handy quite a lot so is a good one to know.
*/
package main

import "fmt"

func Maps(x []int) []int {
	a := []int{}
	for _, num := range x {
		newNum := num * 2
		a = append(a, newNum)
	}
	return a
}

func main() {
	a := []int{1, 2, 3}
	fmt.Println("should give back [2 4 6] ", Maps(a))
}
