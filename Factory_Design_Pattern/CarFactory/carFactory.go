package main

import "fmt"

func getCar(name string)(Icar,error){
	if name=="Kia"{
		return newKia(),nil
	}
	if name=="Hyundai"{
		return newHyundai(),nil
	}
	return nil,fmt.Errorf("Invalid Car Type")
}