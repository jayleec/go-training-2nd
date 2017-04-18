package main

import (
	"fmt"
	"strings"
)

func main(){
	str1 := "KAKAO2"
	str2 := "KAKAO23"

	fmt.Println(strings.Contains(str1, "c"))
	fmt.Println(sameStr(str1, str2))
	fmt.Println(hasSubStr(str1, str2))
	fmt.Println(1 == 2)
}

func hasSubStr(A, B string) bool {
	i := 0
	for i <= len(A){
		if strings.HasPrefix(A[i:], B) {
			return true
		}
		i++
	}
	return false
}

func sameStr(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i:=0; i<len(a); i++{
		if a[i] != b[i]{
			return false
		}
	}
	return true
}