package Solution

import (
	"sort"
	"strconv"
	"strings"
)

type Event struct {
	Type      string
	Timestamp int
	All, Here bool
	Users     []int
}

func Solution(numberOfUsers int, events [][]string) []int {
	cEvents := make([]Event, 0)
	var uid int
	for i := range events {
		t, _ := strconv.Atoi(events[i][1])
		e := Event{
			Type:      events[i][0],
			Timestamp: t,
		}
		if e.Type == "OFFLINE" {
			uid, _ = strconv.Atoi(events[i][2])
			e.Users = append(e.Users, uid)
			cEvents = append(cEvents, e)
			continue
		}

		if events[i][2] == "ALL" {
			e.All = true
		} else if events[i][2] == "HERE" {
			e.Here = true
		} else {
			users := strings.Split(events[i][2], " ")
			e.Users = make([]int, len(users))
			for i, u := range users {
				uid, _ = strconv.Atoi(u[2:])
				e.Users[i] = uid
			}
		}

		cEvents = append(cEvents, e)
	}
	sort.Slice(cEvents, func(i, j int) bool {
		a, b := cEvents[i], cEvents[j]
		if a.Timestamp == b.Timestamp {
			return a.Type == "OFFLINE"
		}
		return a.Timestamp < b.Timestamp
	})
	onlineTime := make([]int, numberOfUsers)
	ret := make([]int, numberOfUsers)

	for _, e := range cEvents {
		if e.Type == "MESSAGE" {
			if e.All {
				for i := range ret {
					ret[i]++
				}
			} else if e.Here {
				for i := range ret {
					if e.Timestamp < onlineTime[i] {
						continue
					}
					ret[i]++
				}
			} else {
				for _, u := range e.Users {
					ret[u]++
				}
			}
			continue
		}

		onlineTime[e.Users[0]] = e.Timestamp + 60
	}

	return ret
}
