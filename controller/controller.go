package controller

import (
	"example/album-serive-go/database"
	"example/album-serive-go/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetALbums(c *gin.Context) {
	db := database.DBPool()
	log.Println(db)
	var albums []types.Album
	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	defer rows.Close()
	for rows.Next() {
		var alb types.Album
		if err := rows.Scan(&alb.ID, &alb.Artist, &alb.Title, &alb.Price); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	c.IndentedJSON(http.StatusOK, albums)
}
