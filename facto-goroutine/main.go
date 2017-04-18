package main

import (
	"fmt"
	"runtime"
)

func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main(){
	c := facto(10)
	for n := range c{
		fmt.Println(n)
	}
}

func facto(n int) chan int{
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i >0; i--{
			total *= i
		}
		out<-total
		close(out)
	}()
	return out
}

//
//package main
//
//import "fmt"
//
//func facto(n int, c chan int) {
//	total := 1
//	for i := n; i>0; i--{
//		total *= i
//	}
//	c <- total
//	close(c)
//}
//
//func main(){
//	c := make(chan int, 10)
//	go facto(cap(c), c)
//	for n := range c{
//		fmt.Println(n)
//	}
//}



//
//func main(){
//	c := gen()
//	out := factorial(c)
//	for n := range out {
//		fmt.Println(n)
//	}
//}
//
//func gen() chan int{
//	out := make(chan int)
//
//	go func(){
//		for i:=0; i<10; i++{
//			for j := 3; j<13; j++{
//				out <- j
//			}
//		}
//		close(out)
//	}()
//	return out
//}
//
//func factorial(in chan int) chan int{
//	out  := make(chan int)
//	go func(){
//		for n := range in{
//			out <- facto(n)
//		}
//		close(out)
//	}()
//	return out
//}
//
//func facto(n int) int{
//	total :=1
//
//	for i := n; i>0; i--{
//		total *= i
//	}
//	return total
//}