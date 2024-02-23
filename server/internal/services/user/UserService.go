package UserService

import (
	UserModel "server/server/internal/models"
	Database "server/server/internal/services/database"
)

func AddUser(user *UserModel.User) {
	db := Database.GetInstance()

	//insert the user into the database
	res, err := db.Query("INSERT INTO users (email, username, password) VALUES ($1, $2, $3) RETURNING id;", user.Email, user.Username, user.Password)

	//check if there was an error and handle it
	if err != nil {
		panic(err)
	}

	if res.Next() {
		err = res.Scan(&user.Id)
		if err != nil {
			panic(err)
		}
	}
}

func VerifyUser(creds *UserModel.Credentials) (int, error) {
	var id int = -1
	db := Database.GetInstance()

	res, err := db.Query("SELECT id FROM users WHERE username = $1 and password = $2", creds.Username, creds.Password)
	if err != nil {
		return id, err
	}

	if res.Next() {
		err = res.Scan(&id)
		if err != nil {
			return -1, err
		}
	}

	return id, nil
}
