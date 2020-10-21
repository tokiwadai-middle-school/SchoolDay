package api

import (
	"SchoolDay/env"
	"encoding/json"
	"net/url"

	"github.com/buger/jsonparser"
)

// 학교 이름으로 학교 정보 검색
func GetSchoolInfo(schul_nm string) (map[string]string, error) {
	apiUrl := "https://open.neis.go.kr/hub/schoolInfo" // API 주소

	// API 파라미터
	params := url.Values{}
	params.Add("key", env.SchoolInfoKey)
	params.Add("schul_nm", schul_nm)

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
