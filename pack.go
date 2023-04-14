package hostutils

import (
	"fmt"
	"sort"
	"strings"
)

// PackString Pack space separated full hosts list into short abbreviated hosts.
func PackString(hosts string) (packedHosts []string) {
	return Pack([]string{hosts})
}

// Pack full hosts list into short abbreviated hosts.
func Pack(hosts []string) (packedHosts []string) {
	regHosts := regularizeHosts(hosts[:])
	if regHosts == nil {
		return nil
	}
	if len(regHosts) == 0 {
		return []string{}
	}
	return packHosts(regHosts)
}

func packHosts(uniqHosts []string) []string {
	hostGroups := make(map[string][]string)
	var result []string
	for _, fqdn := range uniqHosts {
		hostname, domain := parseFQDN(fqdn)
		m := reHostname.FindStringSubmatch(hostname)
		if len(m) == 0 {
			result = append(result, fqdn)
		} else {
			prefix := m[1]
			num := m[2]
			suffix := m[3]
			key := fmt.Sprintf("%s%%s%s%s", prefix, suffix, domain)
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
