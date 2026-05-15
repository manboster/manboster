package browser

import "github.com/manboster/manboster/spec/schema"

type RunArgs struct {
	Name         NameType     `json:"name" description:"Operation to run: search queries a search engine; webpage fetches a URL." validate:"required" enum:"search,webpage" example:"search"`
	Keywords     string       `json:"keywords" description:"Search query. Required when name is search." example:"What is Manboster?"`
	Engine       EngineType   `json:"engine" description:"Search engine for search. Prefer bing or cnbing when unsure." example:"bing"`
	ScrapType    ScrapType    `json:"scrap_type" description:"Fetch mode for webpage. text is lighter; browser renders the page." example:"text"`
	ResponseType ResponseType `json:"response_type" description:"Output format for webpage: raw HTML, body HTML, or markdown." example:"markdown"`
	URL          string       `json:"url" description:"URL to fetch. Required when name is webpage." example:"https://example.com"`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(&RunArgs{})
}
