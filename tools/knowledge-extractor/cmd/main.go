package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/learninfun/learn-with-ai/internal/io"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"

	openai "github.com/sashabaranov/go-openai"
)

var (
	command              string
	workzone             string
	sourseFolder         string
	resultFolder         string
	cacheFolder          string
	pathMK               string
	lang                 string
	sourseFolderLen      int
	errorNum             int
	questionInfoSlice    []QuestionInfo
	answerOutputTemplate *template.Template
	translateMap         map[string]interface{}
)

type QuestionInfo struct {
	desc     string
	cacheKey string
	template string
}

type AnswerInfo struct {
	QuestionDesc string
	Question     string
	Answer       string
}

func main() {
	initGlobalVar()
	if command == "fetch" || command == "regen" {
		initTemplate()
		initQuestion()
		initTranslation()
		if command == "fetch" {
			initRole()
		}

		filepath.Walk(sourseFolder, walkAllFile)
	} else if command == "mdToFile" {
		mdListToTreeFiles()
	}

}

func initGlobalVar() {
	var err error

	pathMK = string(os.PathSeparator)

	currentFolder, err := os.Getwd()
	checkErr(err)
	fmt.Println(currentFolder)

	//lang = "zh-tw"
	if len(os.Args) >= 4 {
		command = os.Args[1]
		lang = os.Args[2]
		workzone = os.Args[3]
	} else {
		command = "regen"
		lang = "zh-cn"
		workzone = currentFolder + pathMK + "workzone"
	}

	sourseFolder = workzone + pathMK + "result" + pathMK + lang
	sourseFolderLen = len(sourseFolder)
	resultFolder = workzone + pathMK + "result" + pathMK + lang
	cacheFolder = workzone + pathMK + "cache" + pathMK + lang

	errorNum = 0 //init no error
}

func initQuestion() {
	if lang == "english" {
		var q1 QuestionInfo
		q1.desc = "Preview"
		q1.cacheKey = "preview"
		q1.template = "give me five question about %s"
		questionInfoSlice = append(questionInfoSlice, q1)

		var q2 QuestionInfo
		q2.desc = "Explain"
		q2.cacheKey = "explain"
		q2.template = "Explain %s and give an example"
		questionInfoSlice = append(questionInfoSlice, q2)

		var q3 QuestionInfo
		q3.desc = "Keypoint"
		q3.cacheKey = "keypoint"
		q3.template = "List the key points of %s"
		questionInfoSlice = append(questionInfoSlice, q3)

		var q4 QuestionInfo
		q4.desc = "Review"
		q4.cacheKey = "test"
		q4.template = "Give me 5 medium-difficulty questions with answers about %s"
		questionInfoSlice = append(questionInfoSlice, q4)

		var q5 QuestionInfo
		q5.desc = "Related webpage"
		q5.cacheKey = "ref"
		q5.template = "List the relevant introduction webpages about %s"
		questionInfoSlice = append(questionInfoSlice, q5)
	} else if lang == "zh-cn" {
		var q1 QuestionInfo
		q1.desc = "习题预习"
		q1.cacheKey = "preview"
		q1.template = "给我5题%s的问题"
		questionInfoSlice = append(questionInfoSlice, q1)

		var q2 QuestionInfo
		q2.desc = "说明知识"
		q2.cacheKey = "explain"
		q2.template = "说明%s并举例"
		questionInfoSlice = append(questionInfoSlice, q2)

		var q3 QuestionInfo
		q3.desc = "汇总重点"
		q3.cacheKey = "keypoint"
		q3.template = "条列%s的重点"
		questionInfoSlice = append(questionInfoSlice, q3)

		var q4 QuestionInfo
		q4.desc = "知识测验"
		q4.cacheKey = "test"
		q4.template = "给我5题%s的中等难度问题，并在后面列出答案"
		questionInfoSlice = append(questionInfoSlice, q4)

		var q5 QuestionInfo
		q5.desc = "网络数据"
		q5.cacheKey = "ref"
		q5.template = "给我5篇%s的网络数据"
		questionInfoSlice = append(questionInfoSlice, q5)
	} else if lang == "zh-tw" {
		var q1 QuestionInfo
		q1.desc = "習題預習"
		q1.cacheKey = "preview"
		q1.template = "給我5題%s的中文問題"
		questionInfoSlice = append(questionInfoSlice, q1)

		var q2 QuestionInfo
		q2.desc = "說明知識"
		q2.cacheKey = "explain"
		q2.template = "以中文說明%s並舉例"
		questionInfoSlice = append(questionInfoSlice, q2)

		var q3 QuestionInfo
		q3.desc = "彙總重點"
		q3.cacheKey = "keypoint"
		q3.template = "以中文條列%s的重點"
		questionInfoSlice = append(questionInfoSlice, q3)

		var q4 QuestionInfo
		q4.desc = "知識測驗"
		q4.cacheKey = "test"
		q4.template = "以中文給我5題%s的中等難度問題，並在後面列出答案"
		questionInfoSlice = append(questionInfoSlice, q4)

		var q5 QuestionInfo
		q5.desc = "網路資料"
		q5.cacheKey = "ref"
		q5.template = "給我5篇%s的中文網路資料"
		questionInfoSlice = append(questionInfoSlice, q5)
	}
}

