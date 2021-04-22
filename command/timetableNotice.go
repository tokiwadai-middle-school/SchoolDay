package command

import (
	"SchoolDay/db"
	"SchoolDay/extension"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/volatiletech/null/v8"
)

func TimetableNotice(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	channelId := m.ChannelID
	discordId := m.Author.ID

	if len(args) < 2 {
		user, err := db.UserGet(discordId)

		if (err == nil && user.TimetableTime != null.String{}) {
			user.TimetableTime = null.String{}
			_, err = db.UserUpdate(discordId, user)

			if err != nil {
				log.Fatalln(err)

				extension.ChannelMessageSend(s, channelId, "알림 삭제에 실패했습니다.")
				return
			}
			extension.ChannelMessageSend(s, channelId, "더 이상 시간표 알림을 받지 않습니다.")
		} else {
			extension.ChannelMessageSend(s, channelId, "사용법: `%s시간표알림 HH:MM`", "%")
		}
		return
	}

	timetableTime, err := time.Parse("0615:04", "01"+args[1])

	if err != nil {
		extension.ChannelMessageSend(s, channelId, "시각 입력이 잘못됐습니다: `%s`", args[1])
		return
	}

	user, err := db.UserGet(discordId)

	if err != nil {
		extension.ChannelMessageSend(s, channelId, "학교 등록을 먼저 하세요.")
		return
	}

	user.TimetableTime = null.String{String: timetableTime.Format("15:04"), Valid: true}
	_, err = db.UserUpdate(discordId, user)

	if err != nil {
		log.Fatalln(err)

		extension.ChannelMessageSend(s, channelId, "알림 추가에 실패했습니다.")
		return
	}

	extension.ChannelMessageSend(s, channelId, "등록 완료: `%d시%d분`", timetableTime.Hour(), timetableTime.Minute())
}
