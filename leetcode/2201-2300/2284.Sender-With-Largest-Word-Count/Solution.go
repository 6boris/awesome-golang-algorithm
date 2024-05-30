package Solution

import "strings"

func Solution(messages []string, senders []string) string {
	count := make(map[string]int)
	mc, ma := 0, ""
	for idx, msg := range messages {
		count[senders[idx]] += len(strings.Split(msg, " "))
		if count[senders[idx]] > mc {
			mc = count[senders[idx]]
			ma = senders[idx]
			continue
		}
		if count[senders[idx]] == mc && senders[idx] > ma {
			ma = senders[idx]
		}
	}
	return ma
}
