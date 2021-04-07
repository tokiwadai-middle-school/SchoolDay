package command

import (
	"SchoolDay/api"
	"SchoolDay/db"
	"SchoolDay/extension"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/volatiletech/null/v8"
)

// discordId 			string
// scCode 				string
// scGrade 				string
// scClass 				string
// scheduleChannelId 	string
// timetableChannelId 	string
// dietChannelId

func AddSchool(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 4 {
		_, err = s.ChannelMessageSend(m.ChannelID, "사용법: `%학교등록 학교명 학년 반`")

		if err != nil {
			log.Warningln(err)
		}

		return
	} else {
		var discordId = m.Author.ID
		scCode, err := api.GetSchoolInfoByName(args[1])

		if err != nil {
			log.Warningln(err)

			_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("학교를 찾을 수 없습니다: `%s`", args[1]))

			if err != nil {
				log.Warningln(err)
			}

			return
		}

		var ScGrade, ScClass int

		if extension.IsInt(args[2]) {
			ScGrade, _ = strconv.Atoi(args[2])
		} else {
			_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("학년 입력이 잘못 됐습니다: `%s`", args[2]))

			if err != nil {
				log.Warningln(err)
			}

			return
		}

		if extension.IsInt(args[3]) {
			ScClass, _ = strconv.Atoi(args[3])
		} else {
			_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("반 입력이 잘못 됐습니다: `%s`", args[3]))

			if err != nil {
				log.Warningln(err)
			}

			return
		}

		var AcGrade int8
		AcGrade = int8(ScGrade)
		scGrade := null.Int8From(AcGrade)

		var BcClass int8
		BcClass = int8(ScClass)
		scClass := null.Int8From(BcClass)

		scheduleChannelId := null.String{}
		timetableChannelId := null.String{}

		_, err = db.UserAdd(discordId, scCode["SD_SCHUL_CODE"], scGrade, scClass, scheduleChannelId, timetableChannelId)

		if err != nil {
			log.Fatalln(err)

			_, err = s.ChannelMessageSend(m.ChannelID, "학교 등록에 실패했습니다.")

			if err != nil {
				log.Warningln(err)
			}

			return
		}

		_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("등록 완료: `%s %d-%d`", scCode["SCHUL_NM"], scGrade.Int8, scClass.Int8))

		if err != nil {
			log.Warningln(err)
			return
		}
	}
}
