package main

import (
	"log"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal(err)
		}
	}

	token := os.Getenv("DISCORD_TOKEN")

	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	dg.AddHandler(onMessageCreate)

	err = dg.Open()

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

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if checkRegexp("!help", m.Content) {
		footer := discordgo.MessageEmbedFooter{Text: "🍎 ご要望、不具合は https://github.com/cohky16/invite までお願いします"}

		embed := discordgo.MessageEmbed{
			Title:       "機能概要",
			Description: "ボイスチャンネルへの招待を送信できます\n\n**__各種コマンド__**\n\n**invite**\nユーザーにボイスチャンネルへの招待情報を送信します\n`!invite @hoge @fuga`\n\n**help**\nヘルプを表示します",
			Footer:      &footer,
		}

		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)

		if err != nil {
			return
		}

	} else if checkRegexp("!invite", m.Content) {
		c, err := s.Channel(m.ChannelID)

		if err != nil {
			return
		}

		if checkRegexp("Talk", c.Name) {
			err := sendMessage(s, m, c, "talk")

			if err != nil {
				return
			}
		} else if checkRegexp("Meeting", c.Name) {
			err := sendMessage(s, m, c, "meeting")

			if err != nil {
				return
			}
		} else {
			return
		}
	}

}

func checkRegexp(reg, str string) bool {
	return regexp.MustCompile("^" + reg + ".*$").Match([]byte(str))
}

func sendMessage(s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel, r string) error {
	users := ""

	for _, user := range m.Mentions {
		users += user.Mention() + " "
	}

	channels, err := s.GuildChannels(c.GuildID)

	if err != nil {
		return err
	}

	alreadyChannelIds, err := makeAlreadyChannelIds(s, m)

	if err != nil {
		return err
	}

	for _, channel := range channels {
		count := 0

		for _, alreadyChannelId := range alreadyChannelIds {
			if alreadyChannelId == channel.ID {
				count++
			}
		}

		if count < 2 && channel.Type == 2 && checkRegexp(r, channel.Name) {
			st, err := s.ChannelInviteCreate(channel.ID, discordgo.Invite{})

			if err != nil {
				return err
			}

			s.ChannelMessageSend(m.ChannelID, users+"\nボイスチャンネルに招待されました\n"+"https://discord.gg/"+st.Code)

			break
		}
	}

	return nil
}

func makeAlreadyChannelIds(s *discordgo.Session, m *discordgo.MessageCreate) ([]string, error) {
	members, err := s.GuildMembers(m.GuildID, "", 1000)

	if err != nil {
		return nil, err
	}

	alreadyChannelIds := []string{}

	for _, member := range members {
		state, err := s.State.VoiceState(m.GuildID, member.User.ID)

		if state == nil || err != nil {
			continue
		}

		alreadyChannelIds = append(alreadyChannelIds, state.ChannelID)
	}

	return alreadyChannelIds, nil
}
