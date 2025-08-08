package main

type Deployment interface{
	Build()
	Test()
	Deploy()
	Notify()
}

type buildTest struct{
	deployment Deployment
}


func(b *buildTest)startDeployment(){
	b.deployment.Build()
	b.deployment.Test()
	b.deployment.Deploy()
	b.deployment.Notify()
}