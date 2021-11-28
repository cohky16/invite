package main_test

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ewohltman/discordgo-mock/mockchannel"
	"github.com/ewohltman/discordgo-mock/mockconstants"
	"github.com/ewohltman/discordgo-mock/mockmember"
	"github.com/ewohltman/discordgo-mock/mockrole"
	"github.com/ewohltman/discordgo-mock/mockuser"
)

type MockSession struct{}

func (ms MockSession) ChannelMessageSendEmbed(channelID string, embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return nil, nil
}

func (ms MockSession) GuildChannels(guildID string) (st []*discordgo.Channel, err error) {
	var channels []*discordgo.Channel

	channel := mockchannel.New(
		mockchannel.WithID(mockconstants.TestChannel),
		mockchannel.WithGuildID(mockconstants.TestGuild),
		mockchannel.WithName(mockconstants.TestChannel),
		mockchannel.WithType(discordgo.ChannelTypeGuildVoice),
	)

	channels = append(channels, channel)

	return channels, nil
}

func (ms MockSession) GuildMembers(guildID string, after string, limit int) (st []*discordgo.Member, err error) {
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

	var members []*discordgo.Member

	members = append(members, botMember)

	return members, nil
}

func (ms MockSession) Channel(channelID string) (st *discordgo.Channel, err error) {
	channel := mockchannel.New(
		mockchannel.WithID(mockconstants.TestChannel),
		mockchannel.WithGuildID(mockconstants.TestGuild),
		mockchannel.WithName(mockconstants.TestChannel),
		mockchannel.WithType(discordgo.ChannelTypeGuildVoice),
	)

	return channel, nil
}

func (ms MockSession) ChannelInviteCreate(channelID string, i discordgo.Invite) (st *discordgo.Invite, err error) {
	st = &discordgo.Invite{Code: "ok"}

	return
}
func (ms MockSession) ChannelMessageSend(channelID string, content string) (*discordgo.Message, error) {
	return nil, nil
}
