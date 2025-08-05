package main

import "fmt"

func main(){
	for i:=0;i<=15;i++{
		go getInstance()
	}

	fmt.Scan()
}