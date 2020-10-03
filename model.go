package main

import (
	"time"
)

type GormModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type User struct {
	GormModel
	Name                  string `json:"name"`
	Email                 string `json:"email"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	RegistrationCompleted bool   `json:"registrationCompleted"`
}

type Role struct {
	GormModel
	Name string `json:"name"`
}

type Partner struct {
	GormModel
	FamilyIDNumber string      `json:"familyIdNumber"`
	FirstName      string      `json:"firstName"`
	LastName       string      `json:"lastName"`
	Name           string      `json:"name"`
	Phone          string      `json:"phone"`
	Address        string      `json:"address"`
	Business       Business    `json:"business"`
	BusinessID     uint        `json:"businessId"`
	City           string      `json:"city"`
	Position       string      `json:"position"`
	Sex            string      `json:"sex"`
	Bank           Bank        `json:"bank"`
	BankID         uint        `json:"bankId"`
	BankHolderName string      `json:"bankHolderName"`
	ServiceType    ServiceType `json:"serviceType"`
	ServiceTypeID  uint        `json:"serviceTypeId"`
	Lat            float32     `json:"lat"`
	Lon            float32     `json:"lon"`
	User           User        `json:"user"`
	UserID         uint        `json:"userId"`
}

type ServiceType struct {
	GormModel
	Name string `json:"name"`
}

type Business struct {
	GormModel
	Name        string `json:"name"`
	Address     string `json:"address" gorm:"type:text"`
	Description string `json:"description" gorm:"type:text"`
}

type Bank struct {
	GormModel
	Name string `json:"name"`
	Code uint   `json:"code"`
}

type Building struct {
	GormModel
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Lat     string `json:"lat"`
	Lon     string `json:"lon"`
}

type Room struct {
	GormModel
	Name               string           `json:"name"`
	Number             string           `json:"number"`
	RoomType           RoomType         `json:"roomType"`
	RoomTypeID         uint             `json:"roomTypeId"`
	HasAC              bool             `json:"hasAc"`
	FoodAccomodation   FoodAccomodation `json:"foodAccomodation"`
	FoodAccomodationID uint             `json:"foodAccomodationId"`
	CancellationFee    CancellationFee  `json:"cancellationFee"`
	BedCapacity        uint             `json:"bedCapacity"`
	GuestType          GuestType        `json:"guestType"`
	GuestTypeID        uint             `json:"guestTypeId`
	Price              uint             `json:"price"`
	Description        string           `json:"description" gorm:"type:text"`
}

type GuestType struct {
	GormModel
	Name string `json:"name"`
}

type CancellationFee struct {
	GormModel
	Name string `json:"name"`
}

type RoomType struct {
	GormModel
	Name string `json:"name"`
}

type FoodAccomodation struct {
	GormModel
	Name string `json:"name"`
}

type ComingSoonEmail struct {
	GormModel
	Email string `json:"email"`
}

type Car struct {
	GormModel
	Name              string   `json:"name"`
	FuelType          FuelType `json:"fuelType"`
	FuelTypeID        uint     `json:"fuelTypeId"`
	Capacity          uint     `json:"capacity"`
	CarType           CarType  `json:"carType"`
	CarTypeID         uint     `json:"carTypeId"`
	RegistrationPlate string   `json:"registrationPlate"`
	Year              uint     `json:"year"`
	Description       string   `json:"description" gorm:"type:text"`
}

type CarType struct {
	GormModel
	Name string `json:"name"`
}

type FuelType struct {
	GormModel
	Name string `json:"name"`
}

type Laundry struct {
	GormModel
	Price         uint        `json:"price"`
	IsPerKg       bool        `json:"isPerKg"`
	LaundryType   LaundryType `json:"laundryType"`
	LaundryTypeID uint        `json:"laundryTypeId"`
	Description   string      `json:"description" gorm:"type:text"`
}

type LaundryType struct {
	GormModel
	Name string `json:"name"`
}
