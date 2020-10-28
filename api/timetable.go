package api

import (
	"SchoolDay/env"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/buger/jsonparser"
)

// 학급 시간표
func GetTimetable(schoolInfo map[string]string, grade string, class string, fromDate time.Time, toDate time.Time) (map[string]map[string]string, error) {
	var apiKind, timetableKey string // API 이름, API 키

	// 학교 종류 별 API 이름, API 키 설정
	switch schoolInfo["SCHUL_KND_SC_NM"] {
	case "초등학교":
		apiKind = "els"
		timetableKey = env.ElsTimetableKey

	case "중학교":
		apiKind = "mis"
		timetableKey = env.MisTimetableKey

	case "고등학교":
		apiKind = "his"
		timetableKey = env.HisTimetableKey
	}

	apiKind += "Timetable"
	apiUrl := "https://open.neis.go.kr/hub/" + apiKind // API 주소

	// API 파라미터
	params := url.Values{}
	params.Add("KEY", timetableKey)
	params.Add("ATPT_OFCDC_SC_CODE", schoolInfo["ATPT_OFCDC_SC_CODE"])
	params.Add("SD_SCHUL_CODE", schoolInfo["SD_SCHUL_CODE"])
	params.Add("GRADE", grade)
	params.Add("CLASS_NM", class)
	params.Add("TI_FROM_YMD", fromDate.Format("20060102"))
	params.Add("TI_TO_YMD", toDate.Format("20060102"))

	resultJson, err := Request(apiUrl, params) // API 호출

	if err != nil {
		return nil, err
	}

	timetable := map[string]map[string]string{}
	rowCount, err := jsonparser.GetInt(resultJson, apiKind, "[0]", "head", "[0]", "list_total_count") // 수업 개수

	// 수업이 없을 경우 에러 리턴
	if rowCount == 0 {
		return nil, errors.New("")
	}

	for index := 0; index < int(rowCount); index++ {
		rowJson, _, _, err := jsonparser.Get(resultJson, apiKind, "[1]", "row", fmt.Sprintf("[%d]", index)) // 수업

		if err != nil {
			return nil, err
		}

		// Map으로 변환
		row := make(map[string]string)
		json.Unmarshal(rowJson, &row)

		// 문자열 처리
		re := regexp.MustCompile(`[a-z A-Z 가-힣 0-9 & \s]`)
		content := strings.Join(re.FindAllString(row["ITRT_CNTNT"], -1)[:], "")

		// 날짜 별로 Map 생성
		_, exists := timetable[row["ALL_TI_YMD"]]

		if !exists {
			timetable[row["ALL_TI_YMD"]] = make(map[string]string)
		}

		timetable[row["ALL_TI_YMD"]][row["PERIO"]] = content // 해당하는 날짜에 수업 내용 삽입
	}

	return timetable, nil
}
