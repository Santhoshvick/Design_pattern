package main

type Kia struct{
    Car
}

func newKia()Icar{
	return &Kia{
		Car: Car{
		    name: "Kia Seltos",
			price: 2499000,
		},
	}
}