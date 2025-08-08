package main

import "fmt"

type Production struct{}

func(p *Production)Build(){
	fmt.Println("************PRODUCTION*****************")
	fmt.Println("The Build is attached to the Production environment")
}

func(p *Production)Test(){
	fmt.Println("The Testing is Done in the Production")
}

func(p *Production)Deploy(){
	fmt.Println("The Deployment is done successfully")
}

func(p *Production)Notify(){
	fmt.Println("Notify the changes in the Production")
}