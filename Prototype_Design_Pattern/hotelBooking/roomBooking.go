package main

import "fmt"

type Room struct{
	RoomNumber string
	RoomType string
	Food string
}

type Customer struct{
	CustomerName string
	CustomerIDProof string
	CustomerEmail string
	CustomerMobileNo int64
	r Room
}

type Clonable interface{
	clone()Clonable
}

func(r *Room)clone()Clonable{
	return &Room{
		RoomNumber: r.RoomNumber,
		RoomType: r.RoomNumber,
		Food: r.Food,
	}
}

func main(){
	room1:=&Room{RoomNumber:"101A",RoomType: "Suite",Food: "Include"}

	customer:=&Customer{CustomerName: "Xyz",CustomerIDProof: "Aadhar Card",CustomerEmail: "abcd@gmail.com",CustomerMobileNo: 9876543210,r: *room1}
	fmt.Println(customer)

	room2:=room1.clone().(*Room)
	room2.RoomNumber="102B"
	room2.RoomType="Luxury Suite"

	customer2:=&Customer{CustomerName: "mnop",CustomerIDProof: "Voter Id",CustomerEmail: "wmnonp@gmail.com",CustomerMobileNo: 9890786712,r:*room2}
	fmt.Println(customer2)

}

