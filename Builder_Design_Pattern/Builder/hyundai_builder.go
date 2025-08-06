package main

type Hyundai struct{
	Car
}

func newHyundai()IcarBuilder{
	return &Hyundai{
		Car: Car{
			color1: "Red",
			engineName: "BS4 Engine",
			tyre1: "Jacob Tyres",
			model: "Creta",
		},
	}
}