package cron

import (
	"context"

	"github.com/manboster/manboster/internal/repository/types"
	"github.com/robfig/cron/v3"
)

func (s *Service) Create(ctx context.Context, arg RunArgs, chatId string, chatProvider string, userId string) error {
	if isDelayFormat(arg.Cron) {
		d, err := parseDelay(arg.Cron)
		if err != nil {
			return err
		}

		go s.DelayRunner(d, buildMessageDataFromArgs(arg, chatId, chatProvider, userId))
		return nil
	}

	var cj types.Cron
	cj.CreateBy = userId
	cj.ChatProvider = chatProvider
	cj.CronTab = arg.Cron
	cj.Name = arg.JobName
	cj.Prompt = arg.Prompt
	cj.Type = string(arg.MessageType)
	cj.Ignore = string(arg.Ignore)
	switch arg.To {
	case ToThisChat:
		cj.ChatID = chatId
	case ToPM:
		cj.ChatID = userId
	default:
		cj.ChatID = chatId
	}
	err := s.Register(cj)
	if err != nil {
		return err
	}
	return s.cronRepo.CreateCronjob(ctx, cj)
}

func (s *Service) Delete(ctx context.Context, name string) error {
	id, avail := s.manager.GetEntry(name)
	if avail {
		s.cron.Remove(cron.EntryID(id))
	}
	return s.cronRepo.DeleteCronjob(ctx, name)
}

func (s *Service) List(ctx context.Context, provider string, chat string) ([]string, error) {
	var resp []string
	data, err := s.cronRepo.GetCronjobByChatID(ctx, chat, provider)
	if err != nil {
		return nil, err
	}
	for _, d := range data {
		resp = append(resp, d.Name)
	}
	return resp, nil
}

func (s *Service) Get(ctx context.Context, name string) (types.Cron, error) {
	return s.cronRepo.GetCronjobByName(ctx, name)
}
