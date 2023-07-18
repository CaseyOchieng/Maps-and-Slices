package main

import "log"
//  How to make a slice in Golang
// var mySlice []string
// How to make a slice in Golang method 1
func Slice() {
	var mySlice []string
	mySlice = append(mySlice, "Apples")
	mySlice = append(mySlice, "Oranges")
	mySlice = append(mySlice, "Mangos")

	log.Println(mySlice)
}

// How to make a slice in Golang method 2
func Slice2() {
	fruits := []string{"Apples", "Oranges", "Mangos"}
	log.Println(fruits)
}

//Loop in Golang
//How to make a For loop in Golang
