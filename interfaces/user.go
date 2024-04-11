package main

type User struct {
	name string
	active bool
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) IsAdmin() bool {
	return false
}

func (u *User) SetActive(active bool) {
	u.active = active
}

