package command

import (
	"SchoolDay/api"
	"SchoolDay/embed"
	"github.com/bwmarrin/discordgo"
	"time"
)

func WeeklyDiet(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		log.Warning("학교가 지정되지 않았습니다.")
		_, err := s.ChannelMessageSend(m.ChannelID, "학교가 지정되지 않았습니다.")
		if err != nil {
			log.Warning(err)
		}
	} else {
		schoolInfo, err := api.GetSchoolInfoByName(args[1])

		if err != nil {
			_, err = s.ChannelMessageSend(m.ChannelID, "학교를 찾을 수 없습니다.")
			if err != nil {
				log.Warning(err)
			}
		}

		weeklyDietEmbed, err := embed.GetWeeklyDietEmbed(schoolInfo, time.Now())

		if err != nil {
			_, err = s.ChannelMessageSend(m.ChannelID, "이번 주는 급식이 없습니다.")
			if err != nil {
				log.Warning(err)
			}
		}

		_, err = s.ChannelMessageSendEmbed(m.ChannelID, weeklyDietEmbed)
		if err != nil {
			log.Warning(err)
		}
	}
}
