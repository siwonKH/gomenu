package model

type Menu []struct {
	//Breakfast string `json:"breakfast"`
	//Lunch     string `json:"lunch"`
	//Dinner    string `json:"dinner"`
	MenuList [3]struct {
		MenuData []string `json:"menu"`
	} `json:"menu_list"`
}
