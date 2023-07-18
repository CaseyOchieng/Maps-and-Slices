package main

import "log"

// How to make a map in Golang

func maps() {
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"

	myMap["Camels"] = "Casey"

	log.Println("My Dogs name is", myMap["dog"]+" and "+myMap["Dog"])

}
