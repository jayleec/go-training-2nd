package main

import "fmt"

func main() {
	FizzBuzz(20)
}

func FizzBuzz(n int){
	for i:=1; i<= n; i++ {
		if i % 15 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i % 5 == 0 {
			fmt.Println(i, "Buzz")
		}else{
			fmt.Println(i)
		}
	}
}