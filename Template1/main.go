package main

func main(){
	smsOTP := &Sms{}
    o := otp{
        iotp: smsOTP,
    }
    o.generSendOtp(4)
}