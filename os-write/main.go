package main

import (
	"os"
	"log"
)

func main(){
	dst, err := os.Create("hello.txt")
	if err != nil{
		log.Fatalln(err.Error())
	}
	defer dst.Close()

	bs := []byte("put some phrase here")

	_, err = dst.Write(bs)
	if err != nil{
		log.Fatal("writing error", err.Error())
	}
}