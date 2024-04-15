package users

import (
	"log"
	"time"
)

var userID int
var datesLayout = "2006-01-02"

func generateUserID() int {
	userID = userID + 1
	return userID
}

func formatDate(sDate string) time.Time {
	date, err := time.Parse(datesLayout, sDate)
	if err != nil {
		log.Printf("Error parsing date: %s", sDate)
	}

	return date
}

var inMemoryDB = []User{
	{ID: generateUserID(), FirstName: "Alice", LastName: "Smith", BirthDay: formatDate("2001-01-01")},
	{ID: generateUserID(), FirstName: "Charlie", LastName: "Garcia", BirthDay: formatDate("1992-11-23")},
	{ID: generateUserID(), FirstName: "Edward", LastName: "Brown", BirthDay: formatDate("1978-09-18")},
	{ID: generateUserID(), FirstName: "Diana", LastName: "Williams", BirthDay: formatDate("2005-03-10")},
	{ID: generateUserID(), FirstName: "Bob", LastName: "Johnson", BirthDay: formatDate("1985-07-12")},
}

type UsersRepository struct {
	db *[]User
}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{&inMemoryDB}
}

func (r UsersRepository) GetByID(id int) User {
	for _, u := range *r.db {
		if u.ID == id {
			return u
		}
	}
	return User{}
}

func (r UsersRepository) CreateUser(u User) {
	u.ID = generateUserID()
	*r.db = append(*r.db, u)
}

func (r UsersRepository) UpdateUser(u User) {
	for idx, dbUser := range *r.db {
		if dbUser.ID == u.ID {
			(*r.db)[idx] = u
		}
	}
}
