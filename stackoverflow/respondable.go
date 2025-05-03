package main

type Respondable interface {
	AddComment(c Comment)
	GetComments() []Comment
	Upvoate()
	Downvote()
}
