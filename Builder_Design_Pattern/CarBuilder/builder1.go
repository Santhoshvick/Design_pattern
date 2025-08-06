package main

import "fmt"

type Car struct{
	Model string
	color string
	Body string
	tyre string
}



type CarBuilder struct{
	Car Car
}

func (c *CarBuilder)setModel(model string)(*CarBuilder){
	c.Car.Model=model
	return c
}

func (c *CarBuilder)setBody(body string)(*CarBuilder){
	c.Car.Body=body
	return c
}
func(c *CarBuilder)setColor(color string)(*CarBuilder){
	c.Car.color=color
	return c
}

func(c *CarBuilder)setTyre(tyre1 string)(*CarBuilder){
	c.Car.tyre=tyre1
	return c
}

func(c *CarBuilder)build()Car{
	return c.Car
}


func main(){

	carBuilder:=&CarBuilder{}

	car1:=carBuilder.setModel("Range Rover Defender").setBody("Fiber Body").setColor("Olive Green").setTyre("JK Tyres").build()
	fmt.Println(car1)

}



