package skill

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/fatih/color"
)

func LoadSkill(path string, name string, isDir bool) (*Skill, error) {
	s := &Skill{}
	s.IsDirectory = isDir
	s.Name = name

	filePath := path
	if isDir {
		filePath = filepath.Join(path, "SKILLS.md")
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	content := string(fileContent)
	content = strings.ReplaceAll(content, "\r\n", "\n")

	var data Frontmatter

	rest, err := frontmatter.Parse(strings.NewReader(content), &data)
	if err != nil {
		fmt.Printf("failed to get: %v\n", err)
		return s, err
	}
	s.Content = string(rest)

	s.DisplayName = data.Name
	s.Description = data.Description
	s.Homepage = data.Homepage

	var ocMeta Metadata
	err = json.Unmarshal([]byte(data.MetadataRaw), &ocMeta)
	if err != nil {
		color.Yellow("failed to unmarshal OpenClaw metadata: %v\n", err)
	}
	s.RepresentEmoji = ocMeta.OpenClaw.Emoji

	return s, nil
}
