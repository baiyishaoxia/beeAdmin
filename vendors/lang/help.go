package lang

type Types int

const (
	ZhCN Types = iota + 1
	EnUS
	Ja
)

const (
	ZhCNCode string = "zh-CN"
	EnUSCode        = "en-US"
	JaCode          = "ja"
)

func GetLangId(lang string) Types {
	switch lang {
	case ZhCNCode:
		return ZhCN
	case EnUSCode:
		return EnUS
	case JaCode:
		return Ja
	default:
		return EnUS
	}
}

func (l *Types) String() string {
	switch *l {
	case ZhCN:
		return ZhCNCode
	case EnUS:
		return EnUSCode
	case Ja:
		return JaCode
	default:
		return EnUSCode
	}
}
