package hostutils

import (
	"golang.org/x/exp/constraints"
	"regexp"
)

var (
	reComment    = regexp.MustCompile(`#.*`)
	reSpaces     = regexp.MustCompile(`\s+`)
	rePackedHost = regexp.MustCompile(`^([^\[]*)\[([-,:0-9\s]+)](.*)$`)
	reCondSpace  = regexp.MustCompile(`,\s*`)
	reCondBlk    = regexp.MustCompile(`^(\d+)([-:](\d+))?$`)
)

func max[T constraints.Ordered](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func min[T constraints.Ordered](a T, b T) T {
	if a < b {
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
		noCmtHosts := reComment.ReplaceAllString(host, "")
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
