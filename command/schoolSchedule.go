package command

import (
	"SchoolDay/api"
	"SchoolDay/db"
	"SchoolDay/embed"
	"SchoolDay/extension"
	"strconv"
	"time"

	"github.com/beevik/ntp"
	"github.com/bwmarrin/discordgo"
)

func SchoolSchedule(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	var schoolInfo map[string]string
	var mealCode int
	var err error

	date, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Warningln(err)
		return
	}

	date = date.Add(time.Hour * 9)

	schoolName := ""

	for index, arg := range args {
		if index == 0 {
			continue
		}

		if extension.IsValidNumber(arg) {
			tempDate, err := time.Parse("20060102", strconv.Itoa(date.Year())+arg)

			if err == nil {
				date = tempDate
			}
		} else if len(schoolName) == 0 {
			schoolName = arg
		}
	}

	if len(schoolName) == 0 {
		user, err := db.UserGet(m.Author.ID)

		if err != nil {
			extension.ChannelMessageSend(s, m, "학교를 등록하지 않으셔서 학교 이름을 생략할 수 없습니다.")
			return
		}
		schoolInfo, _ = api.GetSchoolInfoByCode(user.ScCode)
	} else {
		schoolInfo, err = api.GetSchoolInfoByName(schoolName)

		if err != nil {
			extension.ChannelMessageSend(s, m, "학교를 찾을 수 없습니다: `%s`", schoolName)
			return
		}
	}

	embed, err := embed.DailySchoolScheduleEmbed(schoolInfo, date, mealCode)

	if err != nil {
		extension.ChannelMessageSend(s, m, "%d월 %d일 학사일정이 없습니다.", date.Month(), date.Day())
		return
	}

	extension.ChannelMessageSendEmbed(s, m, embed)
}
