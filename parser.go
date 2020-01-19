package pssh

import (
	"bufio"
	"os"
	"strings"
)

func ParseAnisbleHost(fp string) (hosts map[string][]string, err error) {
	f, err := os.Open(fp)
	if err != nil {
		return
	}

	sc := bufio.NewScanner(f)
	hosts = make(map[string][]string)
	var key string
	var keyHosts []string
	for sc.Scan() {
		str := strings.TrimSpace(sc.Text())
		// comment
		if strings.Index(str, "#") == 0 || len(str) == 0 {
			continue
		} else if strings.Index(str, "[") == 0 && strings.Index(str, "]") == len(str)-1 {
			hosts[key] = keyHosts

			key = str[1 : len(str)-1]
			keyHosts = keyHosts[:0]
		} else {
			keyHosts = append(keyHosts, str)
		}
	}
	if len(keyHosts) > 0 {
		hosts[key] = keyHosts
	}
	return
}
