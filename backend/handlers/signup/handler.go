package signup

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"bitbucket.com/cobinix/backend/user"
	"github.com/jinzhu/gorm"
)

// Request takes the json from POST and its values are passed to the Signup Request struct arg.
type Request struct {
	Auth0UUID     string  `json:"auth0uuid"`
	Email         string  `json:"email"`
	EmailVerified bool    `json:"email_verified"`
	DemoBalance   float64 `json:"demo_balance"`
}

type Response struct {
	ID uint
}

var userRequest Request

// CreateUser Gets called when user posts to the /api/userinfo endpoint
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&userRequest)

	fmt.Println("auth0 uuid: ", userRequest.Auth0UUID)
	fmt.Println("email: ", userRequest.Email)
	fmt.Println("email verified: ", userRequest.EmailVerified)
	fmt.Println("demo balance: ", userRequest.DemoBalance)

	json.NewEncoder(w).Encode(userRequest)

	db, err := gorm.Open("postgres",
		`host=localhost
		user=postgres
		password=test
		dbname=userinfo_db
		sslmode=disable`)
	panicOnError(err)
	defer db.Close()

	signupUser(db)
}

func signupUser(db *gorm.DB) {
	res, err := signup(db, &Request{
		Auth0UUID:     userRequest.Auth0UUID,
		Email:         userRequest.Email,
		EmailVerified: userRequest.EmailVerified,
		DemoBalance:   userRequest.DemoBalance,
	})
	if err != nil {
		switch err.(type) {
		case *user.Auth0UUIDDuplicateError:
			fmt.Println("Bad Request: ", err.Error())
			return
		case *user.EmailDuplicateError:
			fmt.Println("Bad Request: ", err.Error())
			return
		default:
			fmt.Println("Internal Server Error: ", err.Error())
			return
		}
	}
	fmt.Println("Created User: ", res.ID)
}

func signup(db *gorm.DB, req *Request) (*Response, error) {
	newUser := &user.User{
		Auth0UUID:     req.Auth0UUID,
		Email:         req.Email,
		EmailVerified: req.EmailVerified,
		DemoBalance:   req.DemoBalance,
	}

	id, err := user.Create(db, newUser)
	if err != nil {
		return nil, err
	}
	return &Response{ID: id}, err
}

func panicOnError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
