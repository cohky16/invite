package main

import "github.com/bwmarrin/discordgo"

type Users struct {
	users             string
	channels          []*discordgo.Channel
	alreadyChannelIds []string
}

func makeUsers(s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel) (Users, error) {
	var users Users

	for _, user := range m.Mentions {
		users.users += user.Mention() + " "
	}

	channels, err := s.GuildChannels(c.GuildID)

	if err != nil {
		return users, err
	}

	users.channels = channels

	members, err := s.GuildMembers(m.GuildID, "", 1000)

	if err != nil {
		return users, err
	}

	for _, member := range members {
		state, err := s.State.VoiceState(m.GuildID, member.User.ID)

		if state == nil || err != nil {
			continue
		}

		users.alreadyChannelIds = append(users.alreadyChannelIds, state.ChannelID)
	}

	return users, nil
}
