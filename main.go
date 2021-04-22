package main

import (
	"SchoolDay/command"
	"SchoolDay/env"
	"SchoolDay/extension"
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var Token string
var log = extension.Log()

// 봇 토큰 처리
func init() {
	flag.StringVar(&Token, "t", env.BotToken, "Bot Token")
	flag.Parse()
}

// 봇 연결
func main() {
	dg, err := discordgo.New("Bot " + Token)

	if err != nil {

		log.Fatal("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
		return
	}

	log.Infof("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	err = dg.Close()
	if err != nil {
		log.Fatal("error closing listening/heartbeat goroutine", err)
		return
	}

}

// 메시지 핸들러
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	prefix := "%"

	// 접두사 감지 시 명령어 검색
	if strings.HasPrefix(m.Content, prefix) {
		args := strings.Fields(m.Content[len(prefix):])
		cmd, exists := command.Commands[args[0]]

		if !exists {
			return
		}

		cmd.Exec(s, m, args)
	}
}
