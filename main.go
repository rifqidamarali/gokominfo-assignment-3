package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	data := map[string]int{}

	go func() {
		for {
			// Update data
			data["water"] = rand.Intn(100)
			data["wind"] = rand.Intn(100)
			time.Sleep(15 * time.Second)
		}
	}()

	// Apply CORS middleware
	g.Use(cors.Default())

	g.GET("/data", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"status": data,
		})
	})

	g.Run(":8080")
}
