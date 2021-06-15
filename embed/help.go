package embed

import (
	"SchoolDay/extension"
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HelpEmbed(args []string) (*discordgo.MessageEmbed, error) {
	var embed discordgo.MessageEmbed

	embed.Color = 0x43b581
	embed.Title = "SchoolDay 도움말"

	if len(args) >= 2 {
		var fieldName, fieldValue string

		switch args[1] {
		case "학교등록":
			fieldName = "%학교등록 [교명 학년 반]"
			fieldValue = "자신의 학교와 학년, 반을 등록합니다.\n등록함으로써 앞으로 명령어를 쓸 때\n교명과 학년, 반을 생략할 수 있습니다.\n등록한 정보를 삭제하려면 `%학교등록`을 입력하세요.\n(주의: 삭제 시 설정해 두신 알림까지 모두 중지됩니다)"

		case "학사일정":
			fieldName = "%학사일정 [교명] [날짜]"
			fieldValue = "당일 학사일정을 알려줍니다."

		case "시간표":
			fieldName = "%시간표 [교명] [학년] [반] [날짜]"
			fieldValue = "당일 시간표를 알려줍니다."

		case "급식":
			fieldName = "%급식 [교명] [날짜] [급식 종류]"
			fieldValue = "당일 급식 식단을 알려줍니다.\n급식 종류는 `아침`, `점심`, `저녁`이나\n`조식`, `중식`, `석식`이라고 입력하세요.\n급식 종류를 생략하면 당일 급식을 모두 알려줍니다."

		case "수능":
			fieldName = "%수능"
			fieldValue = "다음 대학수학능력시험까지 남은 기간을 알려줍니다."

		case "학사일정알림", "시간표알림", "조식알림", "중식알림", "석식알림":
			noticeType := args[1][:len(args[1])-len("알림")]

			fieldName = "%" + args[1] + " [시각]"
			fieldValue = fmt.Sprintf("봇이 매일 지정한 시각에 %s 알려줍니다.\n알림을 중지하려면 `%s`을 입력하세요.\n학교 정보를 먼저 등록하셔야 사용하실 수 있습니다.", extension.AddKoreanObjectParticle(noticeType), "%"+args[1])

		default:
			return nil, errors.New("")
		}

		if strings.Contains(fieldName, "[") {
			embed.Description += "\n대괄호 쳐진 부분은 생략 가능합니다."
		}

		if strings.Contains(fieldName, "교명") {
			embed.Description += "\n교명이 길면 일부만 쓰셔도 됩니다."
		}

		if strings.Contains(fieldName, "시각") {
			embed.Description += "\n시각은 18:06과 같은 형식으로 입력하세요."
		}

		if strings.Contains(fieldName, "날짜") {
			embed.Description += "\n날짜는 12/06 또는 2020/12/06과 같은 형식이나\n`그제`, `어제`, `오늘`, `내일`, `모레`라고 입력하세요.\n날짜를 생략하면 오늘로 처리됩니다."
		}

		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  fieldName,
			Value: fieldValue,
		})
	} else {
		embed.Description = "`%도움말 [명령어]`로 자세한 설명을 보실 수 있습니다."

		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  "명령어",
			Value: "학교등록\n학사일정\n시간표\n급식\n학사일정알림\n시간표알림\n조식알림\n중식알림\n석식알림\n수능",
		})
	}

	return &embed, nil
}
