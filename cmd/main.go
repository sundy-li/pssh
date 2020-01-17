package main

import (
	"github.com/sundy-li/pssh"
)

func main() {
	options := pssh.Options

	for _, h := range options.Hosts {
		err := pssh.ExeSSH(h, options.Cmd)
		if err != nil {
			println("ERROR IN ssh ", err.Error())
		}
	}
}
