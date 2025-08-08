package main

import (
	"fmt"
)


type Hotel struct{
	Wifi string
	BreakFast string
	SeaView string
	LateCheckout string
}

type HotelBuilder interface{
	addWifi(string)*HotelBuilder
	addBreakfast(string)*HotelBuilder
	addSeaView(string)*HotelBuilder
	addLateCheckOut(string)*HotelBuilder
	build()Hotel
}

type BasicRoom struct{
	hotel Hotel
}

func(b *BasicRoom)addWifi(wifi string)*BasicRoom{
	b.hotel.Wifi=wifi
	fmt.Printf("The Wifi speed for the Basic room is %v",b.hotel.Wifi)
	return b
}

func(b *BasicRoom)addBreakfast(vegStaticMeal string)*BasicRoom{
	b.hotel.BreakFast=vegStaticMeal
	fmt.Printf("The BreakFast for the Basic room is %s",b.hotel.BreakFast)
	return b
}


func(b *BasicRoom)addSeaView(view string)*BasicRoom{
	b.hotel.SeaView=view
	fmt.Println("The Sea view is not applicable for the Basic Room")
	return b
}

func(b *BasicRoom)addLateCheckOut(checkout string)*BasicRoom{
	fmt.Println("The LateCheckout is also not applicable for the basic room")
	return b
}
func(b *BasicRoom) build()*Hotel{
	return &b.hotel
}