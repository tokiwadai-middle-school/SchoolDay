package command

import (
	"SchoolDay/api"
	"SchoolDay/db"
	"SchoolDay/embed"
	"SchoolDay/extension"
	"time"

	"github.com/bwmarrin/discordgo"
)

func SchoolSchedule(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	channelId := m.ChannelID
	discordId := m.Author.ID

	var schoolInfo map[string]string
	var err error

	var date *time.Time = nil

	schoolName := ""

	for index, arg := range args {
		if index == 0 {
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

	embed, err := embed.SchoolScheduleEmbed(schoolInfo, *date)

	if err != nil {
		now, err := extension.NtpTimeKorea()
		format := "1월 2일"
		if err == nil && date.Year() != now.Year() {
			format = "2006년 " + format
		}

		extension.ChannelMessageSend(s, channelId, "%s 학사일정이 없습니다.", date.Format(format))
		return
	}

	extension.ChannelMessageSendEmbed(s, channelId, embed)
}
