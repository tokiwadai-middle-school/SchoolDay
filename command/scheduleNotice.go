package command

import (
	"SchoolDay/db"
	"SchoolDay/extension"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/volatiletech/null/v8"
)

func scheduleNotice(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		extension.ChannelMessageSend(s, m, "사용법: %s학사일정알림 HH:MM", "%")
		return
	}

	scheduleTime, err := time.Parse("0615:04", "01"+args[1])

	if err != nil {
		extension.ChannelMessageSend(s, m, "시각 입력이 잘못됐습니다: `%s`", args[1])
		return
	}

	discordId := m.Author.ID
	user, err := db.UserGet(discordId)

	fmt.Println(err.Error())

	if err != nil {
		extension.ChannelMessageSend(s, m, "학교 등록을 먼저 하셔야됩니다.")
		return
	}

	user.ScheduleTime = null.Time{Time: scheduleTime, Valid: true}
	_, err = db.UserUpdate(discordId, user)

	if err != nil {
		log.Fatalln(err)

		extension.ChannelMessageSend(s, m, "알림 등록에 실패했습니다.")
		return
	}

	extension.ChannelMessageSend(s, m, "등록 완료: `%d시%d분`", scheduleTime.Hour(), scheduleTime.Minute())
}
