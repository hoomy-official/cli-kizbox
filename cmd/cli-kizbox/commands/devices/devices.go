package devices

type Cmd struct {
	List  ListCmd  `cmd:"list" help:"list devices available"`
	Get   GetCmd   `cmd:"get" help:"get device by URL"`
	Apply ApplyCmd `cmd:"apply" help:"apply"`
	RPC   ApplyCmd `cmd:"rpc" help:"rpc"`
}
