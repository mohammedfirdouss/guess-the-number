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

	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random number between 1 and 100
	target := randGen.Intn(100) + 1
	fmt.Println("Debug: The target number is", target)

	// Welcome message and basic rules.
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("")

	// Present difficulty options.
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
	fmt.Print("Enter your choice: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	difficultyChoice, err := strconv.Atoi(input)
if err != nil || difficultyChoice < 1 || difficultyChoice > 3 {
	fmt.Println("Invalid choice. Defaulting to Medium difficulty (5 chances).")
	difficultyChoice = 2
}

var chances int
switch difficultyChoice {
case 1:
	chances = 10
	fmt.Println("Great! You have selected the Easy difficulty level.")
case 2:
	chances = 5
	fmt.Println("Great! You have selected the Medium difficulty level.")
case 3:
	chances = 3
	fmt.Println("Great! You have selected the Hard difficulty level.")
default:
	chances = 5
	fmt.Println("Great! You have selected the Medium difficulty level.")
}


fmt.Printf("You have %d chances to guess the correct number.\n", chances)
fmt.Println("Let's start the game!")

// Generate a secret number between 1 and 100.
secretNumber := rand.Intn(100) + 1

attempts := 0
for attempts < chances {
	fmt.Print("\nEnter your guess: ")
	guessInput, _ := reader.ReadString('\n')
	guessInput = strings.TrimSpace(guessInput)
	guess, err := strconv.Atoi(guessInput)
	if err != nil {
		fmt.Println("Please enter a valid number.")
		continue
	}

	attempts++

	if guess == secretNumber {
		fmt.Printf("Congratulations! You guessed the correct number in %d attempts\n", attempts)
		return
	} else if guess < secretNumber {
		fmt.Printf("Incorrect! The number is greater than %d.\n", guess)
	} else {
		fmt.Printf("Incorrect! The number is less than %d.\n", guess)
	}
}

fmt.Printf("\nYou've run out of chances. The correct number was %d.\n", secretNumber)
}
