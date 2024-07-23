package globals

import (
	"fmt"
	"net/url"

	"github.com/hoomy-official/go-kizbox/pkg/client"
)

type Globals struct {
	APIKey string `env:"API_KEY" help:"apikey (retrieved through developer quickstart)"`
	Host   string `env:"HOST" help:"host of the Kizbox"`
	Scheme string `default:"https" help:"scheme of the Kizbox"`
	Port   int    `default:"8443" help:"port of the Kizbox"`
}

func (c *Globals) Client() *client.APIClient {
	u := &url.URL{
		Scheme: c.Scheme,
		Host:   fmt.Sprintf("%s:%d", c.Host, c.Port),
	}

	return client.NewClient(client.WithBaseURL(u), client.WithAuthToken(c.APIKey))
}
