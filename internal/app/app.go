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
		fmt.Println("If you want to add a password, enter: 1")
		fmt.Println("display a password, enter: 2")
		fmt.Println("update a password, enter: 3")
		fmt.Println("delete a password, enter: 4")
		fmt.Println("exit the program, enter: 5")
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