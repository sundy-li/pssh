package pssh

import "flag"

type (
	Hosts []string

	Opts struct {
		Hosts Hosts
		Cmd   string
	}
)

func (i *Hosts) String() string {
	return "my string representation"
}

func (i *Hosts) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var (
	Options Opts
)

func init() {
	flag.Var(&Options.Hosts, "host", "host to execute the commands")
	flag.StringVar(&Options.Cmd, "cmd", "", "execute commands")
	flag.Parse()
}
