package command

import (
	"SchoolDay/api"
	"SchoolDay/db"
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
// dietChannelId 		string

func AddSchool(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if err != nil {
		log.Fatalln(err)
		return
	}

	if len(args) < 2 {

	} else {
		var discordId = m.Author.ID
		scCode, err := api.GetSchoolInfoByName(args[1])

		if err != nil {
			log.Error(err)
		}

		ScGrade, err := strconv.Atoi(args[2])

		if err != nil {
			log.Fatalln(err)
			return
		}

		ScClass, err := strconv.Atoi(args[3])

		if err != nil {
			log.Fatalln(err)
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
			return
		}

		_, err = s.ChannelMessageSend(m.ChannelID, "ok")

		if err != nil {
			log.Error(err)
			return
		}
		return
	}

}
