package services

import (
	"fmt"
	"password_manager/internal/db"
)

func DeletePassword(title string) {
	for i, password := range db.Passwords {
		if password.Title == title {
			db.Passwords = append(db.Passwords[:i], db.Passwords[i+1:]...)
			fmt.Println("Password deleted successfully!")
            return
		}
	}
	fmt.Println("Password not found.")
}