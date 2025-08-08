package main

import (
	"fmt"

	"google.golang.org/genproto/googleapis/type/date"
)

func main(){

	i2:=getInstance()
	i2.update("Oil","20 liters",date.Date{Year:2030})
	fmt.Println()

	i3:=getInstance()
	i3.update("Oil","20 liters",date.Date{Year:2030})
}