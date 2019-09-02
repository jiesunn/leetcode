package main

import (
	"fmt"
)

func main() {
	v := numDecodings("1212")
	fmt.Println(v)
}

func numDecodings(s string) int {
	count := 0
	l := len(s)
	for i := 0; i < l; i++ {
		if i == 0 {
			if s[i:i+1] == "0" {
				break
			}
			count++
			continue
		}
		if s[i-1:i] != "0" && s[i:i+1] != "0" && s[i-1:i+1] <= "26" {
			count++
		}
		if i >= 2 && s[i:i+1] == "0" && s[i-2:i-1] != "0" && s[i-2:i] <= "26" {
			count--
		}
		if s[i-1:i+1] == "00" || (s[i-1:i] > "2" && s[i:i+1] == "0") {
			count = 0
			break
		}
	}
	return count
}