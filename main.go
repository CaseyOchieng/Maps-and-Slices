package main

import "log"

// How to make a map in Golang

func main() {
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"

	myMap["Dog"] = "Casey"

	log.Println("My Dogs name is", myMap["dog"]+" and "+myMap["Dog"])

}

// Maps are collections of key-value pairs

// How to make a slice in Golang
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
