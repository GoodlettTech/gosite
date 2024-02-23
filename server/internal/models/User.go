package UserModel

type User struct {
	Id       int
	Email    string
	Username string
	Password string
}

type Credentials struct {
	Username string
	Password string
}
