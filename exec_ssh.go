package pssh

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"sync"
)

func ExeParallelSSH(options Opts) (err error) {
	var wg sync.WaitGroup
	wg.Add(len(options.Hosts))

	for i := range options.Hosts {
		go func(i int) {
			defer func() {
				if e := recover(); e != nil {
				}
				wg.Done()
			}()

			switch options.Action {
			case ActionShell:
				cmd, args := buildSshCmd(options.Hosts[i], options.Cmd)
				ExecSSH(options.Hosts[i], cmd, args)
			case ActionRsync:
				cmd, args := buildRsyncCmd(options.Hosts[i], options.SrcPath, options.DstPath)
				ExecSSH(options.Hosts[i], cmd, args)
			}
		}(i)
	}
	wg.Wait()
	return nil
}

func buildSshCmd(host string, command string) (cmd string, args []string) {
	args = append(args, "-q")
	args = append(args, "-oNumberOfPasswordPrompts=0")
	args = append(args, "-oStrictHostKeyChecking=no")
	args = append(args, host)
	args = append(args, command)

	return "ssh", args
}

func buildRsyncCmd(host string, srcPath, dstPath string) (cmd string, args []string) {
	args = append(args, "-vlcr")
	args = append(args, srcPath)
	args = append(args, host+":"+dstPath)
	return "rsync", args
}

func ExecSSH(host string, command string, args []string) (err error) {
	cmd := exec.Command(command, args...)
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
