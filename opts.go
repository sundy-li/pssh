package pssh

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type (
	Hosts []string
	Opts  struct {
		Hosts        Hosts
		Hostfile     string
		Cmd          string
		AnsibleFile  string
		AnsibleGroup string
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
	flag.Var(&Options.Hosts, "h", "host to execute the commands")
	flag.StringVar(&Options.Hostfile, "f", "", "hosts files")
	flag.StringVar(&Options.Cmd, "c", "", "execute commands")
	flag.StringVar(&Options.AnsibleFile, "a", "/etc/ansible/hosts", "ansible file path")
	flag.StringVar(&Options.AnsibleGroup, "g", "", "ansible hosts group")

	flag.Parse()

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
