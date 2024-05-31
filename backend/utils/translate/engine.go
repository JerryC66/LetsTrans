package translate

type Engine interface {
	Translate(text string, sourceLang string, targetLang string) (string, error)
}

func NewEngine(engine string) Engine {
	switch engine {
	case "deepl":
		return &DeepL{}
	case "volc":
		return &Volc{}
	default:
		return &DeepL{}
	}
}

const (
	LangEN = "EN"
	LangZH = "ZH"
	LangJP = "JP"
)
