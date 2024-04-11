package main

import "time"

type User struct {
	FirstName string
	Lastname  string
	Service   string
	JoinDate  time.Time
	Active    bool
}

func (u User) GetFirstname()             {}
func (u User) GetLastname()              {}
func (u User) GetService()               {}
func (u User) GetJoinDate()              {}
func (u User) GetActive()                {}
func (u User) Complex(int, int, string)  {}
func (u User) Complex1(int, int, string) {}
func (u User) Complex2(int, int, string) {}
func (u User) Complex3(int, int, string) {}

type UserModel1 interface {
	Complex1(int, int, string)
}

type UserModel interface {
	GetFirstname()
	GetLastname()
	GetService()
	GetJoinDate()
	GetActive()
}

func ProcessUser(user UserModel1) {

}

func main() {

}
