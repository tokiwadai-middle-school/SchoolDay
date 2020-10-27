package env

import (
	"fmt"
	"os"
	"strings"
)

var BotToken, SchoolInfoKey, SchoolScheduleKey, ElsTimetableKey, MisTimetableKey, HisTimetableKey, MealServiceDietInfoKey string

// 환경 변수 처리
func init() {
	exists := make(map[string]bool)

	// 환경 변수 값, 환경 변수 존재 여부
	BotToken, exists["BOT_TOKEN"] = os.LookupEnv("BOT_TOKEN")
	SchoolInfoKey, exists["SCHOOL_INFO_KEY"] = os.LookupEnv("SCHOOL_INFO_KEY")
	SchoolScheduleKey, exists["SCHOOL_SCHEDULE_KEY"] = os.LookupEnv("SCHOOL_SCHEDULE_KEY")
	ElsTimetableKey, exists["ELS_TIME_TABLE_KEY"] = os.LookupEnv("ELS_TIME_TABLE_KEY")
	MisTimetableKey, exists["MIS_TIME_TABLE_KEY"] = os.LookupEnv("MIS_TIME_TABLE_KEY")
	HisTimetableKey, exists["HIS_TIME_TABLE_KEY"] = os.LookupEnv("HIS_TIME_TABLE_KEY")
	MealServiceDietInfoKey, exists["MEAL_SERVICE_DIET_INFO_KEY"] = os.LookupEnv("MEAL_SERVICE_DIET_INFO_KEY")

	// 존재하지 않는 환경 변수 기록
	var missingEnv []string

	for key, value := range exists {
		if !value {
			missingEnv = append(missingEnv, key)
		}
	}

	// 존재하지 않는 환경 변수가 있을 경우 출력 후 종료
	if len(missingEnv) > 0 {
		fmt.Printf("set following environment variables: %s", strings.Join(missingEnv, "\n"))
		os.Exit(1)
	}
}
