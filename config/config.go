// Copyright (C) GRANDPA HACKER - All Rights Reserved
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	// json token and bot prefix
	Token     string `json : "MTA228NzA1OTM4NzY3MzM0NjA3OA.GpnaHG.HAY9unzUaKfLCFu3a72ptqCEZyU6ELC2xa9rPs"`
	BotPrefix string `json : "$$"`
}

// reading of .json file
func ReadConfig() error {
	fmt.Println("Reading config file...")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
