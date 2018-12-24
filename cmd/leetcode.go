package cmd

import (
	"encoding/json"
	"log"
	"net/http"
)

func getProblem() {
	URL := "https://leetcode.com/api/problems/Algorithms/"

	raw, err := http.Get(URL)

	res := new(algorithms)
	if err := json.Unmarshal(raw, res); err != nil {
		log.Panicf("无法把json转换成Category: %s\n", err.Error())
	}

	// 如果，没有登录的话，也能获取数据，但是用户名，就不是本人
	if res.User != getConfig().Username {
		log.Fatal("没有获取到本人的数据")
	}

	return res
}
