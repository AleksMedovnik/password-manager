package services

import (
	"fmt"
	"strconv"

	"password_manager/internal/db"
	"password_manager/pkg/input"
)

func DisplayPasswords(psw []db.Password) {
	if len(psw) == 0 {
		fmt.Println("No passwords found.")
		return
	}
	fmt.Println("Passwords:")
	for i, password := range psw {
		fmt.Printf("%d. %s\n", i+1, password.Title)
	}

	fmt.Print("Enter the number of the password to display: ")
	choice := input.GetInput()
	index, _ := strconv.Atoi(choice)
	if index <= 0 || index > len(psw) {
		fmt.Println("Invalid choice.")
		return
	}

	password := psw[index-1]
	fmt.Println("Password: ", password.Value)
}

