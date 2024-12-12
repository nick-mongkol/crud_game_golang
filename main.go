package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type game struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Publisher string  `json:"publisher"`
	Price     float64 `json:"price"`
}

var games = []game{
	{ID: "1", Title: "GTA V", Publisher: "Rockstar Games", Price: 50},
	{ID: "2", Title: "God of War", Publisher: "Sony", Price: 60},
	{ID: "3", Title: "Red Dead Redemption 2", Publisher: "Rockstar Games", Price: 70},
}

func getGames(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, games)
}

func addGames(context *gin.Context) {
	var newGame game

	// bind json baru ke newGame
	if err := context.BindJSON(&newGame); err != nil {
		return
	}

	games = append(games, newGame)

	context.IndentedJSON(http.StatusCreated, newGame)
}

func main() {
	fmt.Print("Start server")
	router := gin.Default()
	router.GET("/coba", getGames)
	router.POST("/coba", addGames)
	router.Run("localhost:8888")
}
