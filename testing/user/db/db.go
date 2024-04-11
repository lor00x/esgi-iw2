package db

import "iw2/testing/user/model"

type Db interface {
	QueryUser(string, int) (model.User, error)
}
