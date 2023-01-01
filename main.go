package main

import "github.com/gin-gonic/gin"

func main() {
	// remove this if you don't want to run in release mode
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Static("/static", "./build/static")
	router.Static("/favicon.ico", "./build")
	router.Static("/manifest.json", "./build")

	router.NoRoute(func(c *gin.Context) {
		c.File("./build/index.html")
	})

	router.Run(":8080")
}
