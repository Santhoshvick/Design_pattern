package main

type Iotp interface{
	genOtp(int)string
	saveOtp(string)
	getMessage(string)string
	sendNotification(string)error
}


type otp struct{
	iotp Iotp
}


func (o *otp) generSendOtp(otp int)error{
	otp1:=o.iotp.genOtp(otp)
	o.iotp.saveOtp(otp1)
	Message1:=o.iotp.getMessage(otp1)

	err:=o.iotp.sendNotification(Message1)
	if err!=nil{
		return err
	}
	return nil

}