func initTemplate() {
	templateStr := `## {{ .QuestionDesc}}
{{"{{<"}} ask_chatgpt {{">}}"}}
{{ .Question}}
{{"{{<"}} /ask_chatgpt {{">}}"}}

{{ .Answer}}   

`

	var err error
	answerOutputTemplate, err = template.New("answerOutput").Parse(templateStr)
	checkErr(err)
}

func initRole() {
	if lang == "english" {
		askChatGpt("Assuming you are an IT expert, answer my questions")
	} else if lang == "zh-tw" {
		askChatGpt("假設你是投資專家，請以中文回答我相關問題")
	}
}

func initTranslation() {
	if lang == "english" {
		return
	}

	translationPath := workzone + pathMK + "translation" + pathMK + lang + ".json"
	translationStr := io.FileToString(translationPath)
	err := json.Unmarshal([]byte(translationStr), &translateMap)
	checkErr(err)

}

func walkAllFile(filePath string, fileInfo os.FileInfo, err error) error {
	checkErr(err)

	var pathRelative = filePath[sourseFolderLen:]
	var pathResult = resultFolder + pathRelative
	if fileInfo.IsDir() {
		fmt.Printf("Folder Name: %s\n", pathRelative)
		if !io.FileExists(pathResult) {
			err := os.Mkdir(pathResult, 0755)
			checkErr(err)
		}
	} else {
		if fileInfo.Size() == 0 {
			var pathCache = cacheFolder + pathRelative
			extractKnowledge(fileInfo.Name(), pathResult, pathCache)
		}
	}

	return nil
}

