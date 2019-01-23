// This function has to be called multiply and needs to take two numbers as arguments,
// and has to return the multiplication of the two arguments.

package main

import "fmt"

func Multiply(x, y int) int {
	return x * y
}

func main(){
fmt.Println("should be 30: ",Multiply(5,6))
}