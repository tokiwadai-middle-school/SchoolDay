package env

import (
	"fmt"
	"os"
	"strings"
)

var (
	SchoolInfoKey          string
	SchoolScheduleKey      string
	ElsTimetableKey        string
	MisTimetableKey        string
	HisTimetableKey        string
	MealServiceDietInfoKey string
	BotToken               string
	DBUser				   string // MySQL Username
	DBPwd				   string // MySQL Password
	DBUrl				   string // MySQL Server URL
	DBEngine 			   string // MySQL DB Engine
	DBName 				   string // MySQL DATABASE Name
)

// 환경 변수 처리
func init() {
	exists := make(map[string]bool)

	// 환경 변수 값, 환경 변수 존재 여부
	// API Credentials
	BotToken, exists["BOT_TOKEN"] = os.LookupEnv("BOT_TOKEN")
	SchoolInfoKey, exists["SCHOOL_INFO_KEY"] = os.LookupEnv("SCHOOL_INFO_KEY")
	SchoolScheduleKey, exists["SCHOOL_SCHEDULE_KEY"] = os.LookupEnv("SCHOOL_SCHEDULE_KEY")
	ElsTimetableKey, exists["ELS_TIME_TABLE_KEY"] = os.LookupEnv("ELS_TIME_TABLE_KEY")
	MisTimetableKey, exists["MIS_TIME_TABLE_KEY"] = os.LookupEnv("MIS_TIME_TABLE_KEY")
	HisTimetableKey, exists["HIS_TIME_TABLE_KEY"] = os.LookupEnv("HIS_TIME_TABLE_KEY")
	MealServiceDietInfoKey, exists["MEAL_SERVICE_DIET_INFO_KEY"] = os.LookupEnv("MEAL_SERVICE_DIET_INFO_KEY")
	// MySQL Credentials
	DBUser, exists["MySQL_CREDENTIAL_USERNAME"] = os.LookupEnv("MySQL_CREDENTIAL_USERNAME")
	DBPwd, exists["MySQL_CREDENTIAL_PASSWORD"] = os.LookupEnv("MySQL_CREDENTIAL_PASSWORD")
	DBUrl, exists["MySQL_SERVER_URL"] = os.LookupEnv("MySQL_SERVER_URL")
	DBEngine, exists["MySQL_DB_ENGINE"] = os.LookupEnv("MySQL_DB_ENGINE")
	DBName, exists["MySQL_DB_NAME"] = os.LookupEnv("MySQL_DB_NAME")

	// 존재하지 않는 환경 변수 기록
	var missingEnv []string

	for key, value := range exists {
		if !value {
			missingEnv = append(missingEnv, key)
		}
	}

	// 존재하지 않는 환경 변수가 있을 경우 출력 후 종료
	if len(missingEnv) > 0 {
		fmt.Printf("\nmissing environment variables:\n%s\n\n", strings.Join(missingEnv, "\n"))
		os.Exit(1)
	}
}
