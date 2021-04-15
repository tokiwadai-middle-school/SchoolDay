package command

import (
	"SchoolDay/db"
	"SchoolDay/extension"

	"github.com/bwmarrin/discordgo"
)

var log = extension.Log()
var conn, err = db.Database()

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
	"학교등록": {
		Exec:        AddSchool,
		Description: "학교와 학급 등록",
		Usage:       []string{"학교명", "학년", "반"},
	},
	"학사일정": {
		Exec:        SchoolSchedule,
		Description: "학사일정 출력",
		Usage:       []string{"학교명", "날짜", "학년", "반"},
	},
	"시간표": {
		Exec:        Timetable,
		Description: "시간표 출력",
		Usage:       []string{"학교명", "날짜", "학년", "반"},
	},
	"급식": {
		Exec:        MealService,
		Description: "급식 식단 출력",
		Usage:       []string{"학교명", "날짜", "식사종류"},
	},
}
