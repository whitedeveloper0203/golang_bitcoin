package update

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID            uint `gorm:"primary_key"`
	Auth0UUID     string
	Email         string
	EmailVerified bool
	DemoBalance   float64
	Bets          []byte `gorm:"type:text[]"`
	// StringBets    string
}

type Request struct {
	Auth0UUID     string  `json:"auth0uuid"`
	Email         string  `json:"email"`
	EmailVerified bool    `json:"email_verified"`
	DemoBalance   float64 `json:"demo_balance"`
	Bets          []uint8 `json:"bets"`
	StringBets    string  `json:"stringbets"`
}

type Response struct {
	ID uint
}

// DemoBalance queries/updates the db for a user with an auth0uuid that matches the one requested.
func DemoBalance(w http.ResponseWriter, r *http.Request) {
	userRequest := &Request{}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&userRequest)
	params := mux.Vars(r) // Gets params

	db, err := gorm.Open("postgres",
		`host=localhost
		user=postgres
		password=test
		dbname=userinfo_db
		sslmode=disable`)
	panicOnError(err)
	defer db.Close()

	fmt.Println("update, incoming userRequest:", userRequest.DemoBalance, "stringbets:", userRequest.StringBets)

	// Convert StringBets back to []byte for postgres
	convertedBetsForDB := []byte(userRequest.StringBets)

	user := &User{}
	db.Where("auth0_uuid = ?", params["auth0uuid"]).First(&user)

	db.Find(&user)
	user.DemoBalance = userRequest.DemoBalance
	user.Bets = convertedBetsForDB
	db.Save(&user)

	// DB is done -
	fmt.Println("(updated) userbalance: ", user.DemoBalance, "bets: ", user.Bets)

	json.NewEncoder(w).Encode(&userRequest.DemoBalance)
}

func panicOnError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
