package main

import "fmt"

func main(){
	str := "teeters"
	fmt.Println(noRepeat(str))
}

func noRepeat(A string) string {
	s := make(map[rune]int, len(A))
	for _, r := range A{
		s[r]++
	}
	for _, v := range A {

		fmt.Printf("%T\n", v)
		if s[v] == 1{
			return string(v)
		}

	}
	return ""
}
