package pssh

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	ActionShell = "shell"
	ActionRsync = "rsync"
)

type (
	Hosts []string
	Opts  struct {
		Hosts        Hosts
		Hostfile     string
		Action       string
		Cmd          string
		AnsibleFile  string
		AnsibleGroup string

		//for rsync
		SrcPath string
		DstPath string
	}
)

func (i *Hosts) String() string {
	return "my string representation"
}

func (i *Hosts) Set(value string) error {
	hosts := strings.Split(value, ",")
	for _, h := range hosts {
		*i = append(*i, strings.TrimSpace(h))
	}
	return nil
}

var (
	Options Opts
)

func init() {
	var argsStart = 0
	Options.Action = ActionShell
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case ActionRsync:
			Options.Action = ActionRsync
			argsStart = 1
		case ActionShell:
			Options.Action = ActionShell
		}
	}

	var flagSet = flag.NewFlagSet(os.Args[argsStart], flag.ExitOnError)

	flagSet.Var(&Options.Hosts, "h", "host to execute the commands")
	flagSet.StringVar(&Options.Hostfile, "f", "", "hosts files")
	flagSet.StringVar(&Options.Cmd, "c", "", "execute commands")
	flagSet.StringVar(&Options.AnsibleFile, "a", "/etc/ansible/hosts", "ansible file path")
	flagSet.StringVar(&Options.AnsibleGroup, "g", "", "ansible hosts group")

	flagSet.StringVar(&Options.SrcPath, "s", "", "rsync source path")
	flagSet.StringVar(&Options.DstPath, "d", "", "rsync dist path")

	err := flagSet.Parse(os.Args[argsStart+1:])
	if err != nil {
		panic(err)
	}

	Options.initConfig()
}

func (opt *Opts) initConfig() {
	if opt.Hostfile != "" {
		fp, err := os.Open(opt.Hostfile)
		if err != nil {
			log.Fatalf("Hostfile %s open error:%s", opt.Hostfile, err.Error())
		}
		bs, _ := ioutil.ReadAll(fp)
		for _, str := range strings.Split(string(bs), "\n") {
			h := strings.TrimSpace(str)
			if len(h) > 0 {
				opt.Hosts = append(opt.Hosts, h)
			}
		}
	}

	if opt.AnsibleFile != "" && opt.AnsibleGroup != "" {
		hosts, err := ParseAnisbleHost(opt.AnsibleFile)
		if err != nil {
			log.Fatalf("Parsing ansible file error:%s", err.Error())
		}
		groups := strings.Split(opt.AnsibleGroup, ",")
		for _, g := range groups {
			if hs, ok := hosts[strings.TrimSpace(g)]; ok {
				opt.Hosts = append(opt.Hosts, hs...)
			} else {
				log.Fatalf("Not found group %s", g)
			}
		}
	}
}
