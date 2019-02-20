//Return a new array consisting of elements which are multiple of their own index in input array (length > 1).

package main

func multipleOfIndex(ints []int) []int {
	// good luck
	arr := make([]int, 0)
	for i, v := range ints {
		if i > 0 {
			if v%i == 0 {
				arr = append(arr, v)
			}
		}
	}
	return arr
}

// func main() {
// 	fmt.Println(multipleOfIndex([]int{22, -6, 32, 82, 9, 25}))
// }
