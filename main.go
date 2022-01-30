package main

import (
	"Test5/keyboard"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Enter a grade")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "falling"
	}
	fmt.Println("A grade of", grade, "is", status)
}
