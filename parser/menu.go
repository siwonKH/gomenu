package parser

import (
	"gomenu/handler"
	"gomenu/model"
	"strconv"
	"strings"
)

func stripNums(menuList []string) []string {
	for i := 0; i < len(menuList); i++ {
		strLen := len(menuList[i])
		if menuList[i][strLen-1] == '.' {
			menuList[i] = menuList[i][:strings.Index(menuList[i], ".")-1]
			menuList[i] = strings.Replace(menuList[i], "1", "", 1)
		}
	}
	return menuList
}

func ParseMenu(menuData model.NeisMenu) (model.Menu, error) {
	menuDataList := menuData.MealServiceDietInfo[1].Row

	var menuRes, blank model.Menu
	for i := 0; i < len(menuDataList); i++ {
		menuType, err := strconv.Atoi(menuDataList[i].MmealScCode) // MmealScCode 숫자로 변환
		if err != nil {
			return blank, handler.HandleErr("(menu)neis api parse failed")
		}
		menuRes[i].MenuList[menuType-1].MenuData = stripNums(strings.Split(menuDataList[i].DdishNm, "<br/>"))
	}

	return menuRes, nil
}
