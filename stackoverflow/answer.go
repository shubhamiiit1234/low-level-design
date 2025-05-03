package main

type AddDataForAnswer interface {
	MarkAccepted()
}

type Answer struct {
	ID         int       `json:"id"`
	User       User      `json:"user"`
	Question   *Question `json:"question"`
	Content    string    `json:"content"`
	Comments   []Comment `json:"comments"`
	Vote       Vote      `json:"vote"`
	CreatedAt  string    `json:"created_at"`
	IsAccepted bool      `json:"is_accepted"`
}

func (a *Answer) AddComment(c Comment) {
	a.Comments = append(a.Comments, c)
}

func (a *Answer) GetComments() []Comment {
	return a.Comments
}

func (a *Answer) Upvote() {
	a.Vote.Count++
}

func (a *Answer) Downvote() {
	a.Vote.Count--
}

func (a *Answer) MarkAccepted() {
	a.IsAccepted = true
}
