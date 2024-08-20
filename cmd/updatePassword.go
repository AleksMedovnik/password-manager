/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"password_manager/pkg/input"
	"password_manager/internal/services"
	"github.com/spf13/cobra"
)

// updatePasswordCmd represents the updatePassword command
var updatePasswordCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a Password",
	Long: `Команда обновляет пароли`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Enter the title of the password to update: ")
		title := input.GetInput()
		services.UpdatePassword(title)
	},
}

func init() {
	rootCmd.AddCommand(updatePasswordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updatePasswordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updatePasswordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
