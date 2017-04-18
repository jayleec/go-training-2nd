package main

import (
	"math/rand"
	"fmt"
)

type Tree struct{
	Left *Tree
	Value int
	Right *Tree
}

func New(k int) *Tree{
	var t *Tree
	for _ , v := range rand.Perm(10){
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree{
	if t == nil{
		return &Tree{nil, v, nil}
	}
	if v < t.Value{
		t.Left = insert(t.Left, v)
	}else{
		t.Right = insert(t.Right, v)
	}
	return t
}

func Walk(t *Tree, ch chan int){
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *Tree) bool{
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func(){
		Walk(t1, ch1)
		close(ch1)
	}()
	go func(){
		Walk(t2, ch2)
		close(ch2)
	}()
	for n := range ch1{
		if n != <-ch2{
			return false
		}
	}
	return true
}

func main(){
	//t := New(10)
	//ch := make(chan int)
	//go func(){
	//	Walk(t, ch)
	//	close(ch)
	//}()
	//for n := range ch{
	//	fmt.Println(n)
	//}

	t1 := New(10)
	t2 := New(10)

	fmt.Println(Same(t1, t2))

}