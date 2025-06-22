package main

import "fmt"

/*
	Requirements:
		1. Create user profile
		2. Create group and Add user to group
		3. Add expense (amount, description, participants)
		4. Split the expenses among participants based on their share
		5. User should be able to view the balances with other users
		6. Settle up the balances
		7. Different split methods (Equal Split | Percentage Split | Exact Amount)
		8. View Users transaction history

	Core Entities:
		1. User
		2. Group
		3. Expense
		4. SplitMethod (enum) - EQUAL | PERC | EXACT
		5. Transaction

	Methods:
		1. User
			- CreateUser(user User)
			- SettleBalance(expenseId string)
			- DeleteUser(id string) (TBD)
			- UpdateUser(id string) (TBD)
		2. Group
			- CreateGroup(users []User)
			- AddUser(groupId string, user User)
			- DeleteGroup(groupId string) (TBD)
		3. Expense
			- AddExpense(paidBy string, groupId string, users []User, amount int)
			- DeleteExpense(groupId string, expenseId) (TBD)
			- UpdateExpense(groupId string, expenseId sting) (TBD)
		4. Transaction
			- CreateTransaction(senderId, receiverId string, amount int)

*/

func main() {
	fmt.Println("splitwise lld")

	splitwise := NewSplitwise()

	// Create User
	user1 := NewUser("1", "Alice", "alice@example.com")
	user2 := NewUser("2", "Bob", "bob@example.com")
	user3 := NewUser("3", "Charlie", "charlie@example.com")
	fmt.Println("user1 id: ", user1.UserID, "user2 id: ", user2.UserID, "user3 id: ", user3.UserID)
	splitwise.AddUser(user1)
	splitwise.AddUser(user2)
	splitwise.AddUser(user3)

	// Create Group
	group := NewGroup("Apartment", []*User{user1, user2, user3})
	fmt.Println("Group ID: ", group.GroupId)
	splitwise.AddGroup(group)

	// Create Expense
	splitwise.AddExpense(user1.UserID, group.GroupId, "Month May Rent", []*User{user1, user2, user3}, 12000, &EqualSplit{})

	// View Balance
	splitwise.ViewBalances(user1)
	splitwise.ViewBalances(user2)
	splitwise.ViewBalances(user3)

	// Settle Balance for a user
	splitwise.SettleBalance(user2, group.Expenses[0])
	splitwise.ViewBalances(user2)
	splitwise.ViewBalances(user3)
}
