package devices

import (
	"context"
	"fmt"

	"github.com/hoomy-official/cli-kizbox/globals"
	v1 "github.com/hoomy-official/go-kizbox/pkg/api/v1"
	"github.com/vanyda-official/go-shared/pkg/cmd"
)

type GetCmd struct {
	URL string `arg:"URL"`
}

func (d *GetCmd) Run(global *globals.Globals, common *cmd.Commons) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	cl := global.Client()
	ctx := context.Background()

	var device v1.Device

	err = cl.V1.Devices.Get(ctx, d.URL, &device)
	if err != nil {
		logger.Error("cannot get device")
		return err
	}

	//nolint:forbidigo // let's not overengineer stdout
	fmt.Printf("%v", device)

	return nil
}
