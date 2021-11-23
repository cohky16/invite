package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	token, err := getToken()

	if err != nil {
		log.Fatal(err)
	}

	dg, err := openDg(token)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening...")

	stopBot := make(chan os.Signal, 1)

	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-stopBot

	err = dg.Close()

	if err != nil {
		log.Fatal(err)
	}

	return
}

func getToken() (token string, err error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()

		if err != nil {
			return "", err
		}
	}

	token = os.Getenv("DISCORD_TOKEN")

	return
}

func openDg(token string) (dg *discordgo.Session, err error) {
	dg, err = discordgo.New("Bot " + token)

	if err != nil {
		return nil, err
	}

	dg.AddHandler(onMessageCreate)

	err = dg.Open()

	if err != nil {
		return nil, err
	}

	return
}
