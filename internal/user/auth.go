package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/matthewzhaocc/planetary-deploy/internal/types"
	"gorm.io/gorm"
)

func authenticate(serv types.Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user NewUser
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var userQuery = User{Username: user.Username}
		var dbUser User
		var result = serv.DB.Where(&userQuery).First(&dbUser)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !dbUser.verifyHash(user.Password) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// TODO Implement JWT session
	}
}
