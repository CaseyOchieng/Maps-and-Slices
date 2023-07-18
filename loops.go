package main

import "log"

func loops() {
	//Here is a Slice of animals
	animals := []string{"Cat", "Dog", "Fish", "Horse", "Camels", "cows"}

	// We are go to create a loop that accept two arguments i is the integer and animals .
	for i, animal := range animals {
		log.Println(i, animal)
	}

}

func loopsMaps() {
	//Here is a Maps of animals
	animals := make(map[string]string)
	animals["Cat"] = "Meow"
	animals["Dog"] = "Ben"
	animals["Fish"] = "Erica"
	animals["Horse"] = "Hosea"
	animals["Camels"] = "Casey"
	// We are go to create a loop that accept two arguments animalType and animals.
	for animalType, animal := range animals {
		log.Println(animalType, animal)
	}

}
