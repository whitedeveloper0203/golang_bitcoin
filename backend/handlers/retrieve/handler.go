package retrieve

import (
	"encoding/json"
	"log"
	"net/http"

	"bitbucket.com/cobinix/backend/user"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// GetUser queries the db for a user with an auth0uuid that matches the one requested.
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params

	db, err := gorm.Open("postgres",
		`host=localhost
		user=postgres
		password=test
		dbname=userinfo_db
		sslmode=disable`)
	panicOnError(err)
	defer db.Close()

	var user user.User

	db.Where("auth0_uuid = ?", params["auth0uuid"]).First(&user)

	// Cast the []byte that postgres returns into a string
	user.StringBets = string(user.Bets)

	json.NewEncoder(w).Encode(&user)
}

func panicOnError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
