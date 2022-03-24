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
