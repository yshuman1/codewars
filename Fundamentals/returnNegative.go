//In this simple assignment you are given a number and have to make it negative. But maybe the number is already negative?
// Notes:
//
//The number can be negative already, in which case no change is required.
//Zero (0) is not checked for any specific sign. Negative zeros make no mathematical sense.

package main

import "fmt"

func MakeNegative(x int) int {
	if x <= 0 {
		return x
	} else {
		return -x
	}

}

func main() {
fmt.Println("should be negative 5", MakeNegative(5))
}
