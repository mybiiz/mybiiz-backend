package main

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserPostBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

type RegisterPostBody struct {
	User    User    `json:"user"`
	Partner Partner `json:"partner"`
}
