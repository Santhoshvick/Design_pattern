package main

import "fmt"


type SMS struct{

}


func (s *SMS)sendNotification(receiver string,message string){
	fmt.Printf("hi %s This is to inform you that we have noticied that you are willing to opt for Wifi Service %s",receiver,message)
}