package file

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runGrepInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "grep",
		DisplayName:  i18n.T(keys.FileGrepDisplayName),
		Description:  i18n.T(keys.FileGrepDescription),
		Represent:    "🔍",
		Irreversible: false,
	},
	Args:           schema.ArgsFromStruct(GrepArgs{}),
	Run:            runGrep,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: clientRendererFileName,
}

func runGrep(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, err := util.Unmarshal[GrepArgs](args)
	if err != nil {
		return nil, err
	}
	pwd, err := resolvePwd(ctx, arg.IsPublic)
	if err != nil {
		return nil, err
	}
	sPath, err := getSafePath(pwd, arg.FilePath, arg.FileName)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(sPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", sPath, err)
	}
	defer f.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(f)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if strings.Contains(line, arg.Keyword) {
			fmt.Fprintf(&sb, "%d: %s\n", lineNum, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan file %s: %w", sPath, err)
	}

	return &plugin.RunResponse{Response: sb.String()}, nil
}
