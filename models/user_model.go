package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Balance  int
}

func SetUser(db *gorm.DB, username string, password string, balance int) {
	db.Create(&User{Username: username, Password: password, Balance: balance})
}

func GetUser(db *gorm.DB, id int) User {
	var user User
	db.First(&user, "id =  ?", id)

	return user
}

func GetUsers(db *gorm.DB) []User {
	var users []User
	db.Find(&users)

	return users
}

func UpdateUser(db *gorm.DB, id int, password string, balance int) {
	user := GetUser(db, id)
	db.Model(&user).Update("password", password)
	db.Model(&user).Update("balance", balance)
}

func DeleteUser(db *gorm.DB, id int) {
	user := GetUser(db, id)
	db.Delete(&user, "id = ?", id)
}
