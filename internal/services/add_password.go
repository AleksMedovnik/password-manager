package services

import (
	"fmt"
	"password_manager/internal/db"
	"password_manager/pkg/input"
)

func AddPassword(psw []db.Password) []db.Password {
	fmt.Print("Enter the title: ")
	title := input.GetInput()
	fmt.Print("Enter the password: ")
	value := input.GetInput()

	password := db.Password{
		Title: title,
		Value: value,
	}

	psw = append(psw, password)
	fmt.Println("Password added successfully!")
	return psw
}