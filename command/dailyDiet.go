package command

import (
	"SchoolDay/api"
	"SchoolDay/db"
	"SchoolDay/embed"
	"SchoolDay/extension"
	"SchoolDay/models"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func DailyDiet(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	var schoolInfo map[string]string
	var mealCode int
	var err error

	schoolName := ""
	date := time.Now()

	for index, arg := range args {
		if index == 0 {
			continue
		}

		if extension.IsInt(arg) {
			tempDate, err := time.Parse("20060102", strconv.Itoa(time.Now().Year())+arg)

			if err == nil {
				date = tempDate
			}
		} else if strings.Contains(arg, "조식") || strings.Contains(arg, "아침") {
			mealCode = 1
		} else if strings.Contains(arg, "중식") || strings.Contains(arg, "점심") {
			mealCode = 2
		} else if strings.Contains(arg, "석식") || strings.Contains(arg, "저녁") {
			mealCode = 3
		} else if len(schoolName) == 0 {
			schoolName = arg
		}
	}

	if len(schoolName) == 0 {
		var user *models.User
		user, err = db.UserGet(m.Author.ID)

		if err != nil {
			_, err = s.ChannelMessageSend(m.ChannelID, "등록된 학교가 없습니다.")

			if err != nil {
				log.Warningln(err)
			}

			return
		}
		schoolInfo, err = api.GetSchoolInfoByCode(user.ScCode)
	} else {
		schoolInfo, err = api.GetSchoolInfoByName(schoolName)
	}

	if err != nil {
		log.Warningln(err)
		_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("학교를 찾을 수 없습니다: `%s`", schoolName))

		if err != nil {
			log.Warningln(err)
		}

		return
	}

	dailyDietEmbed, err := embed.GetDailyDietEmbed(schoolInfo, date, mealCode)

	if err != nil {
		mealName := "급식"

		switch mealCode {
			case 1:
				mealName = "조식"

			case 2:
				mealName = "중식"

			case 3:
				mealName = "석식"
		}

		_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%d월 %d일 %s이 없습니다.", date.Month(), date.Day(), mealName))

		if err != nil {
			log.Warningln(err)
		}

		return
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, dailyDietEmbed)

	if err != nil {
		log.Warningln(err)
	}
}
