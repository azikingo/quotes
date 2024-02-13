package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
)

type CronServer struct {
	log  *log.Helper
	cron *cron.Cron
}

// NewCronServer
func NewCronServer(
	logger log.Logger,
	bot *BotServer,
) *CronServer {
	cs := &CronServer{
		log:  log.NewHelper(log.With(logger, "module", "server/cron")),
		cron: cron.New(),
	}

	cs.checkEventsNotify(bot)

	return cs
}

func (cs *CronServer) checkEventsNotify(bot *BotServer) {
	entryId, err := cs.cron.AddFunc("@every 3h", func() {
		err := bot.SendQuotes()
		if err != nil {
			cs.log.Errorf("failed on cron, err: %v", err)
			return
		}
	})
	if err != nil {
		cs.log.Errorf("failed on cron entryId: %v, err: %v", entryId, err)
		return
	}
}

func (cs *CronServer) Start(ctx context.Context) error {
	cs.cron.Start()
	cs.log.Info("cron server started")

	return nil
}

func (cs *CronServer) Stop(ctx context.Context) error {
	cs.cron.Stop()
	cs.log.Info("cron server stopped")

	return nil
}
