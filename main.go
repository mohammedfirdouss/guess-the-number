package main

import (
	"bufio"     
	"fmt"       
	"math/rand" 
	"os"        
	"strconv"   
	"strings"   
	"time"      
)

func main() {

	// Create a new random generator
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random number between 1 and 100
	target := randGen.Intn(100) + 1
	fmt.Println("Debug: The target number is", target)
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("")

	// Present difficulty options.
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
	fmt.Print("Enter your choice: ")

}
