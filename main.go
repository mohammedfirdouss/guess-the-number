package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var target int
var chances int
var attempts int

var indexTmpl *template.Template
var guessTmpl *template.Template

func main() {
	var err error
	indexTmpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		panic(fmt.Sprintf("Error parsing index template: %v", err))
	}

	guessTmpl, err = template.ParseFiles("templates/guess.html")
	if err != nil {
		panic(fmt.Sprintf("Error parsing guess template: %v", err))
	}

	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random number between 1 and 100
	target = randGen.Intn(100) + 1 
	
	// Register HTTP handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/start", startHandler)
	http.HandleFunc("/guess", guessHandler)

	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := indexTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	difficulty := r.FormValue("difficulty")
	switch difficulty {
	case "easy":
		chances = 10
	case "medium":
		chances = 5
	case "hard":
		chances = 3
	default:
		chances = 5
	}

	attempts = 0
	target = rand.Intn(100) + 1

	// Redirect to the guessing page.
	http.Redirect(w, r, "/guess", http.StatusSeeOther)
}


func guessHandler(w http.ResponseWriter, r *http.Request) {
	type Data struct {
		Message   string
		Remaining int
		GameOver  bool
		Target    int
		Success   bool
	}

	data := Data{
		Remaining: chances - attempts,
		GameOver:  false,
	}

	// Handle POST request (making a guess)
	if r.Method == http.MethodPost {
		guessStr := r.FormValue("guess")
		guess, err := strconv.Atoi(guessStr)
		if err != nil {
			data.Message = "Please enter a valid number."
			// Only execute template if valid input is provided
			if err := guessTmpl.Execute(w, data); err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
			}
			return
		}

		attempts++
		if guess == target {
			data.Message = fmt.Sprintf("Congratulations! You guessed the number in %d attempt(s).", attempts)
			data.GameOver = true
			data.Target = target
			data.Success = true
		} else if attempts >= chances {
			data.Message = "Sorry, you've run out of chances!"
			data.GameOver = true
			data.Target = target
		} else if guess < target {
			data.Message = fmt.Sprintf("Incorrect! The number is greater than %d.", guess)
		} else {
			data.Message = fmt.Sprintf("Incorrect! The number is less than %d.", guess)
		}

		data.Remaining = chances - attempts
	} else {
		// For GET request
		data.Message = "Make a guess!"
	}

	// Ensure template rendering is only done once and avoid sending headers twice
	if err := guessTmpl.Execute(w, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
