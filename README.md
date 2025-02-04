# Guess-the-number

### Description
The Number Guessing Game is a simple command-line interface (CLI) game written in Go. In this game, the computer randomly selects a number between 1 and 100, and the player must guess the number within a limited number of attempts. 

The number of allowed attempts depends on the chosen difficulty level:

- Easy: 10 chances
- Medium: 5 chances
- Hard: 3 chances

As the game progresses, the user receives hints whether the secret number is higher or lower than their guess. The game concludes either when the user guesses the correct number or runs out of chances.

### Features

- Interactive CLI: Engage with a simple, text-based interface.
- Difficulty Selection: The user can choose between Easy, Medium, or Hard modes to set the challenge level.
- Real-Time Hints: Also, the user can receive immediate feedback after each guess.
- Randomized Number: The secret number is generated randomly each time the game is played, ensuring a new experience on every run.

## Installation
To run this project, ensure you have Go installed on your machine.

1. ## Clone the repository:
```bash
git clone https://github.com/mohammedfirdouss/guess-the-number.git
cd guess-the-number
```
2. ## Build the project:
```bash
go build -o guess-the-number main.go
```
Alternatively, you can run the game directly using:
```bash
go run main.go
```

3. ## Usage
After installing, you can start the game by executing the compiled binary or running it via go run:

```bash
./guess-the-number
```
4. ## Follow the on-screen instructions:

- Select a difficulty level (Easy, Medium, or Hard).
- Enter your guess when prompted.
- Receive feedback on whether the secret number is higher or lower than your guess.
- Win or lose based on your guesses and the remaining attempts.


Have Fun along the whole process!!
