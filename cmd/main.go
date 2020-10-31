package main

import (
	"log"

	"github.com/sundy-li/pssh"
)

func main() {
	options := pssh.Options

	log.Printf("hosts -> %v", options.Hosts)
	err := pssh.ExeParallelSSH(options)
	if err != nil {
		log.Fatal("ERROR IN ssh ", err.Error())
	}
}
