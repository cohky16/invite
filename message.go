package main

import (
	"github.com/bwmarrin/discordgo"
)

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if checkRegexp("!help", m.Content) {
		if err := onHelp(s, m); err != nil {
			return
		}

	} else if checkRegexp("!invite", m.Content) {
		if err := onInvite(s, m); err != nil {
			return
		}
	}
}

func onHelp(s *discordgo.Session, m *discordgo.MessageCreate) (err error) {
	footer := discordgo.MessageEmbedFooter{Text: "🍎 ご要望、不具合は https://github.com/cohky16/invite までお願いします"}

	embed := discordgo.MessageEmbed{
		Title:       "機能概要",
		Description: "ボイスチャンネルへの招待を送信できます\n\n**__各種コマンド__**\n\n**invite**\nユーザーにボイスチャンネルへの招待情報を送信します\n`!invite @hoge @fuga`\n\n**help**\nヘルプを表示します",
		Footer:      &footer,
	}

	session := newSession(s, m)

	_, err = session.ChannelMessageSendEmbed(m.ChannelID, &embed)

	return
}

func onInvite(s *discordgo.Session, m *discordgo.MessageCreate) (err error) {
	session := newSession(s, m)

	c, err := session.Channel(m.ChannelID)

	if err != nil {
		return
	}

	if checkRegexp("talk", c.Name) {
		if err = sendMessage(s, m, c, "Talk"); err != nil {
			return
		}
	} else if checkRegexp("meeting", c.Name) {
		if err = sendMessage(s, m, c, "Meeting"); err != nil {
			return
		}
	} else {
		return
	}

	return
}

func sendMessage(s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel, r string) error {
	users, err := makeUsers(s, m, c)

	if err != nil {
		return err
	}

	session := newSession(s, m)

	for _, channel := range users.Channels {
		count := 0

		for _, alreadyChannelId := range users.AlreadyChannelIds {
			if alreadyChannelId == channel.ID {
				count++
			}
		}

		if count < 2 && channel.Type == 2 && checkRegexp(r, channel.Name) {
			st, err := session.ChannelInviteCreate(channel.ID, discordgo.Invite{})

			if err != nil {
				return err
			}

			ms, err := session.ChannelMessageSend(m.ChannelID, users.Users+"\nボイスチャンネルに招待されました\n"+"https://discord.gg/"+st.Code)

			if err != nil || ms == nil {
				return err
			}
			break
		}
	}

	return nil
}
