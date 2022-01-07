package request

import (
	"encoding/json"
	"fmt"
	"github.com/siwonKH/gomenu/handler"
	"github.com/siwonKH/gomenu/model"
	"io/ioutil"
	"net/http"
)

func SearchMenu(schoolCode string, aptCode string, date string, date2 string, KEY string) (model.NeisMenu, error) {
	urlIn := "https://open.neis.go.kr/hub/mealServiceDietInfo?Type=json"
	urlIn += "&KEY=" + KEY + "&ATPT_OFCDC_SC_CODE=" + aptCode + "&SD_SCHUL_CODE=" + schoolCode + "&MLSV_FROM_YMD=" + date + "&MLSV_TO_YMD=" + date2

	var menuData, blank model.NeisMenu
	resp, err := http.Get(urlIn)
	defer resp.Body.Close()
	if err != nil {
		return blank, handler.HandleErr("(menu)request failed: " + aptCode + schoolCode)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return blank, handler.HandleErr("(menu)read failed: " + aptCode + schoolCode)
	}
	err = json.Unmarshal([]byte(data), &menuData)
	if err != nil {
		return blank, handler.HandleErr("(menu)json unmarshal failed: " + aptCode + schoolCode)
	}
	if menuData.MealServiceDietInfo == nil {
		return blank, handler.HandleErr("(menu)menu not found: " + aptCode + schoolCode)
	}
	if menuData.MealServiceDietInfo[0].Head[0].ListTotalCount < 1 {
		return blank, fmt.Errorf("no school menu")
	}

	return menuData, nil
}
