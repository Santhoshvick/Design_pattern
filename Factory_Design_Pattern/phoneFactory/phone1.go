package main
type Samsung struct{
	Phone
}


func newSamsung()FPhone{
	return &Samsung{
		Phone:Phone{
			Name: "Samsung",
			Price: 125999,
			Model: "Fold 5",
			Storage: "512 GB",
		},
	}
}