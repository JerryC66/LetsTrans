package letstrans

type RouterGroup struct {
	FileRouter
	ProjectRouter
	SegmentRouter
	GlossaryRouter
	ThirdPartyRouter
	TranslationMemoryRouter
}
