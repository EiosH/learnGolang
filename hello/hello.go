package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	sentences, err := greetings.Hellos([]string{"Gladys", "Samantha", "Darrin"})
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// the Hello function to get a message for each name.
	for key, sentence := range sentences {
		fmt.Println(key + ": " + sentence)
	}

}
