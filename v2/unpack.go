package hostutils

import (
	"fmt"
	"sort"
	"strconv"
)

// UnpackString Unpack space separated short abbreviated hosts into full hosts list.
func UnpackString(packedHosts string) (hosts []string) {
	return Unpack([]string{packedHosts})
}

// Unpack short abbreviated hosts into full hosts list.
func Unpack(packedHosts []string) (hosts []string) {
	regHosts := regularizeHosts(packedHosts)
	if regHosts == nil {
		return nil
	}
	resultSet := make(map[string]bool)

	for _, packedHost := range regHosts {
		unpackHosts(packedHost, resultSet)
	}

	result := make([]string, len(resultSet))
	i := 0
	for key := range resultSet {
		result[i] = key
		i++
	}
	sort.Strings(result)
	return result
}

func unpackHosts(packedHost string, resultSet map[string]bool) {
	m := rePackedHost.FindStringSubmatch(packedHost)
	if m != nil {
		prefix := m[1]
		cond := m[2]
		suffix := m[3]
		for _, num := range unpackCond(cond) {
			newHost := fmt.Sprintf("%s%s%s", prefix, num, suffix)
			unpackHosts(newHost, resultSet)
		}
	} else {
		resultSet[packedHost] = true
	}
}

func unpackCond(cond string) []string {
	var result []string

	for _, blk := range reCondSpace.Split(cond, -1) {
		m := reCondBlk.FindStringSubmatch(blk)
		if m != nil {
			if m[2] == "" {
				result = append(result, m[1])
			} else {
				digit := max(len(m[1]), len(m[3]))
				low, _ := strconv.Atoi(m[1])
				high, _ := strconv.Atoi(m[3])
				if low > high {
					low, high = high, low
				}
				for i := low; i <= high; i++ {
					result = append(result, fmt.Sprintf("%0*d", digit, i))
				}
			}
		}
	}
	return result
}
