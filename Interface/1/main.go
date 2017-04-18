package main

import "fmt"

type IPAddr [4]byte

func (i IPAddr) String() string{
	return	fmt.Sprintf("{%d, %d, %d, %d}", i[0], i[1], i[2], i[3])
}

func main(){
	hosts := map[string]IPAddr{
		"loopback" : {1,1,1,1},
		"dns" : {2,2,2,2},
	}
	for name, ip := range hosts {
		fmt.Printf("%v %v\n", name, ip)
	}
}
