/**
 * url: https://leetcode-cn.com/problems/group-anagrams/
 * id: 49
 *
 * 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
 *
 * 示例 1:
 *
 * 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * 输出:
 * [
 *   ["ate","eat","tea"],
 *   ["nat","tan"],
 *   ["bat"]
 * ]
 *
 * 说明：
 *
 * 所有输入均为小写字母。
 * 不考虑答案输出的顺序。
 *
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat","tea","tan","ate","nat","bat"}

	fmt.Printf("groupAnagrams1:\n")
	arr := groupAnagrams1(strs)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]) ; j++ {
			fmt.Printf("%s ", arr[i][j])
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\ngroupAnagrams2:\n")
	arr = groupAnagrams2(strs)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]) ; j++ {
			fmt.Printf("%s ", arr[i][j])
		}
		fmt.Printf("\n")
	}
}

func groupAnagrams1(strs []string) [][]string {
	table := [26]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	flag := make(map[int]int)
	group := make([][]string, len(strs))

	size := len(strs)
	index := 0
	for i := 0; i < size; i++ {
		//计算key
		key := 1
		for _,value := range strs[i] {
			key *= table[int(value - 'a')]
		}

		//分组
		if flag[key] == 0 {
			index++
			flag[key] = index
		}
		group[flag[key] - 1] = append(group[flag[key] - 1], strs[i])
		//fmt.Printf("flag[%d] = %d\n", curSum, flag[curSum])
	}

	arr := make([][]string, index)
	for i := 0; i < index; i++ {
		arr[i] = group[i]
	}

	return arr
}

func groupAnagrams2(strs []string) [][]string {
	flag := make(map[string]int)
	group := make([][]string, len(strs))

	size := len(strs)
	index := 0
	for i := 0; i < size; i++ {
		//计算key
		key := ""
		arr := make([]string, len(strs[i]))
		for j := 0; j < len(strs[i]); j++ {
			arr = append(arr, strs[i][j:j+1])
		}
		sort.Strings(arr)
		for _, value := range arr {
			key += value
		}

		//分组
		if flag[key] == 0 {
			index++
			flag[key] = index
		}
		group[flag[key] - 1] = append(group[flag[key] - 1], strs[i])
		//fmt.Printf("flag[%d] = %d\n", curSum, flag[curSum])
	}

	arr := make([][]string, index)
	for i := 0; i < index; i++ {
		arr[i] = group[i]
	}

	return arr
}

/**
 思路1： 素数相乘作为hash key
 思路2： 升序字符串作为hash key
 */