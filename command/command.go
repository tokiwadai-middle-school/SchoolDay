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
	"학사일정알림": {
		Exec:        ScheduleNotice,
		Description: "학사일정 자동 알림",
		Usage:       []string{"HH:MM"},
	},
	"시간표알림": {
		Exec:        TimetableNotice,
		Description: "시간표 자동 알림",
		Usage:       []string{"HH:MM"},
	},
	"조식알림": {
		Exec:        BreakfastNotice,
		Description: "조식 자동 알림",
		Usage:       []string{"HH:MM"},
	},
	"중식알림": {
		Exec:        LunchNotice,
		Description: "중식 자동 알림",
		Usage:       []string{"HH:MM"},
	},
	"석식알림": {
		Exec:        DinnerNotice,
		Description: "석식 자동 알림",
		Usage:       []string{"HH:MM"},
	},
	"수능": {
		Exec:        Csat,
		Description: "수능 D-Day 카운트",
		Usage:       []string{""},
	},
}
