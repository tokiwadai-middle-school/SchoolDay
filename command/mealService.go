package command

import (
	"SchoolDay/api"
	"SchoolDay/db"
	"SchoolDay/embed"
	"SchoolDay/extension"
	"time"

	"github.com/bwmarrin/discordgo"
)

func MealService(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	channelId := m.ChannelID
	discordId := m.Author.ID

	var schoolInfo map[string]string
	var err error

	var date *time.Time = nil

	mealCode := 0
	schoolName := ""

	for index, arg := range args {
		if index == 0 {
			continue
		}

		tempMealCode := extension.GetMealCode(arg)

		if tempMealCode != 0 {
			if mealCode == 0 {
				mealCode = tempMealCode
			}
			continue
		}

		tempDate, err := extension.ParseDate(arg)

		if err == nil {
			if date == nil {
				date = tempDate
			}
			continue
		}

		if len(schoolName) == 0 {
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

	if date == nil {
		tempDate, err := extension.NtpTimeKorea()
		if err != nil {
			log.Warningln(err)
			return
		}
		date = &tempDate
	}

	embed, err := embed.MealServiceEmbed(schoolInfo, *date, mealCode)

	if err != nil {
		mealName := extension.GetMealName(mealCode)
		extension.ChannelMessageSend(s, channelId, "%d월 %d일 %s이 없습니다.", date.Month(), date.Day(), mealName)
		return
	}

	extension.ChannelMessageSendEmbed(s, channelId, embed)
}
