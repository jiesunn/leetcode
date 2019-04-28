/**
 * url: https://leetcode-cn.com/problems/valid-palindrome/
 * id: 125
 *
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 *
 * 说明：本题中，我们将空字符串定义为有效的回文串。
 *
 * 示例 1:
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 *
 * 示例 2:
 * 输入: "race a car"
 * 输出: false
 */

package main

import (
	"fmt"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	if isPalindrome(s) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func isPalindrome(s string) bool {
	//符合要求的字符
	needChar := func(b byte) bool {
		return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
	}

	low, hi := 0, len(s)-1
	for low < hi {
		//转换为字节
		sbyte := s[low] | 0x20
		ebyte := s[hi] | 0x20

		//找到对应的符合要求字符
		if !needChar(sbyte) {
			low++
			continue
		}
		if !needChar(ebyte) {
			hi--
			continue
		}

		//比对
		if sbyte != ebyte {
			return false
		}

		//继续寻找对应字符
		low++
		hi--
	}

	return true
}

/**
 * 思路：比较字符串两端往里符合条件的字符
 *
 * 执行用时 : 4 ms, 在Valid Palindrome的Go提交中击败了99.68% 的用户
 * 内存消耗 : 2.8 MB, 在Valid Palindrome的Go提交中击败了86.49% 的用户
 */
