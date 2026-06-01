package skills

import "github.com/manboster/manboster/internal/tool"

func init() {
	fa := tool.NewFactory[NameType, *Service]()
	fa.RegisterProvider(&Service{})
	fa.Init()
}
