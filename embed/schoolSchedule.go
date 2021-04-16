package embed

import (
	"SchoolDay/api"
	"SchoolDay/extension"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// 하루 치 학사일정 embed
func DailySchoolScheduleEmbed(schoolInfo map[string]string, date time.Time, mealCode int) (*discordgo.MessageEmbed, error) {
	schoolSchedule, err := api.GetSchoolSchedule(schoolInfo, date, date)

	if err != nil {
		return nil, err
	}

	dailySchoolSchedule := schoolSchedule[date.Format("20060102")]

	var embed discordgo.MessageEmbed
	embed.Color = 0x43b581
	embed.Title = date.Format("1월 2일") + extension.GetKoreanWeekday(date)

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:  "학사일정",
		Value: strings.Join(dailySchoolSchedule, "\n"),
	})

	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: schoolInfo["SCHUL_NM"],
	}

	return &embed, nil
}
