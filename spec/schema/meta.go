package schema

type MetaData struct {
	Name             string
	DisplayName      string
	Description      string
	MinEngineVersion int               // The lowest version this app required
	AppVersion       string            // The application's version
	APIVersion       int               // This application's feature version
	Requires         []RequirementData // The requirement of this plugin(skill/plugin)
}
