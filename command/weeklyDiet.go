package command

import (
	"SchoolDay/api"
	"SchoolDay/embed"
	"SchoolDay/extension"
	"time"

	"github.com/bwmarrin/discordgo"
)

func WeeklyDiet(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	schoolInfo, err := api.GetSchoolInfoByName(args[1])

	if err != nil {
		_, err = s.ChannelMessageSend(m.ChannelID, "학교를 찾을 수 없습니다.")
		extension.ErrorHandler(err)

	}

	weeklyDietEmbed, err := embed.GetWeeklyDietEmbed(schoolInfo, time.Now())

	if err != nil {
		_, err = s.ChannelMessageSend(m.ChannelID, "이번 주는 급식이 없습니다.")
		extension.ErrorHandler(err)
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, weeklyDietEmbed)
	extension.ErrorHandler(err)
}
