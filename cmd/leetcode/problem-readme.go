package leetcode

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"strings"
)

func GetReadmeTemplateBuffer() string {
	data, err := os.ReadFile(SOURCE_SOLUTION_SUMMARY_FILE_PATH)
	if err != nil {
		log.Panicln("读取 README模板失败", err.Error())
	}
	fmt.Println(string(data))
	return string(data)
}

func GenerateReadme(problem Problem, basePath ...string) {
	log.Println("开始生成 README")
	file, err := os.OpenFile(SOURCE_SOLUTION_README_FILE_PATH, os.O_RDONLY, 0o600)
	if err != nil {
		log.Panicf("README 模板读取失败1：%s", err.Error())
	}
	defer func() { _ = file.Close() }()

	buffer, err := io.ReadAll(file)
	if err != nil {
		log.Panicf("README 模板读取失败2：%s", err.Error())
	}

	var tmpRes bytes.Buffer

	tmpl, err := template.New("README: ").Parse(string(buffer))
	if err != nil {
		log.Panicf("README 模板解析失败：%s", err.Error())
	}
	if err = tmpl.Execute(&tmpRes, problem); err != nil {
		log.Panicf("README 模板渲染失败：%s", err.Error())
	}
	if len(basePath) > 0 {
		dir := strings.Join(basePath, "/")
		dir = strings.TrimSuffix(dir, "/")
		write(dir+"/README.md", tmpRes.String())
		return
	}
	write(SOLUTIONS_PATH+problem.PathName+"/README.md", tmpRes.String())
}
