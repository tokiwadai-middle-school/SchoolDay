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

// 학교 급식 식단
func GetMealServiceDietInfo(schoolInfo map[string]string, fromDate time.Time, toDate time.Time) (map[string]map[string]string, error) {
	apiUrl := "https://open.neis.go.kr/hub/mealServiceDietInfo" // API 주소

	// API 파라미터
	params := url.Values{}
	params.Add("KEY", env.MealServiceDietInfoKey)
	params.Add("ATPT_OFCDC_SC_CODE", schoolInfo["ATPT_OFCDC_SC_CODE"])
	params.Add("SD_SCHUL_CODE", schoolInfo["SD_SCHUL_CODE"])
	params.Add("MLSV_FROM_YMD", fromDate.Format("20060102"))
	params.Add("MLSV_TO_YMD", toDate.Format("20060102"))

	resultJson, err := Request(apiUrl, params) // API 호출

	if err != nil {
		return nil, err
	}

	mealServiceDietInfo := map[string]map[string]string{}
	rowCount, err := jsonparser.GetInt(resultJson, "mealServiceDietInfo", "[0]", "head", "[0]", "list_total_count") // 급식 개수

	// 급식이 없으면 에러 리턴
	if rowCount == 0 {
		return nil, errors.New("")
	}

	for index := 0; index < int(rowCount); index++ {
		rowJson, _, _, err := jsonparser.Get(resultJson, "mealServiceDietInfo", "[1]", "row", fmt.Sprintf("[%d]", index)) // 급식

		if err != nil {
			return nil, err
		}

		// Map으로 변환
		row := make(map[string]string)
		json.Unmarshal(rowJson, &row)

		// 문자열 처리
		re := regexp.MustCompile(`[a-z A-Z 가-힣 & \s]`)
		diet := strings.ReplaceAll(row["DDISH_NM"], "<br/>", "\n")
		diet = strings.Join(re.FindAllString(diet, -1)[:], "")

		// 날짜 별로 Map 생성
		_, exists := mealServiceDietInfo[row["MLSV_YMD"]]

		if !exists {
			mealServiceDietInfo[row["MLSV_YMD"]] = make(map[string]string)
		}

		mealServiceDietInfo[row["MLSV_YMD"]][row["MMEAL_SC_CODE"]] = diet // 해당하는 날짜에 급식 식단 삽입
	}

	return mealServiceDietInfo, nil
}
