package devices

import (
	"context"

	v1 "github.com/hoomy-official/go-kizbox/pkg/api/v1"
	"github.com/vanyda-official/go-shared/pkg/cmd"

	"github.com/hoomy-official/cli-kizbox/cmd/cli-kizbox/commands"
	"github.com/hoomy-official/cli-kizbox/cmd/cli-kizbox/filter"
	"github.com/hoomy-official/cli-kizbox/cmd/cli-kizbox/globals"

	"go.uber.org/zap"
)

// ApplyCmd apply a configuration to a resources by file name or stdin.
type ApplyCmd struct {
	filter.Filter
	Controllables []string
	Action        string `arg:"action"`
	Args          []string
}

func (s ApplyCmd) Run(global *globals.Globals, common *cmd.Commons) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	ctx := context.Background()
	api := global.Client()

	devices, err := commands.DeviceList(ctx, api, s.Controllables, s.Filter)
	if err != nil {
		logger.Error("cannot list device")
		return err
	}

	var actions []v1.Action
	for _, device := range devices {
		logger.Debug("open device", zap.Any("device", device))
		action := v1.Action{
			Commands: []v1.Command{
				{
					Name:       "ping",
					Parameters: []interface{}{},
				},
			},
			DeviceURL: device.DeviceURL,
		}

		actions = append(actions, action)
	}

	return api.V1.Execution.Apply(ctx, v1.Execute{Label: "cli command test", Actions: actions}, nil)
}
