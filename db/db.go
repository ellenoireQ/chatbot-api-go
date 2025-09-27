package db

type Message struct {
	ID        int
	User      User
	Assistant Assistant
}

type User struct {
	Role    string
	Content string
}

type Assistant struct {
	Role    string
	Content string
}