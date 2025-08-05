package main

type Emirates struct{
	flight
}

func Emirates1()Iflight{
	return &Emirates{
		flight: flight{
			name: "Emirates",
			price: 125000,
			class: "first class",
		},
	}
}

