package main

import (
    "fmt"
)

func countCharacters(words []string, chars string) int {
	var charsMap = make(map[string]int)
	charsLen := len(chars)
	for i := 0; i < charsLen; i++ {
		_, ok := charsMap[chars[i:i+1]];
		if ok {
			charsMap[chars[i:i+1]]++
		} else {
			charsMap[chars[i:i+1]] = 1;
		}
	}
	fmt.Println(charsMap)

	res := 0
    for _, word := range words {
		var wordMap = make(map[string]int)
		wordFlag, wordLen := 1, len(word)
		for i := 0; i < wordLen; i++ {
			_, ok := charsMap[word[i:i+1]];
			if !ok || charsMap[word[i:i+1]] - wordMap[word[i:i+1]] == 0 {
				wordFlag = 0;
				break;
			} else {
				_, ok := wordMap[word[i:i+1]];
				if ok {
					wordMap[word[i:i+1]]++
				} else {
					wordMap[word[i:i+1]] = 1;
				}
			}
		}
		if wordFlag == 1 {
			res += wordLen
		}
	}
	return res
}