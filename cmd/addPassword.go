/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"password_manager/internal/db"
	"password_manager/internal/services"
)

// addPasswordCmd represents the addPassword command
var addPasswordCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a Password",
	Long: `Команда добавляет пароль в хранилище паролей`,
	Run: func(cmd *cobra.Command, args []string) {
		db.Passwords = services.AddPassword(db.Passwords)
	},
}

func init() {
	rootCmd.AddCommand(addPasswordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addPasswordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addPasswordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
