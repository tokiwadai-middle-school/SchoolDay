package command

import (
	"SchoolDay/api"
	"SchoolDay/db"
	"SchoolDay/embed"
	"SchoolDay/extension"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func MealService(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	channelId := m.ChannelID
	discordId := m.Author.ID

	var schoolInfo map[string]string
	var mealCode int
	var err error

	date, err := extension.NtpTimeKorea()

	if err != nil {
		log.Warningln(err)
		return
	}

	schoolName := ""

	for index, arg := range args {
		if index == 0 {
			continue
		}

		tempDate, err := time.Parse("200601/02", strconv.Itoa(date.Year())+arg)

		if err == nil {
			date = tempDate
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
		user, err := db.UserGet(discordId)

		if err != nil {
			extension.ChannelMessageSend(s, channelId, "학교 정보를 등록하지 않으셔서 교명을 생략하실 수 없습니다.")
			return
		}
		schoolInfo, _ = api.GetSchoolInfoByCode(user.ScCode)
	} else {
		schoolInfo, err = api.GetSchoolInfoByName(schoolName)

		if err != nil {
			extension.ChannelMessageSend(s, channelId, "학교를 찾을 수 없습니다: `%s`", schoolName)
			return
		}
	}

	embed, err := embed.MealServiceEmbed(schoolInfo, date, mealCode)

	if err != nil {
		mealName := extension.GetMealName(mealCode)
		extension.ChannelMessageSend(s, channelId, "%d월 %d일 %s이 없습니다.", date.Month(), date.Day(), mealName)
		return
	}

	extension.ChannelMessageSendEmbed(s, channelId, embed)
}
