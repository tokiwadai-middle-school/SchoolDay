package api

import (
	"../env"
	"encoding/json"
	"net/url"

	"github.com/buger/jsonparser"
)

// 학교 이름으로 학교 정보 검색
func GetSchoolInfoByName(SCHUL_NM string) (map[string]string, error) {
	return GetSchoolInfo("SCHUL_NM", SCHUL_NM)
}

// 학교 코드로 학교 정보 검색
func GetSchoolInfoByCode(SD_SCHUL_CODE string) (map[string]string, error) {
	return GetSchoolInfo("SD_SCHUL_CODE", SD_SCHUL_CODE)
}

// 학교 정보
func GetSchoolInfo(key string, value string) (map[string]string, error) {
	apiUrl := "https://open.neis.go.kr/hub/schoolInfo" // API 주소

	// API 파라미터
	params := url.Values{}
	params.Add("KEY", env.SchoolInfoKey)
	params.Add(key, value)

	jsonBytes, err := Request(apiUrl, params) // API 호출

	if err != nil {
		return nil, err
	}

	row, _, _, err := jsonparser.Get(jsonBytes, "schoolInfo", "[1]", "row", "[0]") // 검색 결과 중 가장 첫 번째만 파싱

	if err != nil {
		return nil, err
	}

	// Map으로 변환
	schoolInfo := make(map[string]string)
	json.Unmarshal(row, &schoolInfo)

	return schoolInfo, nil
}
