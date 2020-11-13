package api

import (
	"../env"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/buger/jsonparser"
)

// 학사 일정
func GetSchoolSchedule(schoolInfo map[string]string, fromDate time.Time, toDate time.Time) (map[string][]string, error) {
	apiUrl := "https://open.neis.go.kr/hub/SchoolSchedule" // API 주소

	// API 파라미터
	params := url.Values{}
	params.Add("KEY", env.SchoolScheduleKey)
	params.Add("ATPT_OFCDC_SC_CODE", schoolInfo["ATPT_OFCDC_SC_CODE"])
	params.Add("SD_SCHUL_CODE", schoolInfo["SD_SCHUL_CODE"])
	params.Add("AA_FROM_YMD", fromDate.Format("20060102"))
	params.Add("AA_TO_YMD", toDate.Format("20060102"))

	resultJson, err := Request(apiUrl, params) // API 호출

	if err != nil {
		return nil, err
	}

	schoolSchedule := map[string][]string{}
	rowCount, err := jsonparser.GetInt(resultJson, "SchoolSchedule", "[0]", "head", "[0]", "list_total_count") // 학사 일정 개수

	// 학사 일정이 없을 경우 에러 리턴
	if rowCount == 0 {
		return nil, errors.New("")
	}

	for index := 0; index < int(rowCount); index++ {
		rowJson, _, _, err := jsonparser.Get(resultJson, "SchoolSchedule", "[1]", "row", fmt.Sprintf("[%d]", index)) // 학사 일정

		if err != nil {
			return nil, err
		}

		// Map으로 변환
		row := make(map[string]string)
		json.Unmarshal(rowJson, &row)

		// 문자열 처리
		re := regexp.MustCompile(`[a-z A-Z 가-힣 0-9 & \s]`)
		event := strings.Join(re.FindAllString(row["EVENT_NM"], -1)[:], "")

		schoolSchedule[row["AA_YMD"]] = append(schoolSchedule[row["AA_YMD"]], event) // 해당하는 날짜에 학사 일정 삽입
	}

	return schoolSchedule, nil
}
