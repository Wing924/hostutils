package hostutils

import (
	"fmt"
	"sort"
	"strings"
)

// Pack Pack full hosts list into short abbreviated hosts.
func Pack(hosts []string) (packedHosts []string) {
	if hosts == nil {
		return nil
	}
	regHosts := regularizeHosts(hosts[:])
	if len(regHosts) == 0 {
		return []string{}
	}
	return packHosts(regHosts)
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

func packHosts(uniqHosts []string) []string {
	hostGroups := make(map[string][]string)
	var result []string
	for _, host := range uniqHosts {
		m := reHostname.FindStringSubmatch(host)
		if len(m) == 0 {
			result = append(result, host)
		} else {
			prefix := m[1]
			num := m[2]
			suffix := m[3]
			key := fmt.Sprintf("%s%%s%s", prefix, suffix)
			hostGroups[key] = append(hostGroups[key], num)
		}
	}

	for format, nums := range hostGroups {
		result = append(result, fmt.Sprintf(format, groupNums(nums)))
	}
	sort.Strings(result)
	return result
}

func groupNums(nums []string) string {
	if len(nums) == 0 {
		return ""
	}

	sort.Slice(nums, func(i, j int) bool {
		if len(nums[i]) != len(nums[j]) {
			return len(nums[i]) < len(nums[j])
		}
		return atoi(nums[i]) < atoi(nums[j])
	})
	nums = append(nums, "999999")

	minStr := nums[0]
	min := atoi(minStr)
	prevStr := minStr
	prev := min

	packedNums := []string{}
	for _, v := range nums {
		n := atoi(v)
		if n-prev > 1 || len(v) != len(prevStr) {
			if min == prev {
				packedNums = append(packedNums, minStr)
			} else {
				packedNums = append(packedNums, fmt.Sprintf("%s-%s", minStr, prevStr))
			}
			min = n
			minStr = v
		}
		prev = n
		prevStr = v
	}
	result := strings.Join(packedNums, ",")
	if reIsNumber.MatchString(result) {
		return result
	}
	return fmt.Sprintf("[%s]", result)
}
