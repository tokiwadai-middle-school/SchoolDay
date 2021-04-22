package command

import (
	"SchoolDay/db"
	"SchoolDay/extension"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/volatiletech/null/v8"
)

func BreakfastNotice(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	channelId := m.ChannelID
	discordId := m.Author.ID

	if len(args) < 2 {
		user, err := db.UserGet(discordId)

		if (err == nil && user.BreakfastTime != null.String{}) {
			user.BreakfastTime = null.String{}
			_, err = db.UserUpdate(discordId, user)

			if err != nil {
				log.Fatalln(err)

				extension.ChannelMessageSend(s, channelId, "알림 삭제에 실패했습니다.")
				return
			}
			extension.ChannelMessageSend(s, channelId, "더 이상 조식 알림을 받지 않습니다.")
		} else {
			extension.ChannelMessageSend(s, channelId, "사용법: `%s조식알림 HH:MM`", "%")
		}
		return
	}

	breakfastTime, err := time.Parse("0615:04", "01"+args[1])

	if err != nil {
		extension.ChannelMessageSend(s, channelId, "시각 입력이 잘못됐습니다: `%s`", args[1])
		return
	}

	user, err := db.UserGet(discordId)

	if err != nil {
		extension.ChannelMessageSend(s, channelId, "학교 등록을 먼저 하세요.")
		return
	}

	user.BreakfastTime = null.String{String: breakfastTime.Format("15:04"), Valid: true}
	_, err = db.UserUpdate(discordId, user)

	if err != nil {
		log.Fatalln(err)

		extension.ChannelMessageSend(s, channelId, "알림 추가에 실패했습니다.")
		return
	}

	extension.ChannelMessageSend(s, channelId, "등록 완료: `%d시%d분`", breakfastTime.Hour(), breakfastTime.Minute())
}
