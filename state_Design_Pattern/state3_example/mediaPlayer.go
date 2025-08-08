package main

import (
	"fmt"
)


type MediaPlayer interface{
	player(*Button)
}


type Button struct{
	mediaplayer MediaPlayer
}

func(b *Button)setState(media1 MediaPlayer){
	b.mediaplayer=media1
}

func (b *Button)player(){
	b.mediaplayer.player(b)
}


type Playing struct{}


func(p *Playing)player(b *Button){
	b.setState(&Paused{})
	fmt.Println("The Mediaplayer is paused")
}

type Paused struct{}

func(p *Paused)player(b *Button){
	
	fmt.Println("The MediaPlayer is stopped")
	b.setState(&Stopped{})    
}

type Stopped struct{

}  

func(s *Stopped)player(b *Button){
	if(b.mediaplayer==&Stopped{}){
		fmt.Println("the Button is already stopped")
	}
	fmt.Println("The Media Player is  Playing")
	b.setState(&Playing{})
}


func main(){
	b:=&Button{}

	b.setState(&Stopped{})
	b.player()
	b.player()
	b.player()
}



