package schema

type MetaData struct {
	Name             string            `json:"name" yaml:"name"`
	DisplayName      string            `json:"display_name" yaml:"display_name"`
	Description      string            `json:"description" yaml:"description"`
	MinEngineVersion int               `json:"min_engine_version" yaml:"min_engine_version"` // The lowest version this app required
	AppVersion       string            `json:"app_version" yaml:"app_version"`               // The application's version
	APIVersion       int               `json:"api_version" yaml:"api_version"`               // This application's feature version
	Requires         []RequirementData `json:"requires" yaml:"requires"`                     // The requirement of this plugin(skill/plugin)
}
