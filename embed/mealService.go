package embed

import (
	"SchoolDay/api"
	"SchoolDay/extension"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// 하루 급식 식단
func GetDailyDietEmbed(schoolInfo map[string]string, date time.Time, mealCode int) (*discordgo.MessageEmbed, error) {
	mealServiceDietInfo, err := api.GetMealServiceDietInfo(schoolInfo, date, date, mealCode)

	if err != nil {
		return nil, err
	}

	dailyDiet := mealServiceDietInfo[date.Format("20060102")] // 하루 치 급식 식단

	var embed discordgo.MessageEmbed
	embed.Color = 0x43b581
	embed.Title = date.Format("1월 2일") + extension.GetKoreanWeekday(date)

	// 각 급식 별 급식 시간대 코드와 식단
	for mealCode, diet := range dailyDiet {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  extension.GetMealName(mealCode),
			Value: "```" + strings.Join(diet, "\n") + "```",
		})
	}

	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: schoolInfo["SCHUL_NM"],
	}

	return &embed, nil
}
