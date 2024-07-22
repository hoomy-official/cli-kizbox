package commands

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	v1 "github.com/hoomy-official/go-kizbox/pkg/api/v1"
	"github.com/vanyda-official/go-shared/pkg/cmd"

	"github.com/hoomy-official/cli-kizbox/globals"

	"go.uber.org/zap"
)

type ListenCmd struct {
	Interval time.Duration `help:"Interval between 2 polling (ns, ms, s and m)" default:"5s"`
}

func (l ListenCmd) Run(global *globals.Globals, common *cmd.Commons) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	logger = logger.With(zap.Duration("interval", l.Interval))

	cl := global.Client()
	ctx := context.Background()

	var eventRegister v1.EventRegister
	err = cl.V1.Event.Register(ctx, &eventRegister)
	if err != nil {
		return err
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	t := time.NewTicker(l.Interval)
	for {
		logger.Debug("polling new events")

		var events []map[string]interface{}
		if er := cl.V1.Event.Fetch(ctx, eventRegister, &events); er != nil {
			return er
		}

		for _, event := range events {
			logger.Info("new event", zap.Any("event", event))
		}

		select {
		case <-t.C:
			continue
		case <-c:
			return nil
		}
	}
}
