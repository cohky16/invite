package main

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestMakeUsers(t *testing.T) {
	s, err := mockDg()

	if err != nil {
		t.Error(err)
	}

	m := &discordgo.MessageCreate{&discordgo.Message{Author: s.State.User}}

	users, err := makeUsers(s, m, s.State.Guilds[0].Channels[0])

	if err != nil {
		t.Error(err)
	}

	if len(users.AlreadyChannelIds) != 0 || users.Users != "" || len(users.Channels) != 1 {
		t.Error("error users")
	}
}
