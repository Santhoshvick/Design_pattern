package main

type Car struct{
	color1 string
	tyre1 string
	engineName string
	model string
}

func (c *Car)setColor(color1 string){
	c.color1=color1
}

func (c *Car)setTyreSize(tyre1 string){
	c.tyre1=tyre1
}

func (c *Car) setEngine(engineName string){
	c.engineName=engineName
}

func (c *Car)setModel(model string){
	c.model=model
}

func(c *Car)getCar()Car{
	return Car{
		tyre1: c.tyre1,
		color1: c.color1,
		model: c.model,
		engineName: c.engineName,
	}
}