package main

func main(){

    reception1:=&Reception{}
	doctor1:=&Doctor{}
	medication1:=&Medication{}
	cashier1:=&Cashier{}

	patient1:=&Patient{name: "Santhosh"}
	reception1.execute(patient1)

	reception1.setNext(doctor1)

	
    doctor1.setNext(medication1)

	medication1.setNext(cashier1)


	
	

	

	

	
	


	

	
}