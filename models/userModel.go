package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fname   string
	Lname   string
	Pnumber int32
	Course  string
	FCMToken  string
}
