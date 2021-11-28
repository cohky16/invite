package main

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestOnMessageCreate(t *testing.T) {
	s, err := mockDg()

	if err != nil {
		t.Error(err)
	}

	m := &discordgo.MessageCreate{&discordgo.Message{Author: s.State.User}}

	onMessageCreate(s, m)
}

func TestOnHelp(t *testing.T) {
	s, err := mockDg()

	if err != nil {
		t.Error(err)
	}

	m := &discordgo.MessageCreate{&discordgo.Message{Author: s.State.User}}

	err = onHelp(s, m)

	if err != nil {
		t.Error(err)
	}
}

func TestOnInvite(t *testing.T) {
	s, err := mockDg()

	if err != nil {
		t.Error(err)
	}

	m := &discordgo.MessageCreate{&discordgo.Message{Author: s.State.User}}

	err = onInvite(s, m)

	if err != nil {
		t.Error(err)
	}
}

func TestSendMessage(t *testing.T) {
	s, err := mockDg()

	if err != nil {
		t.Error(err)
	}

	m := &discordgo.MessageCreate{&discordgo.Message{Author: s.State.User}}

	err = sendMessage(s, m, s.State.Guilds[0].Channels[0], "Talk")

	if err != nil {
		t.Error(err)
	}
}
