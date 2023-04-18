package hostutils

import (
	"bytes"
	"fmt"
	"golang.org/x/exp/slices"
	"strconv"
)

type host struct {
	NonDigits []string
	Digits    []digit
}

type digit struct {
	Value int
	Digit int
}

type hostGroup struct {
	host
	RangeIndex int
	Conds      []cond
}

type cond struct {
	Low   int
	High  int
	Digit int
}

// PackString Pack space separated full hosts list into short abbreviated hosts.
func PackString(hosts string) []string {
	return Pack([]string{hosts})
}

// Pack full hosts list into short abbreviated hosts.
func Pack(hosts []string) []string {
	regHosts := regularizeHosts(hosts)
	if len(regHosts) == 0 {
		return nil
	}
	if len(regHosts) == 1 {
		return regHosts
	}
	uniqHosts := make([]*host, 0, len(regHosts))
	for _, host := range regHosts {
		uniqHosts = append(uniqHosts, parseHost(host))
	}
	slices.SortFunc(uniqHosts, func(a, b *host) bool {
		return a.Less(b)
	})

	result := make([]string, 0, len(uniqHosts))

	var group *hostGroup
	for i := 1; i < len(uniqHosts); i++ {
		if group == nil {
			group = mergeHost(uniqHosts[i-1], uniqHosts[i])
			if group == nil {
				result = append(result, uniqHosts[i-1].String())
			}
		} else {
			if !group.AppendHost(uniqHosts[i]) {
				result = append(result, group.String())
				group = nil
			}
		}
	}
	if group == nil {
		result = append(result, uniqHosts[len(uniqHosts)-1].String())
	} else {
		result = append(result, group.String())
	}

	return result
}

func mergeHost(h1, h2 *host) *hostGroup {
	if len(h1.NonDigits) != len(h2.NonDigits) || len(h1.Digits) != len(h2.Digits) {
		return nil
	}
	idx := -1
	for i := range h1.NonDigits {
		if h1.NonDigits[i] != h2.NonDigits[i] {
			return nil
		}
		if i < len(h1.Digits) {
			if idx == -1 {
				if h1.Digits[i] != h2.Digits[i] {
					idx = i
				}
			} else if h1.Digits[i] != h2.Digits[i] {
				return nil
			}
		}
	}
	if idx == -1 {
		panic("idx must not be -1 unless h1 and h2 are same")
	}
	g := &hostGroup{
		host:       *h1,
		RangeIndex: idx,
	}
	if h1.Digits[idx].Digit == h2.Digits[idx].Digit && h1.Digits[idx].Value+1 == h2.Digits[idx].Value {
		g.Conds = append(g.Conds, cond{
			Low:   h1.Digits[idx].Value,
			High:  h2.Digits[idx].Value,
			Digit: h1.Digits[idx].Digit,
		})
	} else {
		g.Conds = append(g.Conds, cond{
			Low:   h1.Digits[idx].Value,
			High:  h1.Digits[idx].Value,
			Digit: h1.Digits[idx].Digit,
		}, cond{
			Low:   h2.Digits[idx].Value,
			High:  h2.Digits[idx].Value,
			Digit: h2.Digits[idx].Digit,
		})
	}
	return g
}

func (g *hostGroup) AppendHost(h *host) bool {
	if len(g.NonDigits) != len(h.NonDigits) || len(g.Digits) != len(h.Digits) {
		return false
	}
	idx := g.RangeIndex
	for i := range g.NonDigits {
		if g.NonDigits[i] != h.NonDigits[i] {
			return false
		}
		if i != idx && i < len(g.Digits) && g.Digits[i] != h.Digits[i] {
			return false
		}
	}
	lastRange := g.Conds[len(g.Conds)-1]
	if lastRange.Digit == h.Digits[idx].Digit && lastRange.High+1 == h.Digits[idx].Value {
		g.Conds[len(g.Conds)-1].High = h.Digits[idx].Value
	} else {
		g.Conds = append(g.Conds, cond{
			Low:   h.Digits[idx].Value,
			High:  h.Digits[idx].Value,
			Digit: h.Digits[idx].Digit,
		})
	}
	return true
}

func (g *hostGroup) String() string {
	var buf bytes.Buffer
	for i := range g.NonDigits {
		buf.WriteString(g.NonDigits[i])
		if i != g.RangeIndex {
			if i < len(g.Digits) {
				_, _ = fmt.Fprintf(&buf, "%0*d", g.Digits[i].Digit, g.Digits[i].Value)
			}
		} else {
			buf.WriteByte('[')
			for j, r := range g.Conds {
				if j != 0 {
					buf.WriteByte(',')
				}
				if r.Low == r.High {
					_, _ = fmt.Fprintf(&buf, "%0*d", r.Digit, r.Low)
				} else {
					_, _ = fmt.Fprintf(&buf, "%0*d-%0*d", r.Digit, r.Low, r.Digit, r.High)
				}
			}
			buf.WriteByte(']')
		}
	}
	return buf.String()
}

func parseHost(s string) *host {
	if s == "" {
		return &host{
			NonDigits: []string{""},
		}
	}
	host := &host{}
	buf := make([]rune, 0, len(s))
	digitMode := false
	for _, c := range s {
		digit := isDigit(c)
		if digitMode == digit {
			buf = append(buf, c)
		} else {
			host.appendToken(string(buf), digitMode)
			buf = buf[:0]
			buf = append(buf, c)
			digitMode = !digitMode
		}
	}
	if len(buf) > 0 {
		host.appendToken(string(buf), digitMode)
	}
	return host
}

func (h *host) String() string {
	var buf bytes.Buffer
	for i, s := range h.NonDigits {
		buf.WriteString(s)
		if i < len(h.Digits) {
			_, _ = fmt.Fprintf(&buf, "%0*d", h.Digits[i].Digit, h.Digits[i].Value)
		}
	}
	return buf.String()
}

func (h *host) appendToken(token string, digitMode bool) {
	if digitMode {
		n, _ := strconv.Atoi(token)
		h.Digits = append(h.Digits, digit{
			Value: n,
			Digit: len(token),
		})
	} else {
		h.NonDigits = append(h.NonDigits, token)
	}
}

func (h *host) Less(rhs *host) bool {
	m := min(len(h.NonDigits), len(rhs.NonDigits))
	for i := 0; i < m; i++ {
		if h.NonDigits[i] < rhs.NonDigits[i] {
			return true
		}
		if h.NonDigits[i] > rhs.NonDigits[i] {
			return false
		}
	}
	if len(h.NonDigits) < len(rhs.NonDigits) {
		return true
	}
	if len(h.NonDigits) > len(rhs.NonDigits) {
		return false
	}

	m = min(len(h.Digits), len(rhs.Digits))
	for i := 0; i < m; i++ {
		if h.Digits[i].Digit < rhs.Digits[i].Digit {
			return true
		}
		if h.Digits[i].Digit > rhs.Digits[i].Digit {
			return false
		}
		if h.Digits[i].Value < rhs.Digits[i].Value {
			return true
		}
		if h.Digits[i].Value > rhs.Digits[i].Value {
			return false
		}
	}
	if len(h.Digits) < len(rhs.Digits) {
		return true
	}
	return false
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}
