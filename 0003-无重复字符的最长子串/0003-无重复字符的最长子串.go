/**
 * url: https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
 * id: 3
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 * 示例 1:
 *
 * 输入: "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 * 示例 2:
 *
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 * 示例 3:
 *
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d ", lengthOfLongestSubstring(""))
}

func lengthOfLongestSubstring(s string) int {
	//return dp1(s, 0)
	return dp2(s)
}

func dp2(s string) int {
	m := make(map[int]int)
	left, curLen, maxLen := 0, 0, 0

	for i := 0; i < len(s); i++ {
		curLen++
		_, ok := m[int(s[i])]
		if ok == true {
			for left <= m[int(s[i])] {
				delete(m, int(s[left]))
				left++
				curLen--
			}
		}
		if curLen > maxLen {
			maxLen = curLen
		}
		m[int(s[i])] = i
	}

	return maxLen
}

func dp1(s string, max int) int {
	m := make(map[int]int)
	l := len(s)

	if max >= l {
		return max
	}

	for i := 0; i < l; i++ {
		_, ok := m[int(s[i])]
		if ok == false {
			m[int(s[i])] = i
		} else {
			if i > max {
				max = i
			}
			return dp1(s[m[int(s[i])]+1:], max)
		}
	}

	if max < l {
		return l
	} else {
		return max
	}
}

/**
 * 思路1： 打表标记, 从重复的后一个再开始
 * 思路2： 滑动窗口
 */