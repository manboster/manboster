package file

import (
	"github.com/manboster/manboster/internal/tool"
)

func init() {
	fa := tool.NewFactory[NameType, *Service]()
	fa.RegisterProvider(&Service{})
	fa.RegisterNamespace(NameRead, &runReadInfo)
	fa.RegisterNamespace(NameWrite, &runWriteInfo)
	fa.RegisterNamespace(NameInfo, &runInfoInfo)
	fa.RegisterNamespace(NameList, &runListInfo)
	fa.RegisterNamespace(NameDir, &runDirInfo)
	fa.RegisterNamespace(NameDelete, &runDeleteInfo)
	fa.RegisterNamespace(NameGrep, &runGrepInfo)
	fa.RegisterNamespace(NameReplace, &runReplaceInfo)
	fa.Init()
}
