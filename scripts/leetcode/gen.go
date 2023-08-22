package leetcode

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

type AlgorithmsData struct {
	UserName        string            `json:"user_name"`
	NumSolved       int               `json:"num_solved"`
	NumTotal        int               `json:"num_total"`
	AcEasy          int               `json:"ac_easy"`
	AcMedium        int               `json:"ac_medium"`
	AcHard          int               `json:"ac_hard"`
	StatStatusPairs []*AlgorithmsPair `json:"stat_status_pairs"`
}
type AlgorithmsPair struct {
	Stat       AlgorithmsPairStat `json:"stat"`
	Status     any                `json:"status"`
	Difficulty struct {
		Level int `json:"level"`
	} `json:"difficulty"`
	PaidOnly  bool `json:"paid_only"`
	IsFavor   bool `json:"is_favor"`
	Frequency int  `json:"frequency"`
	Progress  int  `json:"progress"`
}

func (a *AlgorithmsPair) GetFormatID() string {
	strId := strconv.Itoa(a.Stat.FrontendQuestionID)

	if len(strId) == 1 {
		strId = "000" + strId
	}
	if len(strId) == 2 {
		strId = "00" + strId
	}
	if len(strId) == 3 {
		strId = "0" + strId
	}
	return strId
}
func (a *AlgorithmsPair) GetFormatName() string {
	name := strings.Trim(a.Stat.QuestionTitle, " ")
	str := ""
	for i, v := range name {
		if v == ' ' && name[i-1] != '-' {
			str = str + "-"
			continue
		}
		if v == ' ' {
			continue
		}
		if v == '`' {
			str = str + "-"
			continue
		}
		if v == '%' {
			str = str + "p"
			continue
		}

		if v == '(' && name[i-1] != '-' {
			str = str + "-"
			continue
		}

		if v == ')' {
			continue
		}
		if v == ',' {
			continue
		}
		// ' 过滤
		if v == 39 {
			continue
		}
		if v == '?' {
			continue
		}
		if i > 0 && name[i-1] == '-' {
			str = str + string(v)
			continue
		}
		str = str + string(v)
	}
	if name[len(name)-1:] == "-" {
		name = name[:len(name)-1]
	}
	return str
}
func (a *AlgorithmsPair) GetDir() string {
	return fmt.Sprintf("../../leetcode/%d-%d/%s/", (a.Stat.FrontendQuestionID-1)/100*100+1, ((a.Stat.FrontendQuestionID-1)/100+1)*100, fmt.Sprintf("%s.%s", a.GetFormatID(), a.GetFormatName()))
}

type AlgorithmsPairStat struct {
	QuestionID                      int    `json:"question_id"`
	QuestionArticleLive             any    `json:"question__article__live"`
	QuestionArticleSlug             any    `json:"question__article__slug"`
	QuestionArticleHasVideoSolution any    `json:"question__article__has_video_solution"`
	QuestionTitle                   string `json:"question__title"`
	QuestionTitleSlug               string `json:"question__title_slug"`
	QuestionHide                    bool   `json:"question__hide"`
	TotalAcs                        int    `json:"total_acs"`
	TotalSubmitted                  int    `json:"total_submitted"`
	FrontendQuestionID              int    `json:"frontend_question_id"`
	IsNewQuestion                   bool   `json:"is_new_question"`
}

func createDirectory(path string) error {
	// Check if the directory exists
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	// If the directory does not exist, create its parent
	if os.IsNotExist(err) {
		err = createDirectory(filepath.Dir(path))
		if err != nil {
			return err
		}
		// Create the directory
		err = os.Mkdir(path, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

func copyFile(sourceFile, destinationFile string, perm os.FileMode) error {
	if sourceFile == "" || destinationFile == "" {
		return errors.New("sourceFile or destinationFile is null")
	}
	data, err := os.ReadFile(sourceFile)
	if err != nil {
		return err
	}
	err = os.WriteFile(destinationFile, data, perm)
	if err != nil {
		return err
	}
	return nil
}

func GenerateReadmeFromTpl(problem *AlgorithmsPair) error {
	log.Println("开始生成 README")
	file, err := os.OpenFile("../template/solution/README.md", os.O_RDONLY, 0o600)
	defer file.Close()
	if err != nil {
		log.Panicf("README 模板读取失败1：%s", err.Error())
		return err
	}

	buffer, err := io.ReadAll(file)
	if err != nil {
		log.Panicf("README 模板读取失败2：%s", err.Error())
		return err
	}

	var tmpRes bytes.Buffer

	tmpl, err := template.New("README: ").Parse(string(buffer))
	err = tmpl.Execute(&tmpRes, problem)
	//write(SOLUTIONS_PATH+problem.PathName+"/README.md", string(tmpRes.Bytes()))
	err = os.WriteFile(problem.GetDir()+"/README.md", tmpRes.Bytes(), 0o755)
	if err != nil {
		return err
	}
	return nil
}
