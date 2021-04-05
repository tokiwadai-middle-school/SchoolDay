package command

import (
	"SchoolDay/extension"
	"github.com/bwmarrin/discordgo"
)

var log = extension.Log()

type Command struct {
	Exec        func(*discordgo.Session, *discordgo.MessageCreate, []string) // 명령어 함수
	Description string                                                       // 명령어 설명
	Usage       []string                                                     // 명령어 사용법
}

var Commands = map[string]Command{
	"도움말": {
		Exec:        Help,
		Description: "도움말 출력",
		Usage:       []string{""},
	},
	"급식": {
		Exec:        NextDiet,
		Description: "다음 급식 식단 출력",
		Usage:       []string{"학교명"},
	},
	"일간급식": {
		Exec:        DailyDiet,
		Description: "오늘 급식 식단 출력",
		Usage:       []string{"학교명"},
	},
	"주간급식": {
		Exec:        WeeklyDiet,
		Description: "이번 주 급식 식단 출력",
		Usage:       []string{"학교명"},
	},
}
