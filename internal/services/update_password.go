package services

import (
	"fmt"
	"password_manager/pkg/input"
)

func UpdatePassword(title string) {
	password, found := GetPassword(title)
	if !found {
		fmt.Println("Password not found.")
		return
	}
	fmt.Print("Enter the new password: ")
	value := input.GetInput()

	password.Value = value
	fmt.Println("Password updated successfully!")
}