package main

import "fmt"




type HouseTemplate interface{
	buildHouse()string
	buildWindow()string
	buildWalls() string
	buildPillars()string
	buildFoundation()string
}


func buildFoundation()string{
	return "Foundation is build using steel and iron and cement"
}

func buildWalls()string{
	return "Wall is built using bricks and cement"
}


func ConstructHouse1(h HouseTemplate){
	fmt.Println(h.buildFoundation())
	fmt.Println(h.buildWalls())
}









