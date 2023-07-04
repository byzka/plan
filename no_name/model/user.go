package model

type User struct {
	ID           uint
	Name         string
	Passwd       string
	Email        string
	Birthday     string
	Phone_number uint16
	State        int
}

// ID
// Name
// Passwd
// Email
// Birthday
// Phone_number
// State
const (
	Online    = 1
	Notonline = 0
)
