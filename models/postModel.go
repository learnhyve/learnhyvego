package models

import (

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string
	Body string
	Stars int32
	Author string

  }