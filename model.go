package main

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type Role struct {
	gorm.Model
	Name string `json:"name"`
}

type Partner struct {
	gorm.Model
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Lat     string `json:"lat"`
	Lon     string `json:"lon"`
}

type Building struct {
	gorm.Model
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Lat     string `json:"lat"`
	Lon     string `json:"lon"`
}

type Room struct {
	gorm.Model
	Name   string `json:"name"`
	Number string `json:"number"`
}

type ComingSoonEmail struct {
	gorm.Model
	Email string `json:"email"`
}
