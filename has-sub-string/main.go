package main

import (
	"fmt"
	"strings"
)

func main(){
	test1 := "kakao jae kyung"
	test2 := "not"
	fmt.Println(HasSubStr(test1, test2))

}

//부분 문자열 포함 여부
func HasSubStr(origin, sub string) bool {
	for i:=0; i<len(origin); i++{
		if hasPrefix(origin[i:], sub){
			return true
		}
	}
	return false
}

func hasPrefix(s, prefix string) bool{
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func test(){
	strings.HasPrefix()
}