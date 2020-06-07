package main

import (
	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/profile/:id", inDB.GetPerson)
	router.GET("/profiles", inDB.GetAllPerson)
	router.POST("/profile", inDB.CreatePerson)
	router.PUT("/profile/:id", inDB.UpdatePerson)
	router.DELETE("/profile/:id", inDB.DeletePerson)
	router.Run(":3000")
}
