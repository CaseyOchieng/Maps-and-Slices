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

