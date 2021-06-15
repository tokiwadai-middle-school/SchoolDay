package embed

import (
	"SchoolDay/api"
	"SchoolDay/extension"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func SchoolScheduleEmbed(schoolInfo map[string]string, date time.Time) (*discordgo.MessageEmbed, error) {
	schoolSchedule, err := api.GetSchoolSchedule(schoolInfo, date, date)

	if err != nil {
		return nil, err
	}

	dailySchoolSchedule := schoolSchedule[date.Format("20060102")]

	now, err := extension.NtpTimeKorea()

	var embed discordgo.MessageEmbed
	embed.Color = 0x43b581

	format := "1월 2일"
	if err == nil && date.Year() != now.Year() {
		format = "2006년 " + format
	}
	embed.Title = date.Format(format) + extension.GetKoreanWeekday(date)

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:  "학사일정",
		Value: strings.Join(dailySchoolSchedule, "\n"),
	})

	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: schoolInfo["SCHUL_NM"],
	}

	return &embed, nil
}
