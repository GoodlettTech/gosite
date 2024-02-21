package UserService

import (
	"fmt"
	UserModel "server/server/internal/models"
	Database "server/server/internal/services"
)

func AddUser(user *UserModel.User) {
	db := Database.GetInstance()
	res, err := db.Exec("INSERT INTO users (email, username, password) VALUES (?, ?, ?);", user.Email, user.Username, user.Password)

	if err != nil {
		fmt.Println(res, err)
	}
}
