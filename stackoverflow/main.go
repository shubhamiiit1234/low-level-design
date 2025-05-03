package main

import (
	"fmt"
)

func main() {
	user := User{
		ID:    1,
		Name:  "Tyler",
		Email: "tyler@gmail.com",
	}

	question := Question{
		ID:      1,
		Title:   "Goroutine in Go",
		Content: "What is a goroutine?",
		Vote: Vote{
			Count: 0,
		},
		CreatedAt: "2025-05-03",
		Tags:      []Tag{{ID: 1, Name: "Go"}, {ID: 2, Name: "Concurrency"}},
	}

	user.AddQuestion(question)
	user.AddAnswer(&question, Answer{
		ID:        12,
		Content:   "A goroutine is a lightweight thread managed by the Go runtime.",
		Question:  &question,
		User:      user,
		CreatedAt: "2025-05-03",
		Vote: Vote{
			Count: 0,
		},
	})

	fmt.Println("user: ", user)
	fmt.Println("--------------------------------------------------------------------------------------------------------")
	fmt.Println("user's all questions: ", user.GetQuestions())
	fmt.Println("--------------------------------------------------------------------------------------------------------")
	fmt.Println("user's all answers: ", user.GetAnswers())
	fmt.Println("--------------------------------------------------------------------------------------------------------")
	fmt.Println("question: ", question.Title)
	fmt.Println("--------------------------------------------------------------------------------------------------------")
	fmt.Println("answer: ", question.Answers)
	fmt.Println("--------------------------------------------------------------------------------------------------------")
}
