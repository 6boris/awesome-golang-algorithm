package leetcode

import (
	"fmt"
	"io/ioutil"
	"log"
)

func GetReadmeTemplateBuffer() string {
	data, err := ioutil.ReadFile(SOURCE_SOLUTION_README_FILE_PATH)
	if err != nil {
		log.Panicln("读取 README模板失败", err.Error())
	}
	fmt.Println(string(data))
	return string(data)
}
