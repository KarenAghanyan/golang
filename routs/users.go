package routs

import (
	"crypto/sha1"
	"encoding/json"
	"math/rand"
	"net/http"
	"simpleAPI/models"
	"strconv"
)

type Users struct {
	ID              string `json:"id"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Username        string `json:"name"`
	Email           string `json:"email"`
	Password_Hushed []byte `json:"password_hushed"`
}

var Users_Table []Users

func UserRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := models.User{}
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.UserRegister(r)
	// Users_Table = append(Users_Table, authenUser(&user))
	json.NewEncoder(w).Encode(Users_Table)
}

func authenUser(u *models.User) Users {
	user_db := Users{}
	user_db.ID = strconv.Itoa(rand.Intn(100000000))
	user_db.Firstname = u.Firstname
	user_db.Lastname = u.Lastname
	user_db.Username = u.Username
	user_db.Email = u.Email
	hash := sha1.New()
	hash.Write([]byte(u.Password))
	pass_hashed := hash.Sum(nil)
	user_db.Password_Hushed = pass_hashed

	// sqlStatement := `insert into "Users"("id", "firstname", "lastname",
	// 									  "username", "email", "password_hushed")
	// 									   values($1, $2, $3, $4, $5, $6)
	// 									   RETURNING id`
	// id := 0
	// err := db.QueryRow(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun").Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
}
