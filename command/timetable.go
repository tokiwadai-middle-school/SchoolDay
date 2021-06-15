package command

import (
	"SchoolDay/api"
	"SchoolDay/db"
	"SchoolDay/embed"
	"SchoolDay/extension"
	"SchoolDay/models"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Timetable(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	channelId := m.ChannelID
	discordId := m.Author.ID

	var schoolInfo map[string]string

	var date *time.Time = nil

	schoolName := ""
	grade, class := 0, 0

	for index, arg := range args {
		if index == 0 {
			continue
		}

		num, err := strconv.Atoi(arg)

		if err == nil {
			if grade == 0 && num <= 6 && num >= 1 {
				grade = num
			} else if class == 0 && num <= 16 && num >= 1 {
				class = num
			}
		} else {
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
	}

	if len(schoolName) == 0 || grade == 0 || class == 0 {
		var user *models.User
		user, err = db.UserGet(discordId)

		if err != nil {
			extension.ChannelMessageSend(s, channelId, "학교 정보를 등록하지 않으셔서 교명과 학년, 반을 생략하실 수 없습니다.")
			return
		}

		if len(schoolName) == 0 {
			schoolInfo, err = api.GetSchoolInfoByCode(user.ScCode)
		} else {
			schoolInfo, err = api.GetSchoolInfoByName(schoolName)
		}

		if grade == 0 {
			grade = int(user.ScGrade.Int8)
		}

		if class == 0 {
			class = int(user.ScClass.Int8)
		}
	}

	if err != nil {
		log.Warningln(err)
		extension.ChannelMessageSend(s, channelId, "학교를 찾을 수 없습니다: `%s`", schoolName)
		return
	}

	if date == nil {
		tempDate, err := extension.NtpTimeKorea()
		if err != nil {
			log.Warningln(err)
			return
		}
		date = &tempDate
	}

	embed, err := embed.TimetableEmbed(schoolInfo, *date, grade, class)

	if err != nil {
		now, err := extension.NtpTimeKorea()
		format := "1월 2일"
		if err == nil && date.Year() != now.Year() {
			format = "2006년 " + format
		}

		extension.ChannelMessageSend(s, channelId, "%s %d학년 %d반 수업이 없습니다.", date.Format(format), grade, class)
	}

	extension.ChannelMessageSendEmbed(s, channelId, embed)
}
