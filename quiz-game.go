package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type Question struct {
	Text   string `json:"text"`
	Answer string `json:"answer"`
}

type QuestionData struct {
	Question []Question `json:"question_data"`
}

func CheckAnswer(ans, input string) bool {
	if ans == input {
		return true
	} else {
		return false
	}
}

func main() {
	var questions QuestionData
	var userChoice string
	var score int
	counter := 1

	byteValue, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(byteValue, &questions)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println("Welcome to the quiz game!")
loop:
	for {
		currentQuestion := questions.Question[rand.Intn(len(questions.Question))-1]
		fmt.Printf("Q%v. %v (True or False): ", counter, currentQuestion.Text)
		fmt.Scanln(&userChoice)
		if userChoice == "off" {
			break loop
		} else {
			answer := CheckAnswer(userChoice, currentQuestion.Answer)
			switch answer {
			case true:
				fmt.Println("Correct!")
				score += 1
				fmt.Printf("Your curerent score is %v/%v\n", score, counter)
				counter += 1
			case false:
				fmt.Println("You got it wrong")
				fmt.Printf("Your curerent score is %v/%v\n", score, counter)
				counter += 1
			}
		}
	}
}
