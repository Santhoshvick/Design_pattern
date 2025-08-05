package main

func GenerateNotification(serviceName string)NotificationService{
	if(serviceName=="Email"){
		return &Email{}
	}else if (serviceName=="Sms"){
		return &SMS{}
	}
	return nil
}