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
	footer := discordgo.MessageEmbedFooter{Text: "ð ãè¦æãä¸å·åã¯ https://github.com/cohky16/invite ã¾ã§ãé¡ããã¾ã"}

	embed := discordgo.MessageEmbed{
		Title: "æ©è½æ¦è¦",
		Description: "ãã¤ã¹ãã£ã³ãã«ã¸ã®æå¾ãéä¿¡ã§ãã¾ã\n\n" +
			"**__åç¨®ã³ãã³ã__**\n\n" +
			"**invite**\nã¦ã¼ã¶ã¼ã«ãã¤ã¹ãã£ã³ãã«ã¸ã®æå¾æå ±ãéä¿¡ãã¾ã\n`!invite @hoge @fuga`\n\n" +
			"**invite $(RoomNo)**\nã¦ã¼ã¶ã¼ã«ç¹å®ã®ãã¤ã¹ãã£ã³ãã«ã¸ã®æå¾æå ±ãéä¿¡ãã¾ã\n`!invite 1 @hoge @fuga`\n\n" +
			"**help**\nãã«ããè¡¨ç¤ºãã¾ã",
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
			Description: "ãã«ããè¡¨ç¤ºãã¾ã",
		},
	)

	_, err = session.ApplicationCommandCreate(
		s.State.User.ID,
		m.GuildID,
		&discordgo.ApplicationCommand{
			Name:        "invite",
			Description: "ã¦ã¼ã¶ã¼ã«ãã¤ã¹ãã£ã³ãã«ã¸ã®æå¾æå ±ãéä¿¡ãã¾ã",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "to",
					Description: "æå¾ãããã¦ã¼ã¶ã¼åãå¥åãã¾ã @hoge @fuga Tabã§ãã£ã³ãã«ã®æå®ã¸ãEnterã§ç¢ºå®ãã¾ã",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
				{
					Name:        "channel",
					Description: "æå¾ããããã£ã³ãã«ãæå®ãã¾ã",
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

	ms, err := session.ChannelMessageSend(m.ChannelID, users+"\nãã¤ã¹ãã£ã³ãã«ã«æå¾ããã¾ãã\n"+"https://discord.gg/"+st.Code)

	if err != nil || ms == nil {
		return
	}

	return
}
