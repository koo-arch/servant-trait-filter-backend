package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type SyncAtlas interface {
	Sync(ctx context.Context) error
}

type Scheduler struct {
	cron *cron.Cron
	etl SyncAtlas
}

func NewScheduler(etl SyncAtlas) *Scheduler {
	return &Scheduler{
		cron: cron.New(),
		etl: etl,
	}
}

func (s *Scheduler) SetupJobs(ctx context.Context) {
	_, err := s.cron.AddFunc("@daily", func() {
		defer func () {
			if r := recover(); r != nil {
				log.Printf("panic in sync job: %v", r)
			}
		}()

		jobCtx, cancel := context.WithTimeout(ctx, time.Minute * 5)
		defer cancel()

		err := s.etl.Sync(jobCtx)
		if err != nil {
			log.Printf("failed to sync atlas api: %v", err)
		}
	})
	if err != nil {
		log.Printf("failed to add cron job: %v", err)
	}
}

func (s *Scheduler) Start() {
	s.cron.Start()
	log.Println("scheduler started")
}

func (s *Scheduler) Stop() {
	ctx := s.cron.Stop()
	<-ctx.Done()
	log.Println("scheduler stopped")
}