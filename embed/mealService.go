package embed

import (
	"SchoolDay/api"
	"SchoolDay/extension"
	"sort"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func MealServiceEmbed(schoolInfo map[string]string, date time.Time, mealCode int) (*discordgo.MessageEmbed, error) {
	mealServiceDietInfo, err := api.GetMealServiceDietInfo(schoolInfo, date, date, mealCode)

	if err != nil {
		return nil, err
	}

	dailyMealService := mealServiceDietInfo[date.Format("20060102")]

	now, err := extension.NtpTimeKorea()

	var embed discordgo.MessageEmbed
	embed.Color = 0x43b581

	format := "1월 2일"
	if err == nil && date.Year() != now.Year() {
		format = "2006년 " + format
	}
	embed.Title = date.Format(format) + extension.GetKoreanWeekday(date)

	// 급식 정렬

	keys := make([]int, len(dailyMealService))
	index := 0

	for key := range dailyMealService {
		keys[index] = key
		index++
	}

	sort.Ints(keys)

	for _, key := range keys {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   extension.GetMealName(key),
			Value:  strings.Join(dailyMealService[key], "\n"),
			Inline: true,
		})
	}

	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: schoolInfo["SCHUL_NM"],
	}

	return &embed, nil
}
