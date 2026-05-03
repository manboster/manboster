package browser

type NameType string

const (
	NameSearch  NameType = "search"
	NameWebpage NameType = "webpage"
)

type EngineType string

const (
	EngineGoogle     EngineType = "google"
	EngineDuckDuckGo EngineType = "duckduckgo"
	EngineBing       EngineType = "bing"
	EngineCNBing     EngineType = "cnbing"
	EngineBaidu      EngineType = "baidu"
	EngineGitHub     EngineType = "github"
	EngineWikipedia  EngineType = "wikipedia"
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
