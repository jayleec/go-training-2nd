package main

import "fmt"

func main(){
	testCase := []int{1,3,2,4,5}
	fmt.Println(findSwapped(testCase))
}

func findSwapped(A []int) []int{
	result := make([]int, 0)
	for i := 1; i < len(A); i++{
		if Max(A[:i]) > A[i] || Min(A[i:]) < A[i]{
			result = append(result, A[i])
		}
	}
	return result
}

func Max(A []int) int{
	max := A[0]
	for i:=1; i<len(A); i++{
		if max < A[i] {
			max = A[i]
		}
	}
	return max
}

func Min(A []int) int{
	min := A[0]
	for i:=1; i< len(A); i++{
		if min > A[i]{
			min = A[i]
		}
	}
	return min
}