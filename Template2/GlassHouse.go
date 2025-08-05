package main

import "fmt"

type GlassHouse struct{
	HouseTemplate

}


func (g *GlassHouse)buildHouse()string{
	return "House structure is build wiht glass bricks"
}

func (g *GlassHouse)buildWindow()string{
	return "The Windows is made up of Sand wood"
}

func (g * GlassHouse)buildPillars()string{ 
	return "Pillar of the glass house is made of strong japnese marbel"
}

func (g *GlassHouse)buildWalls()string{
	return "Walls of the house is made up of bricke stones and cement"
}




func ConstructHouse(h GlassHouse) {

	fmt.Println(h.buildPillars())
	fmt.Println(h.buildHouse())
	fmt.Println(h.buildWalls())
	fmt.Println(h.buildWindow())
}

