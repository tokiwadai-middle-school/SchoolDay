package db

import (
	"SchoolDay/env"
	"SchoolDay/extension"
	"github.com/jinzhu/gorm" /* TODO: Document -> https://github.com/jirfag/go-queryset#relation-with-gorm */
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
// discordId 			string
// scCode 				string
// scGrade 				string
// scClass 				string
// scheduleChannelId 	string
// timetableChannelId 	string
// dietChannelId 		string

type dbInfo struct {
	user		string
	pwd 		string
	url 		string
	database 	string
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

	query := `CREATE TABLE user (
		discordId CHAR(18) PRIMARY KEY,
		scCode CHAR(7) NOT NULL,
		scGrade TINYINT,
		scClass TINYINT,
		scheduleChannelId CHAR(18),
		timetableChannelId CHAR(18),
		dietChannelId CHAR(18)
		);`

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

func GetGormDB() *gorm.DB {
	Source := dbInterface.user+":"+ dbInterface.pwd+"@tcp("+ dbInterface.url+")/"+ dbInterface.database
	conn, err := gorm.Open("mysql", Source)
	extension.ErrorHandler(err)

	/* TODO: DB가 이미 존재 한지 체크하도록 구현해야 함
	conn.Exec("CREATE DATABASE "+ db_interface.database)
	conn.Exec("USE "+db_interface.database)

	TODO: TABLE 이 존재하는지 체크하도록 구현해야 함
	query := `CREATE TABLE user (
		discordId CHAR(18) PRIMARY KEY,
		scCode CHAR(7) NOT NULL,
		scGrade TINYINT,
		scClass TINYINT,
		scheduleChannelId CHAR(18),
		timetableChannelId CHAR(18),
		dietChannelId CHAR(18)
		);`
	conn.Exec(query)
	*/
	return conn
}