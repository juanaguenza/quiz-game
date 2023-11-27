package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
	"flag"
	// 	"math/rand"
	"strings"
)

var correctAnswers = 0
var wrongAnswers = 0

func main() {

	// Parsing flags
	timePtr := flag.Int("t", 10, "amt of time for the quiz")
	flag.Parse()

	file, err := os.Open("problems.csv")

	if err != nil {
		log.Fatal("Error opening .csv file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	reader.Comma = ','

	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Error reading records")
	}

	var questions []string
	var answers []string

	for _, record := range records {
		questions = append(questions, record[0])
		answers = append(answers, record[1])
	}

	var response string
	for response != "yes" {
		fmt.Print("Are you ready to play? (yes/no):")
		fmt.Scanln(&response)
		response = strings.ToLower(response)
	}

	beginQuiz(questions, answers, *timePtr)

	fmt.Println("Correct answers:", correctAnswers)
	fmt.Println("Wrong answers:", wrongAnswers)
	fmt.Println("Thanks for playing!")

}

// beginQuiz will start the quiz, quiz will end whenever timer runs out
func beginQuiz(problems []string, answers []string, duration int) {
	// Start the timer
	fmt.Println(duration)
	timer := time.NewTimer(time.Duration(duration) * time.Second)

	// Ask each question
	for i := range problems {
		// If the timer finishes, end the game
		// Otherwise continue asking questions
		select {
			case <- timer.C:
				fmt.Println("Time ran out...")
				return
			default:
				askQuestion(problems[i], answers[i])
		}
	}
}

// askQuestion will ask the user a question and receive their input
func askQuestion(question string, correctAnswer string) {
	// Ask the user the question
	fmt.Println(question + ":")

	// Get the user's answer
	var answer string
	fmt.Scanln(&answer)

	// Grade the user's input
	correct := gradeInput(answer, correctAnswer)
	keepScore(correct)
}

// gradeInput will grade the input of the user
// returns true if correct, otherwise false
func gradeInput(answer string, correctAnswer string) bool {
	return answer == correctAnswer
}

// keepScore will keep the score of the user
func keepScore(correct bool) {
	if correct {
		correctAnswers += 1
	} else {
		wrongAnswers += 1
	}
}
