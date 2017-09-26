package hostutils

import (
	"regexp"
	"strconv"
)

var reComent = regexp.MustCompile(`#.*$`)
var reSpaces = regexp.MustCompile(`\s+`)
var reHostname = regexp.MustCompile(`^(.*?)(\d+)(\D*)$`)
var reIsNumber = regexp.MustCompile(`^\d+$`)

var rePackedHost = regexp.MustCompile(`^([^\[]*)\[([-,:0-9\s]+)](.*)$`)
var reCond = regexp.MustCompile(`^\d+([-:]\d+)?(,\s*(\d+([-:]\d+)?))*$`)
var reCondSpace = regexp.MustCompile(`,\s*`)
var reCondBlk = regexp.MustCompile(`^(\d+)([-:](\d+))?$`)

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func maxi(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
