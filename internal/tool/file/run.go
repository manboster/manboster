package file

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/go-viper/mapstructure/v2"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/spec/plugin"
)

func (s *Service) Init(ctx context.Context, conf any) error {
	err := os.MkdirAll(config.Path("workspace"), 0755)
	if err != nil {
		return err
	}
	err = os.MkdirAll(config.Path(filepath.Join("workspace", "public")), 0755)
	if err != nil {
		return err
	}

	var cfg Config
	err = mapstructure.Decode(conf, &cfg)
	if err != nil {
		return err
	}

	s.cfg = &cfg
	err = s.cfg.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Start(ctx context.Context) error {
	return nil
}

func (s *Service) Run(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg := RunArgs{}
	resp := &plugin.RunResponse{
		Hangup: false,
	}

	sessID, ok := ctx.Value("session_id").(string)
	if !ok {
		return nil, fmt.Errorf("session_id not found in context")
	}
	pwd := config.Path(filepath.Join("workspace", "session-"+sessID))
	err := os.MkdirAll(pwd, 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create session dir: %w", err)
	}

	if json.Unmarshal([]byte(args), &arg) == nil {
		if arg.IsPublic {
			pwd = config.Path(filepath.Join("workspace", "public"))
		}
		switch arg.Name {
		case NameRead:
			sPath, err := getSafePath(pwd, arg.FilePath, arg.FileName)
			if err != nil {
				return nil, err
			}
			data, err := os.ReadFile(sPath)
			if err != nil {
				return nil, fmt.Errorf("failed to read file %s: %w", sPath, err)
			}
			resp.Response = string(data)
		case NameInfo:
			sPath, err := getSafePath(pwd, arg.FilePath, arg.FileName)
			if err != nil {
				return nil, err
			}
			info, err := os.Stat(sPath)
			if err != nil {
				return nil, fmt.Errorf("failed to get file info %s: %w", sPath, err)
			}
			jsonify, err := json.Marshal(newFileEntry(info))
			if err != nil {
				return nil, fmt.Errorf("failed to jsonify file info %s: %w", arg.Name, err)
			}
			resp.Response = string(jsonify)
		case NameList:
			if arg.FileName != "" {
				return nil, fmt.Errorf("filename is not allowed to give while name is list")
			}
			sPath, err := getSafePath(pwd, arg.FilePath, arg.FileName)
			if err != nil {
				return nil, err
			}
			d, err := listDir(sPath)
			if err != nil {
				return nil, fmt.Errorf("failed to read dir %s: %w", sPath, err)
			}
			jsonify, err := json.Marshal(d)
			if err != nil {
				return nil, fmt.Errorf("failed to jsonify dir %s: %w", arg.FileName, err)
			}
			resp.Response = string(jsonify)
		case NameDir:
			resp.Response = pwd
		case NameDelete:
			if s.cfg.Mode == "readonly" {
				return nil, fmt.Errorf("failed to delete: read-only mode set by user")
			}
			sPath, err := getSafePath(pwd, arg.FilePath, arg.FileName)
			if err != nil {
				return nil, err
			}
			err = os.Remove(sPath)
			if err != nil {
				return nil, fmt.Errorf("failed to delete file %s: %w", sPath, err)
			}
			resp.Response = "success"
		case NameWrite:
			if s.cfg.Mode == "readonly" {
				return nil, fmt.Errorf("failed to write: read-only mode set by user")
			}
			sPath, err := getSafePath(pwd, arg.FilePath, arg.FileName)
			if err != nil {
				return nil, err
			}
			err = os.WriteFile(sPath, []byte(arg.Content), 0644)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Tool Provider] dev.manboster.file failed to write file: %q", err))
				return nil, fmt.Errorf("failed to write file %s:%w", sPath, err)
			}
			resp.Response = "Success"
		default:
			return nil, fmt.Errorf("unknown argument %q", arg.Name)
		}
	} else {
		return nil, fmt.Errorf("invalid arguments")
	}
	return resp, nil
}

func (s *Service) Continue(ctx context.Context, session string) (*plugin.RunResponse, error) {
	return nil, nil
}

func (s *Service) Close() error {
	return nil
}
