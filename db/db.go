package db

import (
	"SchoolDay/env"
	"SchoolDay/extension"
	"SchoolDay/models"
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

/* TODO: ORM 구현 완료하면 삭제해야 함
func dbCreate(name string) {
	source := db.user+":"+db.pwd+"@tcp("+db.url+")/"
	conn, err := sql.Open(db.engine, source)
	extension.ErrorHandler(err)

	defer func() {
		err = conn.Close()
		extension.ErrorHandler(err)
	}()

	_, err = conn.Exec("CREATE DATABASE "+name)
	extension.ErrorHandler(err)

	_, err = conn.Exec("USE "+name)
	extension.ErrorHandler(err)

	query := `CREATE TABLE user ( discordId CHAR(18) PRIMARY KEY, scCode CHAR(7) NOT NULL, scGrade TINYINT, scClass TINYINT, scheduleChannelId CHAR(18), timetableChannelId CHAR(18), dietChannelId CHAR(18));`

	_, err = conn.Exec(query)
	extension.ErrorHandler(err)
}
*/
/*TODO: ORM 구현 완료하면 삭제해야 함
func dbQuery(db dbInfo, query string) (count int) {
	dataSource := db.user+":"+db.pwd+"@tcp("+db.url+")/"+db.database

	conn, err := sql.Open(db.engine, dataSource)
	extension.ErrorHandler(err)

	err = conn.QueryRow(query).Scan(&count)
	extension.ErrorHandler(err)

	fmt.Println(count)
	return count
}
*/

// discordId 			string
// scCode 				string
// scGrade 				string
// scClass 				string
// scheduleChannelId 	string
// timetableChannelId 	string
// dietChannelId 		string
/*
type User struct {
	DiscordId          string
	ScCode             string `gorm:"not null"`
	ScGrade            string
	ScClass            string
	scheduleChannelId  string
	timetableChannelId string
	dietChannelId      string
}
*/

var schema = `
CREATE TABLE user ( 
	discordId CHAR(18) PRIMARY KEY, 
	scCode CHAR(7) NOT NULL, 
	scGrade TINYINT, 
	scClass TINYINT, 
	scheduleChannelId CHAR(18), 
	timetableChannelId CHAR(18), 
	dietChannelId CHAR(18)
);
`

func Database() (*sqlx.DB, error) {
	dsn := dbInterface.user + ":" + dbInterface.pwd + "@tcp(" + dbInterface.url + ")/" + dbInterface.database + "?charset=utf8mb4"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}
	//db.MustExec(schema)
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
	status, err := models.UserExists(ctx, db, discordId)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func UserAdd(
	discordId string,
	scCode string,
	scGrade null.Int8,
	scClass null.Int8,
	scheduleChannelId null.String,
	timetableChannelId null.String,
	dietChannelId null.String) (interface{}, error) {

	status, err := IsExists(discordId)
	if err != nil {
		return nil, err
	}
	if true != status {
		db, err := Database()

		if err != nil {
			return nil, err
		}
		ctx := context.Background()
		resp := models.User{
			DiscordId:          discordId,
			ScCode:             scCode,
			ScGrade:            scGrade,
			ScClass:            scClass,
			ScheduleChannelId:  scheduleChannelId,
			TimetableChannelId: timetableChannelId,
			DietChannelId: 		dietChannelId,
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

	ctx := context.Background()

	return models.FindUser(ctx, db, discordId)
}
