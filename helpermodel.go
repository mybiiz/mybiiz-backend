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

type Page struct {
	PageInfo
	Content interface{} `json:"content"`
}

type PageInfo struct {
	TotalElements int64 `json:"totalElements"`
	Page          int   `json:"page"`
	PerPage       int   `json:"perPage"`
	Last          int64 `json:"last"`
}
