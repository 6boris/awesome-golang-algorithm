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

//	生成木目录Dir
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
			err := os.Mkdir(SOLUTIONS_PATH+problems[i].PathName, os.ModePerm)
			if err != nil {
				log.Printf("目录创建失败:%s", err.Error())
			} else {
				//	复制文件
				log.Println("开始复制文件：")
				copy(SOURCE_SOLUTION_FILE_PATH, SOLUTIONS_PATH+problems[i].PathName+"/Solution.go")
				copy(SOURCE_SOLUTION_TEST_FILE_PATH, SOLUTIONS_PATH+problems[i].PathName+"/Solution_test.go")
				GenerateReadme(problems[i])
			}
		}

	}
}

// 拷贝文件  要拷贝的文件路径 拷贝到哪里
func copyFile(source, dest string) bool {
	if source == "" || dest == "" {
		log.Println("source or dest is null")
		return false
	}
	// 打开文件资源
	source_open, err := os.Open(source)
	// 养成好习惯。操作文件时候记得添加 defer 关闭文件资源代码
	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer source_open.Close()
	// 只写模式打开文件 如果文件不存在进行创建 并赋予 644的权限。详情查看linux 权限解释
	dest_open, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 644)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	// 养成好习惯。操作文件时候记得添加 defer 关闭文件资源代码
	defer dest_open.Close()
	// 进行数据拷贝
	_, copy_err := io.Copy(dest_open, source_open)
	if copy_err != nil {
		log.Println(copy_err.Error())
		return false
	} else {
		return true
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

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
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
