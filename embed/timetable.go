package embed

import (
	"SchoolDay/api"
	"SchoolDay/extension"
	"fmt"
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

	now, err := extension.NtpTimeKorea()

	var embed discordgo.MessageEmbed
	embed.Color = 0x43b581

	format := "1월 2일"
	if err == nil && date.Year() != now.Year() {
		format = "2006년 " + format
	}
	embed.Title = date.Format(format) + extension.GetKoreanWeekday(date)

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
		Text: fmt.Sprintf("%s %d-%d", schoolInfo["SCHUL_NM"], grade, class),
	}

	return &embed, nil
}
