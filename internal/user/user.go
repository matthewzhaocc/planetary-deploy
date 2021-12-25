package user

import (
	"encoding/json"
	"net/http"

	"github.com/matthewzhaocc/planetary-deploy/internal/types"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *NewUser) toHash() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(bytes)
}

func (u *User) verifyHash(pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
	return err == nil
}

func AddUser(serv types.Server, u NewUser) {
	user := User{Username: u.Username, Password: u.Password}
	serv.DB.Create(&user)
}

func AddUserRoute(serv types.Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var user NewUser
		json.NewDecoder(r.Body).Decode(&user)
		AddUser(serv, user)
	}
}
