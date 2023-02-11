// Copyright (C) GRANDPA HACKER - All Rights Reserved
package main

import (
	"fmt"
	"golang-discord-bot/bot"
	"golang-discord-bot/config"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	err := config.ReadConfig()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// check for errors

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	// router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	// start bot
	bot.Start()

	// start server
	router.Run(":" + port)

	<-make(chan struct{})

	return
}
