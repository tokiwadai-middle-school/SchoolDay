package command

import (
	"SchoolDay/api"
	"SchoolDay/db"
	"SchoolDay/extension"
	"github.com/bwmarrin/discordgo"
	"github.com/volatiletech/null/v8"
	"strconv"
)

// discordId 			string
// scCode 				string
// scGrade 				string
// scClass 				string
// scheduleChannelId 	string
// timetableChannelId 	string
// dietChannelId		string

func AddSchool(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 4 {
		extension.ChannelMessageSend(s, m, "사용법: %s 학교등록 학교명 학년 반", "%")
		return
	}

	var discordId = m.Author.ID
	schoolInfo, err := api.GetSchoolInfoByName(args[1])

	if err != nil {
		extension.ChannelMessageSend(s, m, "학교를 찾을 수 없습니다: `%s`", args[1])
		return
	}

	var grade, class int

	if extension.IsValidNumber(args[2]) {
		argNum, _ := strconv.Atoi(args[2])

		if argNum > 6 {
			extension.ChannelMessageSend(s, m, "학년이 너무 높습니다: `%s`", args[2])
			return
		}

		grade = argNum
	} else {
		extension.ChannelMessageSend(s, m, "학년 입력이 잘못 됐습니다: `%s`", args[2])
		return
	}

	if extension.IsValidNumber(args[3]) {
		argNum, _ := strconv.Atoi(args[3])

		if argNum > 16 {
			extension.ChannelMessageSend(s, m, "반이 너무 큽니다: `%s`", args[3])
			return
		}

		class = argNum
	} else {
		extension.ChannelMessageSend(s, m, "반 입력이 잘못 됐습니다: `%s`", args[3])
		return
	}

	var AcGrade int8
	AcGrade = int8(grade)
	scGrade := null.Int8From(AcGrade)

	var BcClass int8
	BcClass = int8(class)
	scClass := null.Int8From(BcClass)

	scheduleTime  := null.Time{}
	timetableTime := null.Time{}
	breakfastTime := null.Time{}
	lunchTime     := null.Time{}
	dinnerTime    := null.Time{}

	_, err = db.UserAdd(
		discordId,
		schoolInfo["SD_SCHUL_CODE"],
		scGrade,
		scClass,
		scheduleTime,
		timetableTime,
		breakfastTime,
		lunchTime,
		dinnerTime,
		)

	if err != nil {
		log.Fatalln(err)

		extension.ChannelMessageSend(s, m, "학교 등록에 실패했습니다.")
		return
	}

	extension.ChannelMessageSend(s, m, "등록 완료: `%s %d학년 %d반`", schoolInfo["SCHUL_NM"], scGrade.Int8, scClass.Int8)
}
