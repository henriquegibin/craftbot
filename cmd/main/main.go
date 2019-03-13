package main

import (
	"fmt"
	"log"
	"regexp"
	"time"

	bot "github.com/Tnze/gomcbot"
	colorable "github.com/mattn/go-colorable"
)

func main() {

	for {
		daemon()
		log.Println("dormindo...")
		time.Sleep(time.Duration(100) * time.Second)

	}
}

func daemon() {
	for {
		log.SetOutput(colorable.NewColorableStdout())
		auth := bot.Auth{Name: "capiroto"}
		client, err := auth.JoinServer("MY URL HERE", 25565)
		if err != nil {
			log.Fatal(err)
		}

		go client.HandleGame()
		for event := range client.GetEvents() {
			stringSlice := fmt.Sprint(event)
			match, _ := regexp.MatchString("sleepnow", stringSlice) // palavra magina é sleepnow

			if match { // por causa do rountimes,o kill vai demorar um pouco para matar o bot
				client.Chat("capiroto vai dormir também...")
				return
			}
		}
		return
	}
}
