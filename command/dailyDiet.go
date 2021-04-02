package command

import (
	"SchoolDay/api"
	"SchoolDay/embed"
	"SchoolDay/extension"
	"time"

	"github.com/bwmarrin/discordgo"
)

func DailyDiet(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	schoolInfo, err := api.GetSchoolInfoByName(args[1])

	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, "학교를 찾을 수 없습니다.")
		extension.ErrorHandler(err)
		return
	}

	dailyDietEmbed, err := embed.GetDailyDietEmbed(schoolInfo, time.Now())

	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, "오늘은 급식이 없습니다.")
		extension.ErrorHandler(err)
		return
	}

	_, err  = s.ChannelMessageSendEmbed(m.ChannelID, dailyDietEmbed)
	extension.ErrorHandler(err)
}
