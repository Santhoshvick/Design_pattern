package main

import "fmt"

type Sms struct{
	otp
}

func(s *Sms)genOtp(otp int)string{
	randomOtp:="1234"
	fmt.Printf("The random otp is %d",randomOtp)
	return  randomOtp
}

func(s *Sms)saveOtp(otp1 string){
	fmt.Printf("The Generated OTP is Saved Successfully %s",otp1)
}

func(s *Sms)getMessage(otp string)string{
	fmt.Println("The Message For the particular otp")
	return "The Message is send to the user throught the sms"+otp
}

func(s *Sms)sendNotification(otp string)error{
	fmt.Printf("The Notification is ready to send %s",otp)
	return nil
}