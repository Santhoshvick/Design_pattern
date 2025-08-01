package main

type Kia struct{
	Car 
}

func newKia()IcarBuilder{
	return &Kia{
		Car: Car{
			engineName: "BS6 Engine",
			tyre1: "Mitchell Typres",
			color1: "Olive Black",
			model: "Seltos",
		},
	}
}