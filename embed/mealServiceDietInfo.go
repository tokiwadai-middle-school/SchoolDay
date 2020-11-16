package mealServiceDietInfo

import (
	"strconv"
	"strings"
	"time"

	"../../api"
	"../../extension"


	"github.com/bwmarrin/discordgo"
)

// 다음 급식 식단
func GetNextDietEmbed(schoolInfo map[string]string, date time.Time) (*discordgo.MessageEmbed, error) {
	mealServiceDietInfo, err := api.GetMealServiceDietInfo(schoolInfo, date, date)

	if err != nil {
		return nil, err
	}

	mealCode := extension.GetMealCode(date)                   // 급식 시간대 코드
	dailyDiet := mealServiceDietInfo[date.Format("20060102")] // 하루 치 급식 식단
	nextDiet, exists := dailyDiet[mealCode]                   // 시간대에 해당하는 급식 식단

	// 시간대에 해당하는 급식 식단이 없을 경우 다른 시간대 급식 식단 탐색
	for !exists {
		if extension.GetMealCode(date) == 1 { // 조식 시간대일 경우 뒤로 가며 탐색
			mealCode++
		} else if extension.GetMealCode(date) == 3 { // 석식 시간대일 경우 앞으로 가며 탐색
			mealCode--
		} else { // 중식 시간대일 경우 뒤를 먼저, 앞을 나중에 탐색
			if mealCode == 2 {
				mealCode = 3
			} else {
				mealCode = 1
			}
		}

		nextDiet, exists = dailyDiet[mealCode]
	}

	var embed discordgo.MessageEmbed
	embed.Color = 0x43b581
	embed.Title = date.Format("1월 2일") + extension.GetKoreanWeekday(date)
	embed.Fields = []*discordgo.MessageEmbedField{
		&discordgo.MessageEmbedField{
			Name:  extension.GetMealName(mealCode),
			Value: "```" + strings.Join(nextDiet, "\n") + "```",
		},
	}
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: schoolInfo["SCHUL_NM"],
	}

	return &embed, nil
}

// 일간 급식 식단
func GetDailyDietEmbed(schoolInfo map[string]string, date time.Time) (*discordgo.MessageEmbed, error) {
	mealServiceDietInfo, err := api.GetMealServiceDietInfo(schoolInfo, date, date)

	if err != nil {
		return nil, err
	}

	dailyDiet := mealServiceDietInfo[date.Format("20060102")] // 하루 치 급식 식단

	var embed discordgo.MessageEmbed
	embed.Color = 0x43b581
	embed.Title = date.Format("1월 2일") + extension.GetKoreanWeekday(date)

	for mealCode, diet := range dailyDiet { // 각 급식 별 급식 시간대 코드와 식단
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

// 주간 급식 식단
func GetWeeklyDietEmbed(schoolInfo map[string]string, date time.Time) (*discordgo.MessageEmbed, error) {
	weekdayNumber := int(date.Weekday())               // 요일 순서에 해당하는 숫자
	monday := date.AddDate(0, 0, -(weekdayNumber - 1)) // 월요일
	friday := monday.AddDate(0, 0, 4)                  // 금요일

	mealServiceDietInfo, err := api.GetMealServiceDietInfo(schoolInfo, monday, friday) // 일주일 치 급식

	if err != nil {
		return nil, err
	}

	weekNumber := extension.GetWeekNumber(date) // N주차 계산

	var embed discordgo.MessageEmbed
	embed.Color = 0x43b581
	embed.Title = date.Format("1월 ") + strconv.Itoa(weekNumber) + "주차"

	for date := monday; date.Unix() <= friday.Unix(); date = date.AddDate(0, 0, 1) { // 월요일부터 금요일까지 하루씩
		dailyDiet := mealServiceDietInfo[date.Format("20060102")] // 하루치 급식 식단
		var maxLine int

		// 가장 많은 메뉴(=가장 많은 new line) 수
		for mealCode := 1; mealCode <= 3; mealCode++ {
			if len(dailyDiet[mealCode]) > maxLine {
				maxLine = len(dailyDiet[mealCode])
			}
		}

		for mealCode := 1; mealCode <= 3; mealCode++ {
			fieldName := "\u200B"

			// 맨 앞 칸에만 일자 표시
			if mealCode == 1 {
				fieldName = date.Format("2일") + extension.GetKoreanWeekday(date)
			}

			var fieldValue string

			// 최대 메뉴 수에 맞춰 new line 채워놓기
			for line := 0; line < maxLine; line++ {
				fieldValue += "\t\t\t\n"
			}

			_, exists := dailyDiet[mealCode]

			if exists {
				fieldValue = strings.Join(dailyDiet[mealCode], "\n")

				// 최대 메뉴 수에 못 미칠 경우 그 만큼 new line으로 채우기
				for line := len(dailyDiet[mealCode]); line < maxLine; line++ {
					fieldValue += "\n\t"
				}
			}

			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   fieldName,
				Value:  extension.GetMealName(mealCode) + "\n```" + fieldValue + "```",
				Inline: true,
			})
		}
	}

	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: schoolInfo["SCHUL_NM"],
	}

	return &embed, nil
}
