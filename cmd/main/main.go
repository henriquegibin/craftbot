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
		client, err := auth.JoinServer("minecraft.gibin.club", 25565)
		if err != nil {
			log.Fatal(err)
		}

		go client.HandleGame()
		for event := range client.GetEvents() {
			stringSlice := fmt.Sprint(event)

			match, _ := regexp.MatchString("sleepnow", stringSlice)
			if match {
				client.Chat("capiroto vai dormir também...")
				return
			}
		}
		return
	}
}
