package gomenu

type GoMenu struct {
	KEY string
}

func New(key string) *GoMenu {
	m := GoMenu{
		KEY: key,
	}
	return &m
}

//func main() {
//	request.SetKey("")
//	sum := time.Second * 0
//	log.Println(request.KEY)
//	for i := 0; i < 10; i++ {
//		start := time.Now()
//		schoolData, err := request.SearchSchool("경북소")
//		parsedSchool, err := parser.ParseSchool(schoolData)
//		menuData, err := request.SearchMenu(parsedSchool[0].AptCode, parsedSchool[0].SchoolCode, "20220107")
//		if err != nil {
//			log.Fatal(err)
//		}
//		parsed, err := parser.ParseMenu(menuData)
//		if err != nil {
//			log.Fatal(err)
//		}
//		elapsed := time.Since(start)
//		log.Println(parsed, elapsed)
//		sum += elapsed
//	}
//	log.Println(sum / 10)
//	time.Sleep(time.Second * 60)
//}
