package main

import (
	"net/http"
	"os"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/ewohltman/discordgo-mock/mockchannel"
	"github.com/ewohltman/discordgo-mock/mockconstants"
	"github.com/ewohltman/discordgo-mock/mockguild"
	"github.com/ewohltman/discordgo-mock/mockmember"
	"github.com/ewohltman/discordgo-mock/mockrest"
	"github.com/ewohltman/discordgo-mock/mockrole"
	"github.com/ewohltman/discordgo-mock/mocksession"
	"github.com/ewohltman/discordgo-mock/mockstate"
	"github.com/ewohltman/discordgo-mock/mockuser"
)

func TestGetToken(t *testing.T) {
	tests := []struct {
		App   string
		Token string
	}{
		{"dev", "testToken2"},
		{"", "testToken3"},
		{"aaaaaa", ""},
		{"bbbbbbbb", ""},
		{"", "cccccc"},
		{"", ""},
	}

	for _, test := range tests {
		os.Setenv("APP_ENV", test.App)
		os.Setenv("DISCORD_TOKEN", test.Token)

		token, err := getToken()

		if err != nil || token != test.Token {
			t.Error(err)
		}
	}
}

func TestNewDg(t *testing.T) {

	tests := []struct {
		Token  string
		Result string
		Exp    bool
	}{
		{"testing", "Bot testing", true},
		{"mockmock", "Bot testing", false},
	}

	for _, test := range tests {
		dg, err := newDg(test.Token)

		if err != nil || (dg.Token == test.Result) != test.Exp {
			t.Error(err)
		}
	}
}

func TestOpenDg(t *testing.T) {
	dg, err := mockDg()

	if err != nil {
		t.Error(err)
	}

	err = openDg(dg)

	if err.Error() != "HTTP 404 Not Found, 404 page not found\n" {
		t.Error(err)
	}
}

func mockDg() (*discordgo.Session, error) {
	state, err := newState()

	if err != nil {
		return nil, err
	}

	s, err := mocksession.New(
		mocksession.WithState(state),
		mocksession.WithClient(&http.Client{
			Transport: mockrest.NewTransport(state),
		}),
	)

	if err != nil {
		return nil, err
	}

	return s, err
}

func newState() (*discordgo.State, error) {
	role := mockrole.New(
		mockrole.WithID(mockconstants.TestRole),
		mockrole.WithName(mockconstants.TestRole),
		mockrole.WithPermissions(discordgo.PermissionViewChannel),
	)

	botUser := mockuser.New(
		mockuser.WithID(mockconstants.TestUser+"Bot"),
		mockuser.WithUsername(mockconstants.TestUser+"Bot"),
		mockuser.WithBotFlag(true),
	)

	botMember := mockmember.New(
		mockmember.WithUser(botUser),
		mockmember.WithGuildID(mockconstants.TestGuild),
		mockmember.WithRoles(role),
	)

	userMember := mockmember.New(
		mockmember.WithUser(mockuser.New(
			mockuser.WithID(mockconstants.TestUser),
			mockuser.WithUsername(mockconstants.TestUser),
		)),
		mockmember.WithGuildID(mockconstants.TestGuild),
		mockmember.WithRoles(role),
	)

	channel := mockchannel.New(
		mockchannel.WithID(mockconstants.TestChannel),
		mockchannel.WithGuildID(mockconstants.TestGuild),
		mockchannel.WithName(mockconstants.TestChannel),
		mockchannel.WithType(discordgo.ChannelTypeGuildVoice),
	)

	privateChannel := mockchannel.New(
		mockchannel.WithID(mockconstants.TestPrivateChannel),
		mockchannel.WithGuildID(mockconstants.TestGuild),
		mockchannel.WithName(mockconstants.TestPrivateChannel),
		mockchannel.WithType(discordgo.ChannelTypeGuildVoice),
		mockchannel.WithPermissionOverwrites(&discordgo.PermissionOverwrite{
			ID:   botMember.User.ID,
			Type: discordgo.PermissionOverwriteTypeMember,
			Deny: discordgo.PermissionViewChannel,
		}),
	)

	return mockstate.New(
		mockstate.WithUser(botUser),
		mockstate.WithGuilds(
			mockguild.New(
				mockguild.WithID(mockconstants.TestGuild),
				mockguild.WithName(mockconstants.TestGuild),
				mockguild.WithRoles(role),
				mockguild.WithChannels(channel, privateChannel),
				mockguild.WithMembers(botMember, userMember),
			),
		),
	)
}
