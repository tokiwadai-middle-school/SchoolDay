package command

import (
	"SchoolDay/env"
	"SchoolDay/extension"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Csat(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	channelId := m.ChannelID
	date, err := extension.NtpTimeKorea()

	if err != nil {
		log.Warningln(err)
		return
	}

	csatDate, err := time.Parse("20060102", env.CsatDate)

	if err != nil {
		log.Fatal(err)
		return
	}

	dDay := csatDate.Sub(date)

	if dDay.Microseconds() > 0 {
		extension.ChannelMessageSend(s, channelId, "%d 대학수학능력시험까지 **%d일 %d시간 %d분 %d초**", csatDate.Year(), int(dDay.Hours())/24, int(dDay.Hours())%24, int(dDay.Minutes())%60, int(dDay.Seconds())%60)
	} else {
		extension.ChannelMessageSend(s, channelId, "%d 대학수학능력시험 **D-Day**", csatDate.Year())
	}
}
