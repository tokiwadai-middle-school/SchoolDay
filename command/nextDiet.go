package command

import (
	"SchoolDay/api"
	"SchoolDay/embed"
	"github.com/bwmarrin/discordgo"
	"time"
)


func NextDiet(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	schoolInfo, err := api.GetSchoolInfoByName(args[1])

	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, "학교를 찾을 수 없습니다.")
		if err != nil {
			log.Warning(err)
		}
	}

	nextDietEmbed, err := embed.GetNextDietEmbed(schoolInfo, time.Now())

	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, "금일 급식이 없습니다.")
		if err != nil {
			log.Warning(err)
		}
		return
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, nextDietEmbed)
	if err != nil {
		log.Warning(err)
	}
}
