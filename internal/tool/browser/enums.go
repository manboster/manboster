package browser

type NameType string

const (
	NameTypeSearch  NameType = "search"
	NameTypeWebpage NameType = "webpage"
)

type EngineType string

const (
	EngineTypeGoogle     EngineType = "google"
	EngineTypeDuckDuckGo EngineType = "duckduckgo"
	EngineTypeBing       EngineType = "bing"
	EngineTypeCNBing     EngineType = "cnbing"
	EngineTypeBaidu      EngineType = "baidu"
	EngineTypeGitHub     EngineType = "github"
	EngineTypeWikipedia  EngineType = "wikipedia"
)

type ResponseType string

const (
	ResponseTypeRaw      ResponseType = "raw"
	ResponseTypeMarkdown ResponseType = "markdown"
	ResponseTypeBody     ResponseType = "body"
)

type ScrapType string

const (
	ScrapTypeText    ScrapType = "text"
	ScrapTypeBrowser ScrapType = "browser"
)
