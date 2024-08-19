/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"password_manager/internal/db"
	"password_manager/internal/services"
)

// displayPasswordCmd represents the displayPassword command
var displayPasswordCmd = &cobra.Command{
	Use:   "display",
	Short: "Показать пароли",
	Long: `Крманда выводит в терминал все имеющиеся пароли`,
	Run: func(cmd *cobra.Command, args []string) {
		services.DisplayPasswords(db.Passwords)
	},
}

func init() {
	rootCmd.AddCommand(displayPasswordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// displayPasswordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// displayPasswordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
