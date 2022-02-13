package main

import "github.com/bwmarrin/discordgo"

type Users struct {
	Users             string
	Channels          []*discordgo.Channel
	AlreadyChannelIds []string
}

func makeUsers(s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel) (Users, error) {
	var users Users

	for _, user := range m.Mentions {
		users.Users += user.Mention() + " "
	}

	session := newSession(s, m.Author.Bot)

	channels, err := session.GuildChannels(c.GuildID)

	if err != nil {
		return users, err
	}

	users.Channels = channels

	members, err := session.GuildMembers(m.GuildID, "", 1000)

	if err != nil {
		return users, err
	}

	for _, member := range members {
		state, err := s.State.VoiceState(m.GuildID, member.User.ID)

		if state == nil || err != nil {
			continue
		}

		users.AlreadyChannelIds = append(users.AlreadyChannelIds, state.ChannelID)
	}

	return users, nil
}
