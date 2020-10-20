package command

import(
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Exec			func(*discordgo.Session, *discordgo.MessageCreate, []string) // 명령어 함수
	Description		string // 명령어 설명
	Usage			[]string // 명령어 사용법
}

var Commands = map[string]Command{
	"도움말": {
		Exec:			Help,
		Description:	"도움말 출력",
		Usage:			[]string{""},
	},
}