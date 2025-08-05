package main

import "fmt"

func main(){
	i1:=getInstance()
	i1.val()
	i1.val()

	i2:=getInstance()
    fmt.Println()
	fmt.Println(i2.getCount())

}