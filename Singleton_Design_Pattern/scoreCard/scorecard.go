package main

import (
	"fmt"
	"sync"
)

type ScoreCard struct{
	PlayerName string
	PlayerId int64
	GameName string
	Score int64
}

var instance *ScoreCard

var lock=&sync.Mutex{}


func getInstance()*ScoreCard{
	if(instance==nil){
		lock.Lock()
		defer lock.Unlock()
		if(instance==nil){
		fmt.Println("The ScoreCard is created successfully")
		instance=&ScoreCard{PlayerName: "John",PlayerId: 100002,GameName: "Kabadi",Score:20}
	}
}

return instance
}

func(s *ScoreCard)Update(score int64,pid int64,pname string,gname string){
	s.Score=score
	s.GameName=gname
	s.PlayerId=pid
	s.PlayerName=pname

	fmt.Print("Player Id:",s.PlayerId)
	fmt.Print("Player Name:",s.PlayerName)
	fmt.Print("Game Name:",s.GameName)
	fmt.Print("Score:",s.Score)
}