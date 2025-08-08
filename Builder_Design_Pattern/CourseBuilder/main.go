package main

import "fmt"

func main(){

java:=&JavaFullStack{}
fmt.Println("-------------Java FullStack Development--------------------------")
java1:=java.addAssignments("Java Oops Concepts").addVideos("Java"," basic Tutorials").addReadingMaterial("java_notes.pdf").addQuizzes("quizzes has been created for 10 weeks please ensure everyone completes on the assigned week").createCertificate("Certificate.pdf").build()
fmt.Println(java1)



python:=&PythonFullStack{}
fmt.Println("-------------Python FullStack Development--------------------------")
python1:=python.addAssignments("Python Concepts").addVideos("Python","Basic Tutorial").addReadingMaterial("python_notes.pdf").addQuizzes("QUiz for th python has been created successfully").createCertificate("Certificate.pdf").build()
fmt.Println(python1)

}