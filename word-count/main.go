package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(WordCount("Kakao"))
}

func WordCount(s string) map[string]int{
	result := make(map[string]int)
	str := strings.Fields(s)
	fmt.Println(str)
	for _, n := range str {
		result[n]++
	}
	return result
}