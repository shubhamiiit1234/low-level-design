package main

import (
	"fmt"
	"strings"
)

type AddDataForUser interface {
	AddQuestion(q Question)
	AddAnswer(a Answer)
	AddComment(c Comment)
	GetQuestions() []Question
	GetAnswers() []Answer
	GetComments() []Comment
	Upvoate()
	Downvote()
	GetAnswersByQuestion(q Question) []Answer
}

type User struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Questions []Question `json:"questions"`
	Answers   []Answer   `json:"answers"`
	Comments  []Comment  `json:"comments"`
}

func (u *User) AddQuestion(q Question) {
	u.Questions = append(u.Questions, q)
}

func (u *User) AddAnswer(q *Question, a Answer) {
	q.AddAnswer(a)
	u.Answers = append(u.Answers, a)
}

func (u *User) AddComment(c Comment, respondable Respondable) {
	respondable.AddComment(c)
	u.Comments = append(u.Comments, c)
}

func (u *User) GetQuestions() []Question {
	return u.Questions
}

func (u *User) GetAnswers() []Answer {
	return u.Answers
}

func (u *User) GetComments() []Comment {
	return u.Comments
}

func (u *User) GetAnswersByQuestion(q Question) []Answer {
	var answers []Answer
	for _, a := range u.Answers {
		if a.Question.ID == q.ID {
			answers = append(answers, a)
		}
	}
	return answers
}

func (u *User) Upvoate(respondable Respondable) {
	respondable.Upvoate()
}

func (u *User) Downvote(respondable Respondable) {
	respondable.Downvote()
}

func (u User) String() string {
	var qTitles, qContent, aSummaries, cContents []string

	for _, q := range u.Questions {
		qTitles = append(qTitles, q.Title)
		qContent = append(qContent, q.Content)

	}
	for _, a := range u.Answers {
		aSummaries = append(aSummaries, fmt.Sprintf("AnswerID: %d", a.ID)) // customize as needed
	}
	for _, c := range u.Comments {
		cContents = append(cContents, c.Content)
	}

	return fmt.Sprintf("User{ID: %d, Name: %s, Email: %s, Questions: [%s %s], Answers: [%s], Comments: [%s]}",
		u.ID,
		u.Name,
		u.Email,
		strings.Join(qTitles, "; "),
		strings.Join(qContent, "; "),
		strings.Join(aSummaries, "; "),
		strings.Join(cContents, "; "),
	)
}
