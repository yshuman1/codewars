// Complete the solution so that it reverses the string value passed into it.

package main

func Solution(word string) string {
	r := []rune(word)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// func main() {
// 	fmt.Println(Solution("jokes"))
// }
