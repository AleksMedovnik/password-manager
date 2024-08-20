/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"password_manager/internal/services"
	"password_manager/pkg/input"
)

// deletePasswordCmd represents the deletePassword command
var deletePasswordCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a Password",
	Long: `Команда удаляет выбранный пароль`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Enter the title of the password to delete: ")
		title := input.GetInput()
		services.DeletePassword(title)
	},
}

func init() {
	rootCmd.AddCommand(deletePasswordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deletePasswordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deletePasswordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
