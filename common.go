package hostutils

import (
	"regexp"
	"strconv"
	"strings"
)

var reComent = regexp.MustCompile(`#.*`)
var reSpaces = regexp.MustCompile(`\s+`)
var reHostname = regexp.MustCompile(`^(.*?)(\d+)(\D*)$`)
var reIsNumber = regexp.MustCompile(`^\d+$`)

var rePackedHost = regexp.MustCompile(`^([^\[]*)\[([-,:0-9\s]+)](.*)$`)
var reCond = regexp.MustCompile(`^\d+([-:]\d+)?(,\s*(\d+([-:]\d+)?))*$`)
var reCondSpace = regexp.MustCompile(`,\s*`)
var reCondBlk = regexp.MustCompile(`^(\d+)([-:](\d+))?$`)

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func maxi(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func regularizeHosts(hosts []string) []string {
	if hosts == nil {
		return nil
	}
	uniqHosts := make(map[string]bool)
	for _, host := range hosts {
		noCmtHosts := reComent.ReplaceAllString(host, "")
		for _, h := range reSpaces.Split(noCmtHosts, -1) {
			if h != "" {
				uniqHosts[h] = true
			}
		}
	}
	result := make([]string, len(uniqHosts))
	var i = 0
	for host := range uniqHosts {
		result[i] = host
		i++
	}
	return result
}

func parseFQDN(fqdn string) (hostname, domain string) {
	tokens := strings.SplitN(fqdn, ".", 2)
	if len(tokens) == 2 {
		return tokens[0], "." + tokens[1]
	}
	return fqdn, ""
}
