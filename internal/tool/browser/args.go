package browser

import "github.com/manboster/manboster/spec/schema"

type SearchArgs struct {
	Keywords string     `json:"keywords" description:"Search query." example:"What is Manboster?" validate:"required"`
	Engine   EngineType `json:"engine" description:"Search engine. Prefer bing or cnbing when unsure." example:"bing" validate:"required" enum:"google,duckduckgo,bing,cnbing,baidu,github,wikipedia"`
}

type WebpageArgs struct {
	URL          string       `json:"url" description:"URL to fetch." example:"https://example.com" validate:"required"`
	ScrapType    ScrapType    `json:"scrap_type" description:"Fetch mode. text is lighter; browser renders the page." example:"text" enum:"text,browser"`
	ResponseType ResponseType `json:"response_type" description:"Output format: raw HTML, body HTML, or markdown." example:"markdown" enum:"raw,body,markdown"`
}

func (s *Service) Args() *schema.Args {
	return nil
}
