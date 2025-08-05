package main

import (
	"fmt"
	"sync"
)

type single struct{
	count int64

}

var instance1 *single
var once sync.Once

func getInstance()*single{
	if instance1==nil{
		once.Do(func() {
				fmt.Println("Creating a instance once")
				instance1=&single{}
			})
	}
	return instance1
}

func (s *single)val(){
	s.count++
}


func (s *single)getCount()int64{
	return s.count
}