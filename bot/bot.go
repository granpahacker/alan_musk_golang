// Copyright (C) GRANDPA HACKER - All Rights Reserved
package bot

import (
	"context"
	"fmt"
	"golang-discord-bot/config"
	"log"
	"os"
	"strings"

	// discord bot library

	gpt3 "github.com/PullRequestInc/go-gpt3"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var BotId string
var goBot *discordgo.Session

var currentStatus = 0

type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

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

	// message answers
	goBot.AddHandler(messageHandler)
	// bot activity
	goBot.AddHandler(setActivity)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("No errors. Alan Musk kiss your butt!!")
}

// bot messages handler
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	var userID = m.Author.ID

	if userID == BotId {
		return
	}

	// response user data
	var message_s = s
	var message_d = m

	// clear command
	if strings.Contains(m.Content, ".к") {

		if len([]rune(m.Content)) > 3 {
			b := m.Content[4:]

			response := "!clear " + b
			_, _ = s.ChannelMessageSend(m.ChannelID, response)
		}

	}

	// chat gpt command
	if strings.Contains(m.Content, ".п") {

		if currentStatus == 1 {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Зачекайте оброблення попереднього запиту")
			return
		}

		//Stop message handler
		currentStatus = 1

		// timeout handler
		log.SetOutput(new(NullWriter))

		apiKey := goDotEnvVariable("API_KEY")

		fmt.Println(userID)

		if apiKey == "" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Виникла помилка.")
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, "Очікуйте відповіді:")

		ctx := context.Background()
		client := gpt3.NewClient(apiKey)

		question := m.Content[4:]
		questionParam := validateQuestion(question)
		GetResponse(client, ctx, questionParam, message_s, message_d)

		//с ontinue message handler
		currentStatus = 0
	}
}

func setActivity(s *discordgo.Session, r *discordgo.Ready) {
	err := s.UpdateListeningStatus(".п; .к;")
	if err != nil {
		panic(err)
	}
}

// .ENV file reader
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func validateQuestion(question string) string {
	quest := strings.Trim(question, " ")
	keywords := []string{"", "loop", "break", "continue", "cls", "exit", "block"}
	for _, x := range keywords {
		if quest == x {
			return ""
		}
	}
	return quest
}
