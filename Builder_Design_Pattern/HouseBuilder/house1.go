package main

import "fmt"


type House struct{
	Rooms string
	Parking string
	Garden string
	Balacony string
}

type HouseBuilder struct{
	House
}

func (h *HouseBuilder) NewHouseBuilder()House{
	return h.House
}

func (h *HouseBuilder)setRooms(room string)*HouseBuilder{
	h.House.Rooms=room
	return h
}

func (h *HouseBuilder)setParking()string{
	return "Parking should be larger upto 500 square feet "
}

func (h *HouseBuilder)setBalacony(balcony string)*HouseBuilder{
	h.House.Balacony=balcony
	return h
}

func(h *HouseBuilder)setGarden(garden string)*HouseBuilder{
	h.House.Garden=garden
	return h
}


func main(){

	houseBuilder:=&HouseBuilder{}

	house1:=houseBuilder.setRooms("Luxury 3 bed room").setBalacony("Fully opened type").setGarden("200square feet of garden ").NewHouseBuilder()
	fmt.Println(house1)
}


