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

	tests := []struct {
		ChannelID string
		N         int
	}{
		{"talk", 1},
		{"meeting", 1},
		{"talk", 0},
		{"meeting", 0},
		{"other", 1},
		{"other", 1},
		{"other", 0},
		{"other", 0},
	}

	for _, test := range tests {
		m.ChannelID = test.ChannelID

		err = onInvite(s, m, test.N)

		if err != nil {
			t.Error(err)
		}
	}

}

func TestSendMessage(t *testing.T) {
	s, err := mockDg()

	if err != nil {
		t.Error(err)
	}

	m := &discordgo.MessageCreate{&discordgo.Message{Author: s.State.User}}

	err = sendMessage(s, m, s.State.Guilds[0].Channels[0], "Talk", 1)

	if err != nil {
		t.Error(err)
	}
}

func TestSendInvite(t *testing.T) {
	s, err := mockDg()

	if err != nil {
		t.Error(err)
	}

	m := &discordgo.MessageCreate{&discordgo.Message{Author: s.State.User}}

	err = sendInvite(s, m, s.State.Guilds[0].Channels[0], "@hogefuga")

	if err != nil {
		t.Error(err)
	}
}
