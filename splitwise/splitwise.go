package main

import "fmt"

type Splitwise struct {
	Users  map[string]*User
	Groups map[string]*Group
}

func NewSplitwise() *Splitwise {
	return &Splitwise{
		Users:  map[string]*User{},
		Groups: map[string]*Group{},
	}
}

func (s *Splitwise) AddUser(user *User) {
	s.Users[user.ID] = user
}

func (s *Splitwise) AddGroup(group *Group) {
	s.Groups[group.GroupId] = group
}

func (s *Splitwise) CreateGroup(name string, participants []*User) *Group {
	group := NewGroup(name, participants)
	fmt.Println("Group: ", group)
	return group
}

func (s *Splitwise) AddExpense(paidBy, groupID, description string, participants []*User, amount float64, splitStrategy SplitStrategy) {
	newExpense := NewExpense(paidBy, groupID, description, participants, amount, splitStrategy)
	group := s.Groups[groupID]
	group.AddExpense(newExpense)
	fmt.Println("NewExpense: ", newExpense, " added to group: ", group.Name)
	splitStrategy.UpdateExpense(paidBy, participants, newExpense)
}

func (s *Splitwise) ViewBalances(user *User) {
	u := s.Users[user.ID]
	u.ViewBalances()
}

func (s *Splitwise) SettleBalance(user *User, expense *Expense) {
	u := s.Users[user.ID]
	u.SettleBalance(expense)
}
