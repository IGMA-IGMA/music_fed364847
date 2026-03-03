package main

type FakeUser struct {
	Username string `faker:"first_name"`
	Email    string `faker:"email"`
	Pwd      string `faker:"password"`
}

type UserJS struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Pwd      string `json:"pwd"`
}
