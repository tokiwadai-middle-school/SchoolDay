package extension

import (
	"github.com/mattn/go-colorable"
	logHandler "github.com/sirupsen/logrus"
	"time"
)

// 날짜에 해당하는 한국어 요일 반환
func GetKoreanWeekday(date time.Time) string {
	var koreanWeekday string

	switch int(date.Weekday()) {
	case 0:
		koreanWeekday = "일"
	case 1:
		koreanWeekday = "월"
	case 2:
		koreanWeekday = "화"
	case 3:
		koreanWeekday = "수"
	case 4:
		koreanWeekday = "목"
	case 5:
		koreanWeekday = "금"
	case 6:
		koreanWeekday = "토"
	}

	return "(" + koreanWeekday + ")"
}

// 날짜가 속하는 주의 순서 반환
func GetWeekNumber(date time.Time) int {
	firstDay := time.Date(date.Year(), date.Month(), 1, 1, 1, 1, 1, time.UTC)

	_, currentWeek := date.ISOWeek()
	_, firstWeek := firstDay.ISOWeek()

	return currentWeek - firstWeek
}

// 시각에 해당하는 급식 시간대 코드 반환
func GetMealCode(date time.Time) int {
	if date.Hour() < 8 {
		return 1 // 조식
	} else if date.Hour() < 13 {
		return 2 // 중식
	} else {
		return 3 // 석식
	}
}

// 급식 시간대 코드에 해당하는 이름 반환
func GetMealName(mealCode int) string {
	var mealName string

	switch mealCode {
	case 1:
		mealName = "조식"

	case 2:
		mealName = "중식"

	case 3:
		mealName = "석식"
	}

	return mealName
}


func Log() *logHandler.Entry {
	logHandler.SetFormatter(&logHandler.TextFormatter{
		ForceColors: true,
	})
	logHandler.SetOutput(colorable.NewColorableStdout())
	logHandler.SetLevel(logHandler.DebugLevel)
	var lo = logHandler.WithFields(logHandler.Fields{})
	return lo
}
