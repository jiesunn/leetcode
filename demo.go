package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	for {
        str := ScanLine()
		if str == "" {
			break;
		}
		arr := strings.Split(str, ",")
		size := len(arr)
		for i := 0; i < size; i++ {
			if 
		}
	}
}

func ScanLine() string {
	var c byte
	var err error
	var b []byte
	for {
		_, err = fmt.Scanf("%c", &c)
		if err != nil {
			return ""
		}
		
		if c != '\n' {
			b = append(b, c)
		} else {
			break
		}
	}
	return string(b)
}