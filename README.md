# Here is a code that shows you how to create Slices and Maps in Golang


## How to make a map in Golang
package main

import "log"

func main() {
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"

	myMap["Dog"] = "Casey"

	log.Println("My Dogs name is", myMap["dog"]+" and "+myMap["Dog"])

}


## How to make a slice in Golang
// var mySlice []string
##### How to make a slice in Golang method 1
func Slice() {
	var mySlice []string
	mySlice = append(mySlice, "Apples")
	mySlice = append(mySlice, "Oranges")
	mySlice = append(mySlice, "Mangos")

	log.Println(mySlice)
}

#####How to make a slice in Golang method 2
func Slice2() {
	fruits := []string{"Apples", "Oranges", "Mangos"}
	log.Println(fruits)
}
