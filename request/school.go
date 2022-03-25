package request

import (
	"encoding/json"
	"github.com/siwonKH/gomenu/handler"
	"github.com/siwonKH/gomenu/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func SearchSchool(schoolName string, aptCode string, maxRes string, KEY string) (model.NeisSchool, error) {
	var schoolData, blank model.NeisSchool
	schoolNameUrl := url.QueryEscape(schoolName)
	urlIn := "https://open.neis.go.kr/hub/schoolInfo?Type=json&pSize=" + maxRes
	urlIn += "&SCHUL_NM=" + schoolNameUrl + "&ATPT_OFCDC_SC_CODE=" + aptCode + "&KEY=" + KEY

	resp, err := http.Get(urlIn)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(resp.Body)

	if err != nil {
		return blank, handler.HandleErr("(school)request failed: " + schoolName)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return blank, handler.HandleErr("(school)response read failed: " + schoolName)
	}
	err = json.Unmarshal([]byte(data), &schoolData)
	if err != nil {
		return blank, handler.HandleErr("(school)json unmarshal failed: " + schoolName)
	}

	if schoolData.SchoolInfo[0].Head[0].ListTotalCount < 1 {
		return blank, handler.HandleErr("(school)no school found: " + schoolName)
	}
	return schoolData, nil
}
