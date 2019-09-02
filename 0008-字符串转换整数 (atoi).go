/**
 * url: https://leetcode-cn.com/problems/string-to-integer-atoi/
 * id: 8
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d", myAtoi("+ 002147483646"))
}

func myAtoi(str string) int {
	//有效字符串
	symbol, start, end := 0, -1, len(str)
	for i := 0; i < len(str); i++ {
		if start < 0 {
			if symbol == 0 {
				if str[i:i+1] == " " {
					continue
				}
				if str[i:i+1] == "-" {
					symbol = -1
					continue
				}
				if str[i:i+1] == "+" {
					symbol = 1
					continue
				}
			}
			if str[i:i+1] >= "0" && str[i:i+1] <= "9" {
				start = i
				continue
			}
			break
		} else {
			if str[i:i+1] < "0" || str[i:i+1] > "9" {
				end = i
				break
			}
		}
	}

	//计算
	validStr, index := "", 0
	if start >= 0 {
		validStr = str[start:end]
		fmt.Printf("'%s'\n", validStr)
		for index = 0; index < len(validStr); index++ {
			if validStr[index:index+1] != "0" {
				validStr = validStr[index:]
				break
			}
		}
		fmt.Printf("'%s'\n", validStr)
		if index == len(validStr) && validStr[index - 1] == '0' {
			return 0
		}
	}


	if symbol >= 0 && (len(validStr) > 10 || (len(validStr) == 10 && validStr >= "2147483647")) {
		return 2147483647
	}
	if symbol < 0 && (len(validStr) > 10 || (len(validStr) == 10 && validStr >= "2147483648")) {
		return -2147483648
	}

	sum, position := 0, 1
	for i := len(validStr) - 1; i >= 0; i-- {
		num := int(validStr[i] - '0') * position
		if symbol >= 0 {
			sum += num
		} else {
			sum -= num
		}
		position *= 10
	}
	return sum
}