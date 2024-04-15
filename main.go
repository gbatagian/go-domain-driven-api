package main

import (
	"fmt"
	"go-domain-driven-api/domains/users"
	"time"
)

//	func main() {
//		app.RunWithSettings(settings.EnvSettings)
//	}
func main() {
	repo := users.NewUsersRepository()

	fmt.Println(repo.GetByID(1))

	var u users.User

	bDay, _ := time.Parse("2006-01-02", "1985-07-12")
	u = users.User{FirstName: "George", LastName: "Batas", BirthDay: bDay}

	repo.CreateUser(u)
	fmt.Println(repo.GetByID(6))

	u = users.User{ID: 6, FirstName: "George", LastName: "Batagiannis", BirthDay: bDay}

	repo.UpdateUser(u)
	fmt.Println(repo.GetByID(16))
}
