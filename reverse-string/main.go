package main

import "fmt"

func main(){
	test := "kakao = jae kyung"
	fmt.Println(reverse(test))

}

func reverse(str string) string{
	buffer := []rune(str)
	for i, j:=0, len(str)-1; i < len(str)/2; i, j = i+1, j-1{
		buffer[i], buffer[j] = buffer[j], buffer[i]
	}
	return string(buffer)
}