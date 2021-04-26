package db

import (
	"SchoolDay/env"
	"SchoolDay/extension"
	"SchoolDay/models"
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var log = extension.Log()

type dbInfo struct {
	user     string
	pwd      string
	url      string
	database string
}

var dbInterface = dbInfo{
	env.DBUser,
	env.DBPwd,
	env.DBUrl,
	env.DBName,
}

//go:generate sqlboiler
type User struct {
	DiscordId     string
	ScCode        string
	ScGrade       null.Int8
	ScClass       null.Int8
	ScheduleTime  null.String
	TimetableTime null.String
	BreakfastTime null.String
	LunchTime     null.String
	DinnerTime    null.String
}

var schema = `
CREATE TABLE user ( 
	discordId		CHAR(18)	PRIMARY KEY, 
	scCode 			CHAR(7)		NOT NULL, 
	scGrade			TINYINT,
	scClass			TINYINT,
	ScheduleTime 	CHAR(5),
	TimetableTime   CHAR(5),
	BreakfastTime	CHAR(5),
	LunchTime       CHAR(5),
	DinnerTime      CHAR(5)
);
`

func Database() (*sqlx.DB, error) {
	dsn := dbInterface.user + ":" + dbInterface.pwd + "@tcp(" + dbInterface.url + ")/" + dbInterface.database + "?charset=utf8mb4"

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	boil.SetDB(db)
	return db, nil
}

func IsExists(discordId string) (interface{}, error) {
	ctx := context.Background()

	db, err := Database()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer db.Close()

	status, err := models.UserExists(ctx, db, discordId)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func UserAdd(discordId string, scCode string, scGrade null.Int8, scClass null.Int8) (interface{}, error) {
	status, err := IsExists(discordId)
	if err != nil {
		return nil, err
	}

	if true != status {
		db, err := Database()
		if err != nil {
			return nil, err
		}
		defer db.Close()

		ctx := context.Background()
		resp := models.User{
			DiscordId:     discordId,
			ScCode:        scCode,
			ScGrade:       scGrade,
			ScClass:       scClass,
			ScheduleTime:  null.String{},
			TimetableTime: null.String{},
			BreakfastTime: null.String{},
			LunchTime:     null.String{},
			DinnerTime:    null.String{},
		}
		err = resp.Insert(ctx, db, boil.Infer())
		if err != nil {
			return nil, err
		}
		return true, nil
	}
	return false, nil
}

func UserGet(discordId string) (*models.User, error) {
	db, err := Database()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	ctx := context.Background()
	return models.FindUser(ctx, db, discordId)
}

func UserGetAll(format string, args ...interface{}) (models.UserSlice, error) {
	db, err := Database()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	ctx := context.Background()
	whereClause := fmt.Sprintf(format, args...)
	return models.Users(qm.Where(whereClause)).All(ctx, db)
}

func UserUpdate(discordId string, user *models.User) (bool, error) {
	ctx := context.Background()

	db, err := Database()
	if err != nil {
		return false, err
	}
	defer db.Close()

	userResult, err := UserGet(discordId)
	if err != nil {
		return false, err
	}

	userResult.ScCode = user.ScCode
	userResult.ScGrade = user.ScGrade
	userResult.ScClass = user.ScClass
	userResult.ScheduleTime = user.ScheduleTime
	userResult.TimetableTime = user.TimetableTime
	userResult.BreakfastTime = user.BreakfastTime
	userResult.LunchTime = user.LunchTime
	userResult.DinnerTime = user.DinnerTime

	_, err = userResult.Update(ctx, db, boil.Infer())

	if err != nil {
		return false, err
	}
	return true, nil
}

func UserDelete(discordId string) error {
	ctx := context.Background()

	db, err := Database()
	if err != nil {
		return err
	}
	defer db.Close()

	user, err := UserGet(discordId)

	if err != nil {
		return err
	}

	_, err = user.Delete(ctx, db)
	return err
}
