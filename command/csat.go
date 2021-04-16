package command

import (
	"SchoolDay/env"
	"SchoolDay/extension"
	"time"

	"github.com/beevik/ntp"
	"github.com/bwmarrin/discordgo"
)

func Csat(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	date, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Warningln(err)
		return
	}

	date = date.Add(time.Hour * 9)
	csatDate, err := time.Parse("20060102", env.CsatDate)

	if err != nil {
		log.Fatal(err)
		return
	}

	dDay := csatDate.Sub(date)

	if dDay.Microseconds() > 0 {
		extension.ChannelMessageSend(s, m, "%d 대학수학능력시험까지 **%d일 %d시간 %d분 %d초**", csatDate.Year(), int(dDay.Hours())/24, int(dDay.Hours())%24, int(dDay.Minutes())%60, int(dDay.Seconds())%60)
	} else {
		extension.ChannelMessageSend(s, m, "%d 대학수학능력시험 **D-Day**", csatDate.Year())
	}
}
