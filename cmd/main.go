package main

import (
	"github.com/sundy-li/pssh"
	"log"
)

func main() {
	options := pssh.Options

	log.Printf("hosts -> %v", options.Hosts)
	err := pssh.ExeParallelSSH(options.Hosts, options.Cmd)
	if err != nil {
		log.Fatal("ERROR IN ssh ", err.Error())
	}
}
