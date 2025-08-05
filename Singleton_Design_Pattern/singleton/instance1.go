package main

import (
	"fmt"
	"sync"
)

type single struct{

}
var singleInstance *single
var lock = &sync.Mutex{}


func getInstance() *single{
	if singleInstance==nil{
		lock.Lock()
        defer lock.Unlock()
		if singleInstance==nil{
			fmt.Println("Creating a single Instance now")
			singleInstance=&single{}
		}else{
			fmt.Println("The instance is already created")
		}
	}else{
		fmt.Println("The instance is already created")
	}
	return singleInstance
}
