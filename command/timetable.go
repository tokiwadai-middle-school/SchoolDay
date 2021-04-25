package command

import (
	"SchoolDay/api"
	"SchoolDay/db"
	"SchoolDay/embed"
	"SchoolDay/extension"
	"SchoolDay/models"
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Timetable(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	channelId := m.ChannelID
	discordId := m.Author.ID

	var schoolInfo map[string]string

	schoolName := ""
	grade, class := 0, 0

	date, err := extension.NtpTimeKorea()

	if err != nil {
		log.Warningln(err)
		return
	}

	for index, arg := range args {
		if index == 0 {
			continue
		}

		tempDate, err := time.Parse("200601/02", strconv.Itoa(date.Year())+arg)

		if err == nil {
			date = tempDate
		} else if extension.IsGradeNumber(arg) && grade == 0 {
			grade, _ = strconv.Atoi(arg)
		} else if extension.IsClassNumber(arg) && class == 0 {
			class, _ = strconv.Atoi(arg)
		} else if len(schoolName) == 0 {
			schoolName = arg
		}
	}

	if len(schoolName) == 0 || grade == 0 || class == 0 {
		var user *models.User
		user, err = db.UserGet(discordId)

		if err != nil {
			_, err = s.ChannelMessageSend(channelId, "학교 정보를 등록하지 않으셔서 교명과 학년, 반을 생략하실 수 없습니다.")

			if err != nil {
				log.Warningln(err)
			}

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
	} else {
		schoolInfo, err = api.GetSchoolInfoByName(schoolName)
	}

	if err != nil {
		log.Warningln(err)
		_, err := s.ChannelMessageSend(channelId, fmt.Sprintf("학교를 찾을 수 없습니다: `%s`", schoolName))

		if err != nil {
			log.Warningln(err)
		}

		return
	}

	embed, err := embed.TimetableEmbed(schoolInfo, date, grade, class)

	if err != nil {
		_, err := s.ChannelMessageSend(channelId, fmt.Sprintf("%d월 %d일 %d학년 %d반 수업이 없습니다.", date.Month(), date.Day(), grade, class))

		if err != nil {
			log.Warningln(err)
		}

		return
	}

	_, err = s.ChannelMessageSendEmbed(channelId, embed)

	if err != nil {
		log.Warningln(err)
	}
}
