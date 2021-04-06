package api

import (
	"SchoolDay/extension"
	"io/ioutil"
	"net/http"
	"net/url"
)

var log = extension.Log()

// API 호출
func Request(apiUrl string, params url.Values) ([]byte, error) {
	baseUrl, err := url.Parse(apiUrl) // string에서 URL로 변환

	if err != nil {
		return nil, err
	}

	params.Add("type", "json")         // json 타입으로 호출
	baseUrl.RawQuery = params.Encode() // 파라미터 삽입

	req, err := http.Get(baseUrl.String()) // GET으로 API 요청

	if err != nil {
		return nil, err
	}

	defer func() {
		err = req.Body.Close()
		if err != nil {
			log.Error(err)
		}
	}()
	body, err := ioutil.ReadAll(req.Body) // 바이트 배열로 저장

	if err != nil {
		return nil, err
	}

	return body, nil
}
