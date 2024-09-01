package main

import (
	"learnhyve/initializers"
	"learnhyve/models"
)

func init(){
initializers.LoadInitializers()
initializers.ConnectToDB()	
}

func main(){
	initializers.DB.AutoMigrate(&models.Post{})

}