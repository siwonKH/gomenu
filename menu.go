package gomenu

import (
	"github.com/siwonKH/gomenu/model"
	"github.com/siwonKH/gomenu/parser"
	"github.com/siwonKH/gomenu/request"
	"time"
)

func searchMenu(schoolCode string, aptCode string, date string, date2 string) (model.Menu, error) {
	var menuRes, blank model.Menu
	resp, err := request.SearchMenu(schoolCode, aptCode, date, date2)
	if err != nil {
		return blank, err
	}
	menuRes, err = parser.ParseMenu(resp)
	if err != nil {
		return blank, err
	}
	return menuRes, nil
}

func (m *m) SearchMenu(schoolCode string, aptCode string, date string) (model.Menu, error) {
	return searchMenu(schoolCode, aptCode, date, m)
}

func SearchMenuByRange(schoolCode string, aptCode string, date string, date2 string) (model.Menu, error) {
	return searchMenu(schoolCode, aptCode, date, date2)
}

func SearchTodayMenu(schoolCode string, aptCode string) (model.Menu, error) {
	today := time.Now().Format("20060102")
	return searchMenu(schoolCode, aptCode, today, today)
}

func SearchNextDayMenu(schoolCode string, aptCode string) (model.Menu, error) {
	tomorrow := time.Now().AddDate(0, 0, 1).Format("20060102")
	return searchMenu(schoolCode, aptCode, tomorrow, tomorrow)
}

func SearchThisMonthMenu(schoolCode string, aptCode string) (model.Menu, error) {
	today := time.Now()
	startMonth := time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, time.Local)
	endMonth := startMonth.AddDate(0, 1, 0).AddDate(0, 0, -1)
	return searchMenu(schoolCode, aptCode, startMonth.Format("20060102"), endMonth.Format("20060102"))
}

func SearchNextMonthMenu(schoolCode string, aptCode string) (model.Menu, error) {
	today := time.Now().AddDate(0, 1, 0)
	startMonth := time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, time.Local)
	endMonth := startMonth.AddDate(0, 1, 0).AddDate(0, 0, -1)
	return searchMenu(schoolCode, aptCode, startMonth.Format("20060102"), endMonth.Format("20060102"))
}
