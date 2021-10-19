package db

type Users struct {
	ID              string `json:"id"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Username        string `json:"name"`
	Email           string `json:"email"`
	Password_Hushed []byte `json:"password_hushed"`
}

var Users_Table []Users
