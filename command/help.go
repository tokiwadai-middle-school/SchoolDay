package command

import (
	"SchoolDay/embed"
	"SchoolDay/extension"

	"github.com/bwmarrin/discordgo"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	channelId := m.ChannelID

	embed, err := embed.HelpEmbed(args)

	if err != nil {
		extension.ChannelMessageSend(s, channelId, "알 수 없는 명령어입니다: `%s`", args[1])
	} else {
		extension.ChannelMessageSendEmbed(s, channelId, embed)
	}
}
