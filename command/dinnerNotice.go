package command

import (
	"SchoolDay/db"
	"SchoolDay/extension"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/volatiletech/null/v8"
)

func DinnerNotice(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	channelId := m.ChannelID
	discordId := m.Author.ID

	if len(args) < 2 {
		user, err := db.UserGet(discordId)

		if (err == nil && user.DinnerTime != null.String{}) {
			user.DinnerTime = null.String{}
			_, err = db.UserUpdate(discordId, user)

			if err != nil {
				log.Fatalln(err)

				extension.ChannelMessageSend(s, channelId, "알림을 중지하지 못했습니다.")
				return
			}
			extension.ChannelMessageSend(s, channelId, "더 이상 알림을 받지 않습니다.")
		} else {
			extension.ChannelMessageSend(s, channelId, "알림을 받고 계시지 않습니다.")
		}
		return
	}

	dinnerTime, err := time.Parse("0615:04", "01"+args[1])

	if err != nil {
		extension.ChannelMessageSend(s, channelId, "시각을 잘못 입력하셨습니다: `%s`", args[1])
		return
	}

	user, err := db.UserGet(discordId)

	if err != nil {
		extension.ChannelMessageSend(s, channelId, "학교 정보를 먼저 등록하세요.")
		return
	}

	user.DinnerTime = null.String{String: dinnerTime.Format("15:04"), Valid: true}
	_, err = db.UserUpdate(discordId, user)

	if err != nil {
		log.Fatalln(err)

		extension.ChannelMessageSend(s, channelId, "알림을 설정하지 못했습니다.")
		return
	}

	extension.ChannelMessageSend(s, channelId, "등록 완료: `%s`", dinnerTime.Format("15시04분"))
}
