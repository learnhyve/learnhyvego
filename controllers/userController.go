package controllers

import (
	"fmt"
	"learnhyve/initializers"
	"learnhyve/models"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var body struct {
		Fname   string
		Lname   string
		Pnumber int32
		Course  string
		Token   string
	}

	c.Bind(&body)
	fmt.Println(body)
	post := models.User{
		Fname:    body.Fname,
		Lname:    body.Lname,
		Pnumber:  body.Pnumber,
		Course:   body.Course,
		FCMToken: body.Token,
	}

	res := initializers.DB.Create(&post)
	if res.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, res)
}
