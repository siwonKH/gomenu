package gomenu

import (
	"github.com/siwonKH/gomenu/model"
	"github.com/siwonKH/gomenu/parser"
	"github.com/siwonKH/gomenu/request"
	"strconv"
)

func searchSchool(schoolName string, aptCode string, maxResp int, m *GoMenu) ([]model.School, error) {
	maxRes := strconv.Itoa(maxResp)
	var schoolRes, blank []model.School

	resp, err := request.SearchSchool(schoolName, aptCode, maxRes, m.KEY)
	if err != nil {
		return blank, err
	}
	schoolRes, err = parser.ParseSchool(resp)
	if err != nil {
		return blank, err
	}
	return schoolRes, nil
}

func (m *GoMenu) SearchSchool(schoolName string, aptCode string, maxResp int) ([]model.School, error) {
	return searchSchool(schoolName, aptCode, maxResp, m)
}

func (m *GoMenu) SearchFirstSchool(schoolName string) (model.School, error) {
	resp, err := searchSchool(schoolName, "", 1, m)
	return resp[0], err
}
