package user

import (
	"iw2/testing/user/db"
	"iw2/testing/user/model"
)

func GetUserById(db db.Db, id int) (model.User, error) {
	return db.QueryUser("SELECT * FROM users WHERE id = ?", id)
}
