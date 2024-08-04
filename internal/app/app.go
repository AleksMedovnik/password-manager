package app

import (
	"fmt"
	"os"

	"password_manager/internal/db"
	"password_manager/internal/services"
	"password_manager/pkg/input"
)


func App() {
	for {
		fmt.Println("[1] Add Password")
		fmt.Println("[2] Display Passwords")
		fmt.Println("[3] Update Password")
		fmt.Println("[4] Delete Password")
		fmt.Println("[5] Exit")
		fmt.Print("Enter your choise: ")

		choice := input.GetInput()
		switch choice {
		case "1":
			db.Passwords = services.AddPassword(db.Passwords)
		case "2":
			services.DisplayPasswords(db.Passwords)
		case "3":
			fmt.Print("Enter the title of the password to update: ")
			title := input.GetInput()
			services.UpdatePassword(title)
		case "4":
			fmt.Print("Enter the title of the password to delete: ")
			title := input.GetInput()
			services.DeletePassword(title)
		case "5":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choise.")
		}
	}
}