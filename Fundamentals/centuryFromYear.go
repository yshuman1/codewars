//The first century spans from the year 1 up to and including the year 100, The second - from the year 101 up to and including the year 200, etc.

//Given a year, return the century it is in.

package main

func century(year int) int {
	if year <= 100 {
		return 1
	}
	if year%100 != 0 {
		return year/100 + 1
	}
	return year / 100

	// Finish this :)

}
