package leetcode

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	SOLUTIONS_PATH                   = "leetcode/"
	SOURCE_SOLUTION_FILE_PATH        = "cmd/template/solution/Solution.go"
	SOURCE_SOLUTION_TEST_FILE_PATH   = "cmd/template/solution/Solution_test.go"
	SOURCE_SOLUTION_README_FILE_PATH = "cmd/template/solution/README.md"
)

func MakeDirFromTo(problems []Problem, from, to int) {
	if from%100 != 1 || to-from+1 != 100 {
		return
	}
	baseDir := fmt.Sprintf("%d-%d/", from, to)
	for i := 0; i < len(problems); i++ {
		if problems[i].Stat.FrontendQuestionID >= from && problems[i].Stat.FrontendQuestionID <= to {
			log.Printf("~~ 开始生成第 %d 题的文件夹 ~~", problems[i].Stat.FrontendQuestionID)

			if problems[i].PaidOnly {
				log.Printf("%d 号题需要付费。如果已经订阅，请注释掉本代码。", problems[i].Stat.FrontendQuestionID)
				continue
			}
			_path := SOLUTIONS_PATH + baseDir + problems[i].PathName
			if is_DirExists, _ := PathExists(_path); is_DirExists {
				log.Println("目录已经存在：", _path)
			} else {
				err := os.MkdirAll(_path, 0o750)
				if err != nil {
					log.Printf("目录创建失败:%s", err.Error())
				} else {
					//	复制文件
					log.Println("开始复制文件：")
					if _, err := copy(SOURCE_SOLUTION_FILE_PATH, _path+"/Solution.go"); err != nil {
						log.Printf("复制文件失败:%s", err.Error())
					}
					if _, err := copy(SOURCE_SOLUTION_TEST_FILE_PATH, _path+"/Solution_test.go"); err != nil {
						log.Printf("复制文件失败:%s", err.Error())
					}
					GenerateReadme(problems[i], _path)
				}
			}
		}
	}
}

// 生成木目录Dir
func MakeDir(problems []Problem) {
	for i := 0; i < len(problems); i++ {
		// fmt.Println(problems[i].PathName)
		log.Printf("~~ 开始生成第 %d 题的文件夹 ~~", problems[i].Stat.FrontendQuestionID)

		//	检查数据
		if problems[i].Stat.QuestionID == 0 {
			log.Printf("%d 号题不存，请核查题号。", problems[i].Stat.FrontendQuestionID)
		}

		if problems[i].PaidOnly {
			log.Printf("%d 号题需要付费。如果已经订阅，请注释掉本代码。", problems[i].Stat.FrontendQuestionID)
			continue
		}

		if is_DirExists, _ := PathExists(SOLUTIONS_PATH + problems[i].PathName); is_DirExists {
			log.Println("目录已经存在：", SOLUTIONS_PATH+problems[i].PathName)
		} else {
			err := os.Mkdir(SOLUTIONS_PATH+problems[i].PathName, 0o750)
			if err != nil {
				log.Printf("目录创建失败:%s", err.Error())
			} else {
				//	复制文件
				log.Println("开始复制文件：")
				if _, err := copy(SOURCE_SOLUTION_FILE_PATH, SOLUTIONS_PATH+problems[i].PathName+"/Solution.go"); err != nil {
					log.Printf("复制文件失败:%s", err.Error())
				}
				if _, err := copy(SOURCE_SOLUTION_TEST_FILE_PATH, SOLUTIONS_PATH+problems[i].PathName+"/Solution_test.go"); err != nil {
					log.Printf("复制文件失败:%s", err.Error())
				}
				GenerateReadme(problems[i])
			}
		}

	}
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src) // #nosec G304 -- local dev tool copies template files by path
	if err != nil {
		return 0, err
	}
	defer func() { _ = source.Close() }()

	destination, err := os.Create(dst) // #nosec G304 -- local dev tool writes generated files by path
	if err != nil {
		return 0, err
	}
	defer func() { _ = destination.Close() }()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
