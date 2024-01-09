package Solution

import "fmt"

func Solution(cpdomains []string) []string {
	domains := make(map[string]int)
	var (
		rep    int
		domain string
	)
	for _, cnt := range cpdomains {
		fmt.Sscanf(cnt, "%d %s", &rep, &domain)
		domains[domain] += rep
		for idx := len(domain) - 1; idx >= 0; idx-- {
			if domain[idx] == '.' {
				x := domain[idx+1:]
				domains[x] += rep
			}
		}
	}
	ans := make([]string, 0)
	for domain, cnt := range domains {
		ans = append(ans, fmt.Sprintf("%d %s", cnt, domain))
	}
	return ans
}
