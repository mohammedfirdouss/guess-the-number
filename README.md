# Guess-the-number

## Description
The **Number Guessing Game** is a simple project written in Go that comes in a basic web version.

In the webrsion, the computer randomly selects a number between 1 and 100, and the player must guess the number within a limited number of attempts. The allowed attempts depend on the chosen difficulty level:

- **Easy:** 10 chances  
- **Medium:** 5 chances  
- **Hard:** 3 chances  

As the game progresses, the player receives real-time hints on whether the secret number is higher or lower than their guess. The game concludes either when the player guesses correctly or runs out of attempts.

## Features

- **Interactive Experience:** A web interface.
- **Difficulty Selection:** Pick from Easy, Medium, or Hard to set the challenge level.
- **Real-Time Hints:** Get immediate feedback after every guess.
- **Randomized Number:** Each game session uses a freshly generated secret number for a new experience every time.

---

## Installation

Ensure you have [Go installed](https://golang.org/dl/) on your machine.

### Clone the Repository

```bash
git clone https://github.com/mohammedfirdouss/guess-the-number.git
cd guess-the-number
```
### Build the project:
```bash
go build -o guess-the-number main.go
```
Alternatively, you can run the game directly using:
```bash
go run main.go
```

### Usage
After installing, you can start the game by executing the compiled binary or running it via go run:

```bash
./guess-the-number
```
### Follow the on-screen instructions:

- Select a difficulty level (Easy, Medium, or Hard).
- Enter your guess when prompted.
- Receive feedback on whether the secret number is higher or lower than your guess.
- Win or lose based on your guesses and the remaining attempts.


Have Fun along the whole process!!
