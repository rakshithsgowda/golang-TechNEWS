package models

import "time"

// models used ofr authentication

type User struct {
	ID 				int `db:"id, omitempty"`
	Name 			string `db:"name"`
	Email 		string `db:"email"`
	Password 	string `db:"password_hash"`
	CreatedAt time.Time `db:"created_at"`
	
}

// find a user
// find the user by user id
// insert a user using the password 

// Authentication. using the compare password using the jwt feature.
