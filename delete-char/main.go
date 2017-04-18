package main

import (
	"strings"
	"fmt"
)

func main(){
	o := "nosense"
	r := "no"
	removeStr(o,r)
}

func removeStr(A , R string) string{
	for _, v := range R{
		A = strings.Replace(A, string(v), "", -1)
	}
	fmt.Println(A)
	return A
}