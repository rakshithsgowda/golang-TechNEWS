package models

import "github.com/upper/db/v4"

// use the propable models to use the session data.
type Models struct {
	Users    UsersModel
	Posts    PostsModel
	Comments CommentsModel
}


// use the eror constants 
var (
// no records found
// email already in db
// account is inactive.
// invalid login


)




// new sesssion for the users..
func New(db db.Session)Models{
	return Models{
		Users: UsersModel{db:db},
		Posts: postsModel{db:db},
		Comments: CommentsModel{db:db},
	}
}