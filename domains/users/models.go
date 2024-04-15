package users

import "time"

type User struct {
	ID        int
	FirstName string
	LastName  string
	BirthDay  time.Time
}
