package controllers

import (
	"fmt"
	"learnhyve/initializers"
	"learnhyve/models"

	"github.com/gin-gonic/gin"
)

func CreatePosts(c *gin.Context) {
	var body struct {
		Title  string
		Body   string
		Author string
		Stars  int
	}
	c.Bind(&body)
	fmt.Println(body)
	post := models.Post{Title: body.Title, Body: body.Body, Author: body.Author, Stars: int32(body.Stars)}

	res := initializers.DB.Create(&post)
	if res.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, res)
}
func GetAllPosts(c *gin.Context) {

	var posts []models.Post
	res := initializers.DB.Find(&posts)

	if res.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func FetchPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	res := initializers.DB.First(&post, id)

	if res.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"result": post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title  string
		Body   string
		Author string
		Stars  int
	}
	c.Bind(&body)
	fmt.Println(body)
	var post models.Post
	initializers.DB.First(&post, id)
	res := initializers.DB.Model(&post).Updates(models.Post{Title: body.Title,
		Body: body.Body, Author: body.Author, Stars: int32(body.Stars)})
	if res.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"result": res,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	res := initializers.DB.Delete(&models.Post{}, id)
	if res.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"result": res,
	})
}
