package embed

import (
	"SchoolDay/api"
	"SchoolDay/extension"
	"sort"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func TimetableEmbed(schoolInfo map[string]string, date time.Time, grade int, class int) (*discordgo.MessageEmbed, error) {
	timetable, err := api.GetTimetable(schoolInfo, date, date, grade, class)

	if err != nil {
		return nil, err
	}

	dailyTimetable := timetable[date.Format("20060102")]

	var embed discordgo.MessageEmbed
	embed.Color = 0x43b581
	embed.Title = date.Format("1월 2일") + extension.GetKoreanWeekday(date)

	// 시간표 정렬

	keys := make([]int, len(dailyTimetable))
	index := 0

	for key := range dailyTimetable {
		keys[index] = key
		index++
	}

	sort.Ints(keys)

	for _, key := range keys {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  strconv.Itoa(key) + "교시",
			Value: dailyTimetable[key],
		})
	}

	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: schoolInfo["SCHUL_NM"],
	}

	return &embed, nil
}
