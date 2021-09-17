package leetcode

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

const SOURCE_SOLUTION_SUMMARY_FILE_PATH = "cmd/template/gitbook/SUMMARY.md"

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

//	Auto make the Gitbook SUMMARY
func MakeGitbookSummary(problems []Problem) {
	problems = CheckProblemExists(problems)

	file, err := os.OpenFile(SOURCE_SOLUTION_SUMMARY_FILE_PATH, os.O_RDONLY, 0o600)
	defer file.Close()
	if err != nil {
		log.Panicf("README 模板读取失败1：%s", err.Error())
	}

	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		log.Panicf("README 模板读取失败2：%s", err.Error())
	}

	var tmpRes bytes.Buffer

	tmpl, err := template.New("SUMMARY: ").Parse(string(buffer))
	err = tmpl.Execute(&tmpRes, problems)
	write("SUMMARY.md", string(tmpRes.Bytes()))
}

func CheckProblemExists(problems []Problem) []Problem {
	tmp := []Problem{}

	for i := 0; i < len(problems); i++ {
		isExist, _ := PathExists("leetcode/" + problems[i].PathName)
		if isExist {
			tmp = append(tmp, problems[i])
		} else {
			fmt.Println(problems[i].PathName)
		}
	}
	return tmp
}

func write(path, content string) {
	err := ioutil.WriteFile(path, []byte(content), 0o755)
	if err != nil {
		log.Fatal(err)
	}
}

func readTMPL(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func getSUMMARYBuffer(filepath string) {
}
