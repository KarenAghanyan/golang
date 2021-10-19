package models

import (
	"fmt"
	"net/http"
)

type User struct {
	Firstname        string `json:"firstname"`
	Lastname         string `json:"lastname"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Password_Confirm string `json:"password_confirm"`
}

func (u *User) UserRegister(r *http.Request) {
	if len(u.Password) < 6 || len(u.Password_Confirm) < 6 {
		fmt.Println("The Password must be at least 6 long")
		return
	}

	if u.Password != u.Password_Confirm {
		fmt.Println("Passwords don't match")
		return
	}

	valid := false
	for i := 0; i < len(u.Email); i++ {
		if u.Email[i] == '@' {
			valid = true
		}
	}
	if !valid {
		fmt.Println("Email isn't valid")
		return
	}

	insertDynStmt := `insert into "Students"("Name", "Roll") values($1, $2)`
    _, e = db.Exec(insertDynStmt, "Jane", 2)
    CheckError(e)
}

}
