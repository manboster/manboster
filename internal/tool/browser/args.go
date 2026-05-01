package browser

import "github.com/manboster/manboster/spec/schema"

type RunArgs struct {
	Name         NameType     `json:"name" description:"The name you want call, it would be enum, only 2 values: search and webpage. 'search' returns data from the search engine while 'webpage' returns the information you want to get." validate:"required" enum:"search,webpage" example:"search"`
	Keywords     string       `json:"keywords" description:"The keywords you want to search in search engine. Only valid when name = 'search'." example:"What's the meaning of Manboster?'"`
	Engine       EngineType   `json:"engine" description:"The search engine you want to use, it would be enum, only 5 values: google, bing, cnbing, duckduckgo and baidu. Only valid when name = 'search'." example:"google"`
	ScrapType    ScrapType    `json:"scrap_type" description:"The method you want to use to scrap the webpage, it would be enum, only 2 values: text and browser. 'text' would use get method Golang builtin functions and cost less resources while 'browser' will return the webpage rendered to you. Only valid when name = 'webpage'. Use webpage first, if this couldn't solve the problem, try browser." example:"text"`
	ResponseType ResponseType `json:"response_type" description:"The type you want to get in response, it would be enum, only 3 values: raw, body and markdown. 'raw' returns the raw html rendered data from the page. 'body' only returns html's body content back, and 'markdown' returns markdown converted webpage to you." example:"raw"`
	Page         string       `json:"page" description:"The page URL you want to scrap. Only valid when name='webpage'." example:"https://example.com"`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(&RunArgs{})
}
