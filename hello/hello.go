package main

import (
	"fmt"
	"log"
	"regexp"

	"chef.io/greetings"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("bob")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

	// for a sample test run, we'll see if a string is actually an int
	s := "hello world"
	ret := matchInt(s)
	fmt.Println(ret)
	
	s = "10"
	ret = matchInt(s)
	fmt.Println(ret)

	// return or os.Exit(0)
}
