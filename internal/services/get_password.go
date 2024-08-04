package services

import "password_manager/internal/db"

func GetPassword(title string) (db.Password, bool) {
	for _, password := range db.Passwords{
		if password.Title == title {
			return password, true
		}
	}
	return db.Password{}, false
}