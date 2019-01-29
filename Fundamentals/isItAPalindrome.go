// Write function isPalindrome that checks if a given string (case insensitive) is a palindrom.
package main

import (
	"strings"
)

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
