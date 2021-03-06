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
	footer := discordgo.MessageEmbedFooter{Text: "ð ãè¦æãä¸å·åã¯ https://github.com/cohky16/invite ã¾ã§ãé¡ããã¾ã"}

	embed := discordgo.MessageEmbed{
		Title: "æ©è½æ¦è¦",
		Description: "ãã¤ã¹ãã£ã³ãã«ã¸ã®æå¾ãéä¿¡ã§ãã¾ã\n\n" +
			"**__åç¨®ã³ãã³ã__**\n\n" +
			"**invite**\nã¦ã¼ã¶ã¼ã«ãã¤ã¹ãã£ã³ãã«ã¸ã®æå¾æå ±ãéä¿¡ãã¾ã\n`/invite to: @hoge @fuga`\n\n" +
			"**invite $(RoomNo)**\nã¦ã¼ã¶ã¼ã«ç¹å®ã®ãã¤ã¹ãã£ã³ãã«ã¸ã®æå¾æå ±ãéä¿¡ãã¾ã\n`/invite to: @hoge @fuga channel: Piyo`\n\n" +
			"**help**\nãã«ããè¡¨ç¤ºãã¾ã",
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
			Content: i.Data.Options[0].StringValue() + "\nãã¤ã¹ãã£ã³ãã«ã«æå¾ããã¾ãã\n" + "https://discord.gg/" + st.Code,
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
