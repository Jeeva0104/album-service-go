package main

import (
	"database/sql"
	"example/album-serive-go/controller"
	"example/album-serive-go/database"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {
	database.DBConnect()
	router := gin.Default()
	router.GET("/albums", controller.GetALbums)
	router.Run("localhost:3000")

}
