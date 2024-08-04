package services

import (
	"fmt"
	"password_manager/internal/db"
	"password_manager/pkg/input"
)

func UpdatePassword(title string) {
	for i, password := range db.Passwords {
		if password.Title == title {
			fmt.Print("Enter the new password: ")
			value := input.GetInput()
			db.Passwords[i].Value = value
			fmt.Println("Password updated successfully!")
			return
		}
	}
	fmt.Println("Password not found.")
}
