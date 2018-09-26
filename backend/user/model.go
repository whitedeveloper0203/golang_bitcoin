package user

import (
	"fmt"
)

const (
	UniqueConstraintAuth0UUID = "users_auth0UUID_key"
	UniqueConstraintEmail     = "users_email_key"
)

type User struct {
	ID            uint `gorm:"primary_key"`
	Auth0UUID     string
	Email         string
	EmailVerified bool
	DemoBalance   float64
	Bets          []byte `gorm:"type:text[]"`
	StringBets    string
}

type Auth0UUIDDuplicateError struct {
	Auth0UUID string
}

func (e *Auth0UUIDDuplicateError) Error() string {
	return fmt.Sprintf("Auth0UUID '%s' already exists", e.Auth0UUID)
}

type EmailDuplicateError struct {
	Email string
}

func (e *EmailDuplicateError) Error() string {
	return fmt.Sprintf("Email '%s' already exists", e.Email)
}

/*
userinfo_db=# \d users
                                    Table "public.users"
    Column      |          Type          |                     Modifiers
----------------+------------------------+----------------------------------------------------
 id             | integer                | not null default nextval('users_id_seq'::regclass)
 auth0_uuid     | character varying(50)  | not null
 email          | character varying(255) | not null
 email_verified | bool                   | not null
 demo_balance   | uint                   | not null
Indexes:
    "users_pkey" PRIMARY KEY, btree (id)
    "users_email_key" UNIQUE CONSTRAINT, btree (email)
    "users_auth0_uuid_key" UNIQUE CONSTRAINT, btree (auth0_uuid)
*/
