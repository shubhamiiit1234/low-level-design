package main

type AddDataForQuestion interface {
	AddAnswer(a Answer)
	GetAnswers(q Question) []Answer
	AddTags(tags []Tag)
	GetTags() []Tag
}

type Question struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      []Tag     `json:"tags"`
	Comments  []Comment `json:"comments"`
	Answers   []Answer  `json:"answers"`
	Vote      Vote      `json:"vote"`
	CreatedAt string    `json:"created_at"`
	User      User      `json:"user"`
}

func (q *Question) AddComment(c Comment) {
	q.Comments = append(q.Comments, c)
}

func (a *Question) GetComments() []Comment {
	return a.Comments
}

func (q *Question) AddAnswer(a Answer) {
	q.Answers = append(q.Answers, a)
}

func (q *Question) GetAnswers() []Answer {
	return q.Answers
}

func (q *Question) Upvoate() {
	q.Vote.Count++
}

func (q *Question) Downvote() {
	q.Vote.Count--
}

func (q *Question) AddTags(tags []Tag) {
	q.Tags = append(q.Tags, tags...)
}

func (q *Question) GetTags() []Tag {
	return q.Tags
}
