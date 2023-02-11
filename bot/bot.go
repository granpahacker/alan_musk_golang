// Copyright (C) Oleg Lysiak - All Rights Reserved
package bot

import (
	"fmt"
	"golang-discord-bot/config"

	// discord bot library
	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

// Start func
func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("No errors. Lviv Alarm is online!")
}

// bot messages handler
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
}
