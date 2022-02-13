package main

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestOnCommand(t *testing.T) {
	s, err := mockDg()

	if err != nil {
		t.Error(err)
	}

	i := &discordgo.InteractionCreate{Interaction: &discordgo.Interaction{Member: &discordgo.Member{User: s.State.User}}}

	onCommand(s, i)
}

func TestOnHelpForCommand(t *testing.T) {
	s, err := mockDg()

	if err != nil {
		t.Error(err)
	}

	i := &discordgo.InteractionCreate{Interaction: &discordgo.Interaction{Member: &discordgo.Member{User: s.State.User}}}

	onHelpForCommand(s, i)
}

func TestOnInviteForCommand(t *testing.T) {
	s, err := mockDg()

	if err != nil {
		t.Error(err)
	}

	i := &discordgo.InteractionCreate{Interaction: &discordgo.Interaction{Member: &discordgo.Member{User: s.State.User}, Data: discordgo.ApplicationCommandInteractionData{Name: "test"}, GuildID: "1"}}

	onInviteForCommand(s, i)
}

func TestMakeChannelId(t *testing.T) {
	s, err := mockDg()

	if err != nil {
		t.Error(err)
	}

	i := &discordgo.InteractionCreate{Interaction: &discordgo.Interaction{Member: &discordgo.Member{User: s.State.User}, Data: discordgo.ApplicationCommandInteractionData{Name: "test"}, GuildID: "1"}}

	makeChannelId(s, i)
}

func TestMakeSenderChannelName(t *testing.T) {
	s, err := mockDg()

	if err != nil {
		t.Error(err)
	}

	i := &discordgo.InteractionCreate{Interaction: &discordgo.Interaction{Member: &discordgo.Member{User: s.State.User}, Data: discordgo.ApplicationCommandInteractionData{Name: "test"}, GuildID: "1"}}

	makeSenderChannelName(s, i)
}

func TestMakeAlreadyChannelIds(t *testing.T) {
	s, err := mockDg()

	if err != nil {
		t.Error(err)
	}

	i := &discordgo.InteractionCreate{Interaction: &discordgo.Interaction{Member: &discordgo.Member{User: s.State.User}, Data: discordgo.ApplicationCommandInteractionData{Name: "test"}, GuildID: "1"}}

	makeAlreadyChannelIds(s, i)
}
