package parser

import "github.com/siwonKH/gomenu/model"

func ParseSchool(schoolData model.NeisSchool) ([]model.School, error) {
	schoolDataList := schoolData.SchoolInfo[1].Row

	var schoolRes []model.School
	for i := 0; i < len(schoolDataList) && i < 10; i++ {
		var ele model.School
		ele.SchoolName = schoolDataList[i].SchulNm
		ele.SchoolCode = schoolDataList[i].SdSchulCode
		ele.AptName = schoolDataList[i].AtptOfcdcScNm
		ele.AptCode = schoolDataList[i].AtptOfcdcScCode
		schoolRes = append(schoolRes, ele)
	}

	return schoolRes, nil
}
