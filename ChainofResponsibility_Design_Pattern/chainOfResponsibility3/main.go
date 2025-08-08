package main

func main(){
	cart:=&Cart{}
	order:=&Order{}
	payment:=&Payment{}
	shipment:=&Shipment{}
	delivery:=&Delivery{}

	
	
	cart.setNext(order)
	order.setNext(payment)
	payment.setNext(shipment)
	shipment.setNext(delivery)
	delivery.setNext(nil)

	
	customer:=&Customer{CustomerEmailId: "abcdeef@gmail.com",CustomerName: "xyz",CustomerMobileNo: 9876543210,CustomerAddress: " No:20 Anna Street velacherry Chennai-600003"}
	cart.execute(customer)


	

}