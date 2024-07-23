package commands

import (
	"context"
	"log"
	"time"

	"github.com/hoomy-official/go-kizbox/pkg/discover"
	"github.com/vanyda-official/go-shared/pkg/cmd"

	"github.com/hoomy-official/cli-kizbox/cmd/cli-kizbox/filter"
)

type DiscoverCmd struct {
	filter.Filter

	Timeout time.Duration `default:"5s" help:"timeout for discovering (ns, ms, s & m)"`
}

func (d *DiscoverCmd) Run(_ *cmd.Commons) error {
	ctx := context.Background()

	ch := make(chan *discover.Gateway, 1)

	go func() {
		for gateway := range ch {
			log.Println(gateway)
		}
	}()

	return discover.NewKizboxDiscover(d.Timeout).Discover(ctx, ch)
}
