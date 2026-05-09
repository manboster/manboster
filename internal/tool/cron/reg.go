package cron

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/repository/types"
)

func (s *Service) Register(cj types.Cron) error {
	color.Blue(fmt.Sprintf("[Manboster Tool Provider] Register cron job: %q", cj.Name))
	entryID, err := s.cron.AddFunc(cj.CronTab, func() {
		s.Runner(buildMessageDataFromCronJob(cj))
	})

	fmt.Printf("%+v", s.cron.Entry(entryID))
	if err != nil {
		return err
	}
	s.manager.SetEntry(cj.Name, int(entryID))
	s.manager.Load(cj.Name, true)
	return nil
}

func (s *Service) FullRegister(ctx context.Context) error {
	allData, err := s.cronRepo.GetAllCronjob(ctx)
	if err != nil {
		return err
	}

	c := 0
	for _, entry := range allData {
		err := s.Register(entry)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Tool Provider] Failed to register entry %q: %q", entry.Name, err))
		} else {
			c++
		}
	}
	color.Green(fmt.Sprintf("[Manboster Tool Provider] Register %d jobs!", c))
	return nil
}
