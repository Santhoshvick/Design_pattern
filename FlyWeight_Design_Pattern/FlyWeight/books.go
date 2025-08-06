package main

import "fmt"

type Book struct{
	Title string
	Author string
}

func(b *Book)display(location string){
	fmt.Printf("The Title of the book %s and the Author %s and location %s",b.Title,b.Author,location)
}