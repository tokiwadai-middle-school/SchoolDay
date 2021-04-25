package main

import (
	"SchoolDay/api"
	"SchoolDay/command"
	"SchoolDay/db"
	"SchoolDay/embed"
	"SchoolDay/env"
	"SchoolDay/extension"
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/volatiletech/null/v8"
)

var Token string
var log = extension.Log()

// 봇 토큰 처리
func init() {
	flag.StringVar(&Token, "t", env.BotToken, "Bot Token")
	flag.Parse()
}

// 자동 알림
func NotifyEveryMinute(s *discordgo.Session) {
	date, err := extension.NtpTimeKorea()

	if err != nil {
		log.Warningln(err)
		return
	}

	time.Sleep(time.Duration(int(time.Second) * (60 - date.Second())))

	for c := time.Tick(1 * time.Minute); ; <-c {
		date, err := extension.NtpTimeKorea()

		if err != nil {
			log.Warningln(err)
			continue
		}

		users, err := db.UserGetAll("'%s' in (ScheduleTime, TimetableTime, BreakfastTime, LunchTime, DinnerTime)", date.Format("15:04"))

		if err != nil {
			continue
		}

		for _, user := range users {
			channel, err := s.UserChannelCreate(user.DiscordId)

			if err != nil {
				log.Warningln(err)
				continue
			}

			dateStr := null.String{String: date.Format("15:04"), Valid: true}
			schoolInfo, _ := api.GetSchoolInfoByCode(user.ScCode)

			if user.ScheduleTime == dateStr {
				embed, err := embed.SchoolScheduleEmbed(schoolInfo, date)

				if err == nil {
					extension.ChannelMessageSendEmbed(s, channel.ID, embed)
				}
			}

			if user.TimetableTime == dateStr {
				embed, err := embed.TimetableEmbed(schoolInfo, date, int(user.ScGrade.Int8), int(user.ScClass.Int8))

				if err == nil {
					extension.ChannelMessageSendEmbed(s, channel.ID, embed)
				}
			}

			if user.BreakfastTime == dateStr {
				embed, err := embed.MealServiceEmbed(schoolInfo, date, 1)

				if err == nil {
					extension.ChannelMessageSendEmbed(s, channel.ID, embed)
				}
			}

			if user.LunchTime == dateStr {
				embed, err := embed.MealServiceEmbed(schoolInfo, date, 2)

				if err == nil {
					extension.ChannelMessageSendEmbed(s, channel.ID, embed)
				}
			}

			if user.DinnerTime == dateStr {
				embed, err := embed.MealServiceEmbed(schoolInfo, date, 3)

				if err == nil {
					extension.ChannelMessageSendEmbed(s, channel.ID, embed)
				}
			}
		}
	}
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

	go NotifyEveryMinute(dg)
	dg.UpdateGameStatus(0, "%도움말")

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