func extractKnowledge(fileName, pathResult, pathCache string) error {
	if strings.ToLower(filepath.Ext(fileName)) != ".md" { //only handle markdown
		return nil
	}

	var err error
	fmt.Println("extractKnowledge: " + fileName)
	keyWord := io.FileNameNoExt(fileName)

	f, err := os.OpenFile(pathResult, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(err)

	if lang != "english" {
		keyWordTranslated := translateMap[keyWord]
		if keyWordTranslated != nil {
			f.WriteString(fmt.Sprintf("+++\ntitle = \"%s\"\n+++\n", keyWordTranslated))
		}
	}

	for _, questionInfo := range questionInfoSlice {
		var answerInfo AnswerInfo
		answerInfo.QuestionDesc = questionInfo.desc
		answerInfo.Question = fmt.Sprintf(questionInfo.template, keyWord)
		folderCache := pathCache[:len(pathCache)-len(fileName)]
		//fmt.Println("pathResult=" + pathResult)
		//fmt.Println("folderCache=" + folderCache)
		thisPathCache := pathCache[:len(pathCache)-len(fileName)] + keyWord + "_" + questionInfo.cacheKey + ".md"
		//fmt.Println("thisPathCache=" + thisPathCache)

		if io.FileExists(thisPathCache) {
			answerInfo.Answer = io.FileToString(thisPathCache)
		} else {
			answerInfo.Answer = strings.Trim(askChatGpt(answerInfo.Question), " ")

			if answerInfo.Answer != "" {
				errorNum = 0 //get answer reset errorNum

				if !io.FileExists(folderCache) {
					os.MkdirAll(folderCache, 0755)
				}
				io.StringToFile(thisPathCache, answerInfo.Answer)
			} else {
				fmt.Println("No answer=" + keyWord)
				errorNum++

				if errorNum > 10 {
					err = errors.New("Too many empty answer:" + string(errorNum))
					panic(err)
				}
			}
		}

		err = answerOutputTemplate.Execute(f, answerInfo)
		checkErr(err)
	}

	defer f.Close()

	return nil
}

func askChatGpt2(question string) string {
	return "abc"
}

func askChatGpt(question string) string {
	fmt.Println("askChatGpt: " + question)

	for i := 0; i < 3; i++ {
		client := openai.NewClient("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")
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo0301,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: question,
					},
				},
			},
		)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}

		return resp.Choices[0].Message.Content
	}

	return ""
}

func mdListToTreeFiles() {
	inputFilePath := workzone + "\\input.md"
	fmt.Println(inputFilePath)
	inputStr := io.FileToString(inputFilePath)
	fmt.Println(inputStr)

	source := []byte(inputStr)
	reader := text.NewReader(source)

	mdParser := goldmark.DefaultParser()
	node := mdParser.Parse(reader)

	// Traverse the Markdown AST and find all list nodes
	// Traverse the Markdown AST and find all list nodes
	resultPath := resultFolder
	Walk(node, func(node ast.Node, entering bool, idx int) (ast.WalkStatus, error) {
		if listItem, ok := node.(*ast.ListItem); ok {
			listText := string(listItem.FirstChild().Text(source))

			if entering {
				resultPath += "\\" + listText
				//fmt.Printf("- %s [%d]\n", resultPath, node.ChildCount())

				if node.ChildCount() == 1 { //leaf
					filePath := resultPath + ".md"
					fmt.Println("mkobj " + strconv.Itoa(idx) + ": " + filePath)
					if !io.FileExists(filePath) {
						io.StringToFile(filePath, "")
					}
				} else { //node
					filePath := resultPath + "\\" + listText + ".md"

					fmt.Println("mkdir: " + resultPath)
					fmt.Println("mkobj " + strconv.Itoa(idx) + ": " + filePath)

					if !io.FileExists(resultPath) {
						err := os.Mkdir(resultPath, 0755)
						checkErr(err)
					}

					if !io.FileExists(filePath) {
						io.StringToFile(filePath, "")
					}
				}

			} else {
				resultPath = resultPath[:len(resultPath)-1-len(listText)]
			}
		}

		return ast.WalkContinue, nil
	}, 0)
}

func checkErr(e error) {
	if e != nil {
		fmt.Printf("error: %w", e)
		panic(e)
	}
}

type Walker func(n ast.Node, entering bool, idx int) (ast.WalkStatus, error)

// Walk walks a AST tree by the depth first search algorithm.
func Walk(n ast.Node, walker Walker, idx int) error {
	_, err := walkHelper(n, walker, idx)
	return err
}

func walkHelper(n ast.Node, walker Walker, idx int) (ast.WalkStatus, error) {
	status, err := walker(n, true, idx)
	if err != nil || status == ast.WalkStop {
		return status, err
	}
	if status != ast.WalkSkipChildren {
		var lIdx = 1
		for c := n.FirstChild(); c != nil; c = c.NextSibling() {
			if st, err := walkHelper(c, walker, lIdx); err != nil || st == ast.WalkStop {
				return ast.WalkStop, err
			}

			lIdx++
		}
	}
	status, err = walker(n, false, idx)
	if err != nil || status == ast.WalkStop {
		return ast.WalkStop, err
	}
	return ast.WalkContinue, nil
}
