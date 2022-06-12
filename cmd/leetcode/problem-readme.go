package leetcode

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GetReadmeTemplateBuffer() string {
	data, err := ioutil.ReadFile(SOURCE_SOLUTION_SUMMARY_FILE_PATH)
	if err != nil {
		log.Panicln("读取 README模板失败", err.Error())
	}
	fmt.Println(string(data))
	return string(data)
}

func GenerateReadme(problem Problem, basePath ...string) {
	log.Println("开始生成 README")
	file, err := os.OpenFile(SOURCE_SOLUTION_README_FILE_PATH, os.O_RDONLY, 0o600)
	defer file.Close()
	if err != nil {
		log.Panicf("README 模板读取失败1：%s", err.Error())
	}

	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		log.Panicf("README 模板读取失败2：%s", err.Error())
	}

	var tmpRes bytes.Buffer

	tmpl, err := template.New("README: ").Parse(string(buffer))
	err = tmpl.Execute(&tmpRes, problem)
	if len(basePath) > 0 {
		dir := strings.Join(basePath, "/")
		dir = strings.TrimSuffix(dir, "/")
		write(dir+"/README.md", string(tmpRes.Bytes()))
		return
	}
	write(SOLUTIONS_PATH+problem.PathName+"/README.md", string(tmpRes.Bytes()))
}
