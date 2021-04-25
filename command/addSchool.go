package command

import (
	"SchoolDay/api"
	"SchoolDay/db"
	"SchoolDay/extension"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/volatiletech/null/v8"
)

func AddSchool(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	channelId := m.ChannelID
	discordId := m.Author.ID

	if len(args) == 1 {
		err := db.UserDelete(discordId)

		if err != nil {
			extension.ChannelMessageSend(s, channelId, "삭제할 학교 정보가 없습니다.")
		} else {
			extension.ChannelMessageSend(s, channelId, "등록된 학교 정보를 삭제했습니다.")
		}
		return
	} else if len(args) < 4 {
		extension.ChannelMessageSend(s, channelId, "사용법: `%s학교등록 [교명 학년 반]`", "%")
		return
	}

	schoolInfo, err := api.GetSchoolInfoByName(args[1])

	if err != nil {
		extension.ChannelMessageSend(s, channelId, "학교를 찾을 수 없습니다: `%s`", args[1])
		return
	}

	var grade, class int

	if extension.IsGradeNumber(args[2]) {
		grade, _ = strconv.Atoi(args[2])
	} else {
		extension.ChannelMessageSend(s, channelId, "학년을 잘못 입력하셨습니다: `%s`", args[2])
		return
	}

	if extension.IsClassNumber(args[3]) {
		class, _ = strconv.Atoi(args[3])
	} else {
		extension.ChannelMessageSend(s, channelId, "반을 잘못 입력하셨습니다: `%s`", args[3])
		return
	}

	scCode := schoolInfo["SD_SCHUL_CODE"]
	scGrade := null.Int8{Int8: int8(grade), Valid: true}
	scClass := null.Int8{Int8: int8(class), Valid: true}

	user, err := db.UserGet(discordId)
	var upsertErr error

	if err != nil {
		_, upsertErr = db.UserAdd(
			discordId,
			scCode,
			scGrade,
			scClass,
		)
	} else {
		user.ScCode = scCode
		user.ScGrade = scGrade
		user.ScClass = scClass

		_, upsertErr = db.UserUpdate(discordId, user)
	}

	if upsertErr != nil {
		log.Fatalln(err)

		extension.ChannelMessageSend(s, channelId, "학교 등록에 실패했습니다.")
		return
	}

	extension.ChannelMessageSend(s, channelId, "등록 완료: `%s %d학년%d반`", schoolInfo["SCHUL_NM"], scGrade.Int8, scClass.Int8)
}
