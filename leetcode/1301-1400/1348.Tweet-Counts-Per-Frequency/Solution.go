package Solution

import "sort"

type TweetCounts struct {
	data map[string][]int
}

func Constructor() TweetCounts {
	t := TweetCounts{
		data: make(map[string][]int),
	}
	return t
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	this.data[tweetName] = append(this.data[tweetName], time)
	sort.Ints(this.data[tweetName])
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	interval := 60
	if freq == "hour" {
		interval = 3600
	}
	if freq == "day" {
		interval = 86400
	}

	s := startTime
	ans := make([]int, 0)
	data := this.data[tweetName]
	for s <= endTime {
		e := s + interval - 1
		e = min(endTime, e)
		l := sort.Search(len(data), func(i int) bool {
			return data[i] >= s
		})
		c := 0
		if l != len(data) {
			r := sort.Search(len(data), func(i int) bool {
				return data[i] > e
			})
			c = r - l
		}
		ans = append(ans, c)
		s = s + interval
	}
	return ans
}

type opt struct {
	name  string
	tname string
	cname string

	ttime int
	s, e  int
}

func Solution(opts []opt) [][]int {
	ans := make([][]int, 0)
	c := Constructor()
	for _, o := range opts {
		if o.name == "r" {
			c.RecordTweet(o.tname, o.ttime)
			continue
		}
		ans = append(ans, c.GetTweetCountsPerFrequency(o.cname, o.tname, o.s, o.e))
	}
	return ans
}
