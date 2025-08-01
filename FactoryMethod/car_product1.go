package main

type Hyundai struct{
	Car
}

func newHyundai()Icar{
	return &Hyundai{
		Car:Car{
			name: "Hyundai",
			price: 200000,
		},
	}
}