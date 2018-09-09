package main

import (
	"flag"
	"strings"
)

var (
	hostMasq    = flag.String("host-masq", "", "rewrite nexthop hosts (format: from1=to1,from2=to2)")
	hostMasqMap = map[string]string{}
)

func init() {
	if *hostMasq != "" {
		hostMasqSetup(*hostMasq)
	}
}

func hostMasqSetup(cfg string) {
	for _, line := range strings.Split(cfg, ",") {
		elts := strings.Split(line, "=")
		hostMasqMap[elts[0]] = elts[1]
	}
}

func hostRewrite(host string) string {
	if len(hostMasqMap) == 0 {
		return host
	}
	for k, v := range hostMasqMap {
		if strings.HasSuffix(host, k) {
			host = strings.Replace(host, k, v, -1)
		}
	}
	return host
}