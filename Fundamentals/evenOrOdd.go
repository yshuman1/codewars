package main

// Create a function (or write a script in Shell) that takes an integer as an argument and
// returns "Even" for even numbers or "Odd" for odd numbers.

func EvenOrOdd(number int) string {
	// your code here
	if number%2 != 0 {
		return "Odd"
	} else {
		return "Even"
	}
}
