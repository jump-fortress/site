package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	database "github.com/spiritov/jump/api/db/queries"
)

func main() {
	db, err := sql.Open("sqlite3", "db/jump.db")
	if err != nil {
		log.Fatalf("[fatal] failed to open db: %v", err)
	}
	defer db.Close()

	queries := database.New(db)
	ctx := context.Background()

	player, err := queries.InsertPlayer(ctx, database.InsertPlayerParams{
		SteamID:     "1",
		SteamPfpID:  "1",
		DisplayName: "spiritov",
	})
	if err != nil {
		//log.Fatal(err)
	} else {
		log.Println(player)
	}

	r := gin.Default()

	r.GET("/players/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		_, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			return
		}

		// select player from db
		player, err := queries.SelectPlayer(ctx, 1)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "player not found"})
		}

		c.JSON(http.StatusOK, player)
	})

	r.POST("/players", func(c *gin.Context) {
		var player database.InsertPlayerParams

		if err := c.BindJSON(&player); err != nil {
			return
		}

		// add player to db
		queries.InsertPlayer(ctx, player)
		c.JSON(http.StatusCreated, player)
	})

	r.Run("localhost:8080")
}
