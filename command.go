package main

import (
	"github.com/bwmarrin/discordgo"
)

func onCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Data.Name {
	case "help":
		onHelpForCommand(s, i)
	case "invite":
		onInviteForCommand(s, i)
	default:
		break
	}
}

func onHelpForCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	footer := discordgo.MessageEmbedFooter{Text: "🍎 ご要望、不具合は https://github.com/cohky16/invite までお願いします"}

	embed := discordgo.MessageEmbed{
		Title: "機能概要",
		Description: "ボイスチャンネルへの招待を送信できます\n\n" +
			"**__各種コマンド__**\n\n" +
			"**invite**\nユーザーにボイスチャンネルへの招待情報を送信します\n`/invite to: @hoge @fuga`\n\n" +
			"**invite $(RoomNo)**\nユーザーに特定のボイスチャンネルへの招待情報を送信します\n`/invite to: @hoge @fuga channel: Piyo`\n\n" +
			"**help**\nヘルプを表示します",
		Footer: &footer,
	}

	embeds := []*discordgo.MessageEmbed{
		&embed,
	}

	session := newSession(s, i.Member.User.Bot)

	if err := session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Embeds: embeds,
		},
	}); err != nil {
		return
	}

	return
}

func onInviteForCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channelId := makeChannelId(s, i)

	if channelId == "" {
		return
	}

	st, err := s.ChannelInviteCreate(channelId, discordgo.Invite{})

	if st == nil || err != nil {
		return
	}

	session := newSession(s, i.Member.User.Bot)

	if err := session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: i.Data.Options[0].StringValue() + "\nボイスチャンネルに招待されました\n" + "https://discord.gg/" + st.Code,
		},
	}); err != nil {
		return
	}

	return
}

func makeChannelId(s *discordgo.Session, i *discordgo.InteractionCreate) string {
	if len(i.Data.Options) == 2 {
		return i.Data.Options[1].StringValue()
	}

	session := newSession(s, i.Member.User.Bot)

	channels, _ := session.GuildChannels(i.GuildID)
	state, _ := s.State.VoiceState(i.GuildID, i.Member.User.ID)
	senderChannelName := makeSenderChannelName(s, i)
	alreadyChannelIds := makeAlreadyChannelIds(s, i)

	for _, channel := range channels {
		if state != nil && state.ChannelID == channel.ID {
			return channel.ID
		} else if state == nil && channel.Type == 2 && checkRegexp(senderChannelName, channel.Name) {
			count := 0

			for _, alreadyChannelID := range alreadyChannelIds {
				if alreadyChannelID == channel.ID {
					count++
				}
			}

			if count < 2 {
				return channel.ID
			}
		}
	}

	return ""
}

func makeSenderChannelName(s *discordgo.Session, i *discordgo.InteractionCreate) (senderChannelName string) {
	session := newSession(s, i.Member.User.Bot)

	senderChannel, _ := session.Channel(i.ChannelID)

	if checkRegexp("talk", senderChannel.Name) {
		senderChannelName = "Talk"
	} else if checkRegexp("meeting", senderChannel.Name) {
		senderChannelName = "Meeting"
	}

	return
}

func makeAlreadyChannelIds(s *discordgo.Session, i *discordgo.InteractionCreate) (alreadyChannelIds []string) {
	members, _ := s.GuildMembers(i.GuildID, "", 1000)
	alreadyChannelIds = []string{}

	for _, member := range members {
		state, err := s.State.VoiceState(i.GuildID, member.User.ID)

		if state == nil || err != nil {
			continue
		}

		alreadyChannelIds = append(alreadyChannelIds, state.ChannelID)
	}

	return
}
