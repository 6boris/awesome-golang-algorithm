package leetcode

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestLeekCode_Unite(t *testing.T) {
	t.Run("Gen tpl file", func(t *testing.T) {
		resp, err := req.Get("https://leetcode.com/api/problems/Algorithms/")
		assert.Nil(t, err)
		algorithmData := &AlgorithmsData{}
		err = json.Unmarshal(resp.Bytes(), algorithmData)
		assert.Nil(t, err)
		for _, v := range algorithmData.StatStatusPairs {
			id := v.Stat.FrontendQuestionID
			//if id <= 2820 {
			//	continue
			//}
			log.Printf("~~ 开始生成第 %d 题的文件夹 ~~\n", id)
			if v.Stat.QuestionID == 0 {
				log.Printf("%d 号题不存，请核查题号。\n", id)
				continue
			}
			//if v.PaidOnly {
			//	log.Printf("%d 号题需要付费。如果已经订阅，请注释掉本代码。\n", id)
			//	continue
			//}
			// 复制 Go 文件
			sourceDir := "../template/solution/"
			destinationDir := v.GetDir()

			if _, err := os.Stat(destinationDir); !os.IsNotExist(err) {
				log.Printf("目录已经存在不处理: %s \n", destinationDir)
				continue
			}

			_ = createDirectory(destinationDir)
			err = copyFile(sourceDir+"Solution.go", destinationDir+"/Solution.go", os.ModePerm)
			assert.Nil(t, err)
			err = copyFile(sourceDir+"Solution_test.go", destinationDir+"/Solution_test.go", os.ModePerm)
			assert.Nil(t, err)
			//err = GenerateReadmeFromTpl(v)
			//assert.Nil(t, err)
		}
	})
	t.Run("copyFile", func(t *testing.T) {
		sourceFile := "../template/solution/Solution.go"
		destinationFile := "./tmp/dst.txt"
		path := "tmp"
		data, err := os.ReadFile(sourceFile)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(path, os.ModePerm)
			if err != nil {
				log.Println(err)
			}
		}
		err = os.WriteFile(destinationFile, data, 0755)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		fmt.Println("File copied successfully")
	})
}
