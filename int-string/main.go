package main

import (
	"fmt"
	"strconv"
)

func main(){
	nums := "-54321"
	fmt.Println(strToInt(nums))

	nums2 := -54321
	fmt.Println("string:", intToStr(-55888))
	result := strconv.Itoa(nums2)

	fmt.Printf("%v  , %T\n", result, result)

}

func intToStr(num int) string {
	isNag := false
	result := []int{}
	var str string
	if num < 0{
		isNag = true
		num = -1 * num
	}

	for num != 0{
		result = append([]int{num%10 + int('0')}, result[0:]...)
		num /= 10
	}


	for _, n := range result{
		str = str + string(n)
	}
	if isNag == true{
		return "-" + str
	}

	return str
}

func strToInt(str string) int {
	isNag := false
	num := 0

	if rune(str[0]) == '-'{
		isNag = true
		str = str[1:]
	}

	for _, v := range str{
		num *= 10
		num += int(v) - int('0')
	}
	if isNag == true{
		return -1 * num
	}
	return num
}