package main

import (
	"strconv"

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
	} else if checkRegexp("!invite \\d", m.Content) {
		num := makeNumber(m.Content)

		if err := onInvite(s, m, num); err != nil {
			return
		}
	} else if checkRegexp("!invite", m.Content) {
		if err := onInvite(s, m, 0); err != nil {
			return
		}
	}
}

func onHelp(s *discordgo.Session, m *discordgo.MessageCreate) (err error) {
	footer := discordgo.MessageEmbedFooter{Text: "🍎 ご要望、不具合は https://github.com/cohky16/invite までお願いします"}

	embed := discordgo.MessageEmbed{
		Title: "機能概要",
		Description: "ボイスチャンネルへの招待を送信できます\n\n" +
			"**__各種コマンド__**\n\n" +
			"**invite**\nユーザーにボイスチャンネルへの招待情報を送信します\n`!invite @hoge @fuga`\n\n" +
			"**invite $(RoomNo)**\nユーザーに特定のボイスチャンネルへの招待情報を送信します\n`!invite 1 @hoge @fuga`\n\n" +
			"**help**\nヘルプを表示します",
		Footer: &footer,
	}

	if err = setCommand(s, m); err != nil {
		return
	}

	session := newSession(s, m.Author.Bot)

	if _, err = session.ChannelMessageSendEmbed(m.ChannelID, &embed); err != nil {
		return
	}

	return
}

func setCommand(s *discordgo.Session, m *discordgo.MessageCreate) (err error) {
	session := newSession(s, m.Author.Bot)

	_, err = session.ApplicationCommandCreate(
		s.State.User.ID,
		m.GuildID,
		&discordgo.ApplicationCommand{
			Name:        "help",
			Description: "ヘルプを表示します",
		},
	)

	_, err = session.ApplicationCommandCreate(
		s.State.User.ID,
		m.GuildID,
		&discordgo.ApplicationCommand{
			Name:        "invite",
			Description: "ユーザーにボイスチャンネルへの招待情報を送信します",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "to",
					Description: "招待したいユーザー名を入力します @hoge @fuga Tabでチャンネルの指定へ、Enterで確定します",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
				{
					Name:        "channel",
					Description: "招待したいチャンネルを指定します",
					Type:        discordgo.ApplicationCommandOptionChannel,
					Required:    false,
				},
			},
		},
	)

	return
}

func onInvite(s *discordgo.Session, m *discordgo.MessageCreate, n int) (err error) {
	session := newSession(s, m.Author.Bot)

	c, err := session.Channel(m.ChannelID)

	if err != nil {
		return
	}

	if checkRegexp("talk", c.Name) {
		if err = sendMessage(s, m, c, "Talk", n); err != nil {
			return
		}
	} else if checkRegexp("meeting", c.Name) {
		if err = sendMessage(s, m, c, "Meeting", n); err != nil {
			return
		}
	}

	return
}

func sendMessage(s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel, r string, n int) (err error) {
	users, err := makeUsers(s, m, c)

	if err != nil {
		return
	}

	state, _ := s.State.VoiceState(m.GuildID, m.Author.ID)

	for _, channel := range users.Channels {
		if n > 0 && channel.Type == 2 && checkRegexp(r+".*"+strconv.Itoa(n), channel.Name) {
			if err = sendInvite(s, m, channel, users.Users); err != nil {
				return
			}
			break
		} else if n == 0 && state != nil && state.ChannelID == channel.ID {
			if err = sendInvite(s, m, channel, users.Users); err != nil {
				return
			}
			break
		} else if n == 0 && state == nil {
			count := 0

			for _, alreadyChannelID := range users.AlreadyChannelIds {
				if alreadyChannelID == channel.ID {
					count++
				}
			}

			if count < 2 && channel.Type == 2 && checkRegexp(r, channel.Name) {
				if err = sendInvite(s, m, channel, users.Users); err != nil {
					return
				}
				break
			}
		}
	}

	return
}

func sendInvite(s *discordgo.Session, m *discordgo.MessageCreate, channel *discordgo.Channel, users string) (err error) {
	session := newSession(s, m.Author.Bot)

	st, err := session.ChannelInviteCreate(channel.ID, discordgo.Invite{})

	if err != nil {
		return
	}

	ms, err := session.ChannelMessageSend(m.ChannelID, users+"\nボイスチャンネルに招待されました\n"+"https://discord.gg/"+st.Code)

	if err != nil || ms == nil {
		return
	}

	return
}
