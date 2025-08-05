package main

type Apple struct{
	Phone
}


func newApple()FPhone{
	return &Apple{
		Phone: Phone{
			Name: "Apple",
			Price: 159999,
			Model: "16 Pro Max",
			Storage: "1 TB",
		},
	}
}