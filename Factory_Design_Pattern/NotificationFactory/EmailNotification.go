package main

import "fmt"

type Email struct{
}


func(e *Email)sendNotification(receiver string,message string){
	fmt.Printf("hi %s This is to inform you that we have noticied that you are willing to opt for Wifi Service %s",receiver,message)
}