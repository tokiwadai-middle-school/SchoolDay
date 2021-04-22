package command

import (
	"SchoolDay/api"
	"SchoolDay/db"
	"SchoolDay/embed"
	"SchoolDay/extension"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func SchoolSchedule(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	channelId := m.ChannelID
	discordId := m.Author.ID

	var schoolInfo map[string]string
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
		user, err := db.UserGet(discordId)

		if err != nil {
			extension.ChannelMessageSend(s, channelId, "학교를 등록하지 않으셔서 학교 이름을 생략할 수 없습니다.")
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

	embed, err := embed.SchoolScheduleEmbed(schoolInfo, date)

	if err != nil {
		extension.ChannelMessageSend(s, channelId, "%d월 %d일 학사일정이 없습니다.", date.Month(), date.Day())
		return
	}

	extension.ChannelMessageSendEmbed(s, channelId, embed)
}
