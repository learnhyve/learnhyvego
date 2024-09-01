package main

import (
	"learnhyve/controllers"
	"learnhyve/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadInitializers()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.GET("/getAllposts", controllers.GetAllPosts)
	r.GET("/getPost/:id", controllers.FetchPost)
	r.POST("/createPost", controllers.CreatePosts)
	r.POST("/updatePost/:id", controllers.UpdatePost)
	r.DELETE("/deletePost/:id", controllers.DeletePost)

	r.Run()
}
