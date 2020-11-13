package db

import (
	"fmt"
	"log"
	"../env"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
// discordId 			string
// scCode 				string
// scGrade 			string
// scClass 			string
// scheduleChannelId 	string
// timetableChannelId 	string
// dietChannelId 		string

type dbInfo struct {
	user		string
	pwd 		string
	url 		string
	engine 		string
	database 	string
}

var db = dbInfo{
	env.DBUser,
	env.DBPwd,
	env.DBUrl,
	env.DBEngine,
	env.DBName,
}
func dbCreate(name string) {
	source := db.user+":"+db.pwd+"@tcp("+db.url+")/"
	conn, err := sql.Open(db.engine, source)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Exec("CREATE DATABASE "+name)
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec("USE "+name)
	if err != nil {
		panic(err)
	}

	//CREATE TABLE user(
	//	discordId 			CHAR(18) PRIMARY KEY,
	//	scCode 				CHAR(7) NOT NULL,
	//	scGrade 			TINYINT(1),
	//	scClass 			TINYINT(1),
	//	scheduleChannelId 	CHAR(18),
	//	timetableChannelId 	CHAR(18),
	//	dietChannelId 		CHAR(18)
	//);
	query := "CREATE TABLE user (" +
		"discordId CHAR(18) PRIMARY KEY, " +
		"scCode CHAR(7) NOT NULL, " +
		"scGrade TINYINT(1), " +
		"scClass TINYINT(1), " +
		"scheduleChannelId CHAR(18), " +
		"timetableChannelId CHAR(18), " +
		"dietChannelId CHAR(18)" +
		");"
	_, err = conn.Exec(query)
	if err != nil {
		panic(err)
	}
}

func dbQuery(db dbInfo, query string) (count int) {
	dataSource := db.user+":"+db.pwd+"@tcp("+db.url+")/"+db.database
	conn, err := sql.Open(db.engine, dataSource)
	err = conn.QueryRow(query).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
	return count
}
