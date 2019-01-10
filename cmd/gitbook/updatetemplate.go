package gitbook

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

//	Gitbook summary template file path
const GITBOOK_SUMMARY_TEMPLATE_FILE_PATH = "../template/gitbook/SUMMARY.md"

//	更新GitBook SUMMARY
func UpdateGitBookSummary() {
	summarybuffer := readSummaryFileBuffer()


	t := template.Must(template.New("letter").Parse(string(summarybuffer)))
	fmt.Println(t)

	problems := leetcode.


	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}

//	获取 Summary file buffer
func readSummaryFileBuffer() []byte {
	fileBuffer, err := ioutil.ReadFile(GITBOOK_SUMMARY_TEMPLATE_FILE_PATH)
	if err != nil {
		log.Panicln("读取Gitbook SUMMARY 文件失败：", err.Error())
	}
	return fileBuffer
}
