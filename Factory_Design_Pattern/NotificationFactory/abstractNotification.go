package main

type NotificationService interface{
	sendNotification(receiver string,message string)
	// cacheNotification(receiver string,name string)
}
