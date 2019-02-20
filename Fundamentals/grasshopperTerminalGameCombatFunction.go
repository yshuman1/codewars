//Create a combat function that takes the player's current health and the amount of damage recieved, and returns the player's new health. Health can't be less than 0.

package main

func combat(health, damage float64) float64 {
	if damage >= health {
		return 0
	}
	return health - damage
}

// or a cleaner example i saw after submitting is the one below. in i tthey use math.Max which gives you the greater of the two arguments.

// import "math"
// func combat(health, damage float64) float64 {
//   return math.Max(health-damage, 0.0)
// }
