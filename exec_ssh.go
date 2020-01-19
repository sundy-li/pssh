package pssh

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"sync"
)

func ExeParallelSSH(hosts []string, cmd string) (err error) {
	var wg sync.WaitGroup
	wg.Add(len(hosts))
	for _, h := range hosts {
		go func(h, cmd string) {
			defer func() {
				if e := recover(); e != nil {
				}
				wg.Done()
			}()

			ExeSSH(h, cmd)
		}(h, cmd)
	}
	wg.Wait()
	return nil
}

func ExeSSH(host string, command string) (err error) {
	args := []string{}

	args = append(args, "-q")
	args = append(args, "-oNumberOfPasswordPrompts=0")
	args = append(args, "-oStrictHostKeyChecking=no")
	args = append(args, host)
	args = append(args, command)

	// fmt.Printf("ssh %v", args)
	cmd := exec.Command("ssh", args...)

	stderr, err := cmd.StderrPipe()
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := ioutil.ReadAll(stderr)
	slurt, _ := ioutil.ReadAll(stdout)

	if len(slurp) > 0 {
		fmt.Printf("%s Err: %s\n", host, slurp)
	}

	strs := strings.Split(string(slurt), "\n")
	for _, str := range strs {
		if len(str) > 0 {
			fmt.Printf("[%s]: %s\n", host, str)
		}
	}
	if err = cmd.Wait(); err != nil {
		return
	}
	return
}
