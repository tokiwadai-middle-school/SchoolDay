package command

import (
	"SchoolDay/api"
	"SchoolDay/embed"
	"time"

	"github.com/bwmarrin/discordgo"
)

func NextDiet(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	schoolInfo, err := api.GetSchoolInfoByName(args[1])

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "학교를 찾을 수 없습니다.")
		return
	}

	nextDietEmbed, err := embed.GetNextDietEmbed(schoolInfo, time.Now())

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "금일 급식이 없습니다.")
		return
	}

	s.ChannelMessageSendEmbed(m.ChannelID, nextDietEmbed)
}
