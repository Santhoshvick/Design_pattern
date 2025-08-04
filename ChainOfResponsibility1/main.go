package main




func main(){
	testDrive1:=&TestDrive{}
	carBooking1:=&CarBooking{}
	payment1:=&Payments{}
	delivery1:=&Delivery{}

	customer1:=&Customer{name: "Santhosh",CarName: "BMW"}

	testDrive1.execute(customer1)
	testDrive1.setNext(carBooking1)
	carBooking1.setNext(payment1)
	payment1.setNext(delivery1)
}