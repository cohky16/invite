package main

import (
	main_internal "invite/internal"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

type Session interface {
	ChannelMessageSendEmbed(channelID string, embed *discordgo.MessageEmbed) (*discordgo.Message, error)
	GuildChannels(guildID string) (st []*discordgo.Channel, err error)
	GuildMembers(guildID string, after string, limit int) (st []*discordgo.Member, err error)
	Channel(channelID string) (st *discordgo.Channel, err error)
	ChannelInviteCreate(channelID string, i discordgo.Invite) (st *discordgo.Invite, err error)
	ChannelMessageSend(channelID string, content string) (*discordgo.Message, error)
	ApplicationCommandCreate(appID string, guildID string, cmd *discordgo.ApplicationCommand) (*discordgo.ApplicationCommand, error)
	InteractionRespond(*discordgo.Interaction, *discordgo.InteractionResponse) error
}

func main() {
	token, err := getToken()

	if err != nil {
		log.Fatal(err)
	}

	dg, err := newDg(token)

	if err != nil {
		log.Fatal(err)
	}

	err = openDg(dg)

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
	if os.Getenv("APP_ENV") != "production" && os.Getenv("CI_ENV") != "TRUE" {
		err := godotenv.Load("../../../.env")

		if err != nil {
			return "", err
		}
	}

	token = os.Getenv("DISCORD_TOKEN")

	return
}

func newDg(token string) (dg *discordgo.Session, err error) {
	dg, err = discordgo.New("Bot " + token)

	if err != nil {
		return nil, err
	}

	return
}

func openDg(dg *discordgo.Session) (err error) {
	dg.AddHandler(onMessageCreate)
	dg.AddHandler(onCommand)

	err = dg.Open()

	if err != nil {
		return
	}

	return
}

func newSession(s *discordgo.Session, bot bool) Session {
	if bot {
		var mockSession main_internal.MockSession
		return mockSession
	}
	return s
}
