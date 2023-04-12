package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, i'm Cryptonit. This dapp for asic repair center! Follow me",
		})
	})

	router.Run("https://www.cryptonit.fun/navi")
}
