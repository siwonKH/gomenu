package gomenu

import (
	"github.com/siwonKH/gomenu/model"
	"github.com/siwonKH/gomenu/parser"
	"github.com/siwonKH/gomenu/request"
	"strconv"
)

func searchSchool(schoolName string, aptCode string, maxResp int) ([]model.School, error) {
	maxRes := strconv.Itoa(maxResp)
	var schoolRes, blank []model.School

	resp, err := request.SearchSchool(schoolName, aptCode, maxRes)
	if err != nil {
		return blank, err
	}
	schoolRes, err = parser.ParseSchool(resp)
	if err != nil {
		return blank, err
	}
	return schoolRes, nil
}

func SearchSchool(schoolName string, aptCode string, maxResp int) ([]model.School, error) {
	return searchSchool(schoolName, aptCode, maxResp)
}

func SearchFirstSchool(schoolName string) (model.School, error) {
	resp, err := searchSchool(schoolName, "", 1)
	return resp[0], err
}
