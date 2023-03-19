package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
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
	resultFolder         string
	cacheFolder          string
	pathMK               string
	lang                 string
	resultFolderLen      int
	errorNum             int
	questionInfoSlice    []QuestionInfo
	answerOutputTemplate *template.Template
	translateMap         map[string]interface{}
	openAIKey            string
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

type Data struct {
	Root []interface{} `yaml:"tree"`
}

func main() {
	initGlobalVar()
	if command == "fetch" || command == "regen" {
		initTemplate()
		initQuestion()
		initTranslation()
		if command == "fetch" {
			//initRole()
		}

		err := os.RemoveAll(resultFolder) //clear ori result
		checkErr(err)

		rootTreeNode := mdListToTreeNode()
		for childIdx, child := range rootTreeNode.Children {
			traverseTreeNode(child, resultFolder, childIdx)
		}

		//filepath.Walk(sourseFolder, walkAllFile)
	} else if command == "yamlListToFile" {
		//mdListToTreeFiles()
	}

}

func initGlobalVar() {
	var err error

	pathMK = string(os.PathSeparator)

	currentFolder, err := os.Getwd()
	checkErr(err)
	fmt.Println(currentFolder)

	//exePath, err := os.Executable()
	//if err != nil {
	//	panic(err)
	//}

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

	resultFolder = workzone + pathMK + "result" + pathMK + lang
	resultFolderLen = len(resultFolder)
	cacheFolder = workzone + pathMK + "cache" + pathMK + lang

	if command == "fetch" {
		openAIKey = io.FileToString(currentFolder + pathMK + ".." + pathMK + "config" + pathMK + "openAI_api_key.txt")
		if openAIKey == "" {
			panic(errors.New("openAI_api_key.txt not exist"))
		}
	}

	errorNum = 0 //init no error
}

type TreeNode struct {
	Name     string
	Parent   *TreeNode
	Children []*TreeNode
}

func mdListToTreeNode() *TreeNode {
	inputFilePath := workzone + "\\input.md"
	fmt.Println(inputFilePath)
	inputStr := io.FileToString(inputFilePath)
	fmt.Println(inputStr)

	source := []byte(inputStr)
	reader := text.NewReader(source)

	mdParser := goldmark.DefaultParser()
	node := mdParser.Parse(reader)

	rootTreeNode := &TreeNode{Name: "Dummy"}
	//child := &TreeNode{Name: "Child", Parent: rootNode}
	//rootNode.Children = append(rootNode.Children, child)

	currentTreeNode := rootTreeNode
	//level := 0
	// Traverse the Markdown AST and find all list nodes
	// Traverse the Markdown AST and find all list nodes
	//resultPath := resultFolder
	Walk(node, func(node ast.Node, entering bool, idx int) (ast.WalkStatus, error) {
		if listItem, ok := node.(*ast.ListItem); ok {
			listText := string(listItem.FirstChild().Text(source))
			if entering {
				fmt.Println("enter:" + listText + ", current:" + currentTreeNode.Name)
				childTreeNode := &TreeNode{Name: listText, Parent: currentTreeNode}
				currentTreeNode.Children = append(currentTreeNode.Children, childTreeNode)
				currentTreeNode = childTreeNode

			} else {
				//resultPath = resultPath[:len(resultPath)-1-len(listText)]
				fmt.Println("leave:" + listText + ", current:" + currentTreeNode.Name)
				currentTreeNode = currentTreeNode.Parent
			}
		}

		return ast.WalkContinue, nil
	}, 0)

	fmt.Println(currentTreeNode)

	return rootTreeNode
}

func traverseTreeNode(node *TreeNode, folderPath string, idx int) {
	var storeFolder string
	if node.Children == nil { //leaf
		storeFolder = folderPath
	} else { //folder
		storeFolder = folderPath + pathMK + node.Name
		fmt.Println("mkDir:" + storeFolder)
		err := os.MkdirAll(storeFolder, 0755)
		checkErr(err)

		for childIdx, child := range node.Children {
			traverseTreeNode(child, storeFolder, childIdx)
		}
	}

	var pathRelative = storeFolder[resultFolderLen:]
	var pathCache = cacheFolder + pathRelative
	extractKnowledge(node.Name, storeFolder, pathCache, idx)
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
	if io.FileExists(translationPath) {
		translationStr := io.FileToString(translationPath)
		err := json.Unmarshal([]byte(translationStr), &translateMap)
		checkErr(err)
	}
}

func extractKnowledge(keyWord, thisResultFolder, thisCacheFolder string, idx int) error {
	var err error
	fmt.Println("extractKnowledge: " + keyWord)

	pathResult := thisResultFolder + pathMK + keyWord + ".md"
	f, err := os.OpenFile(pathResult, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(err)

	var keyWordTranslated string
	if translateMap != nil {
		keyWordTranslated = translateMap[keyWord].(string)
	}

	if keyWordTranslated != "" {
		f.WriteString(fmt.Sprintf("+++\n"+
			"title = \"%s\"\n"+
			"weight = \"%d\"\n"+
			"+++\n", keyWordTranslated, idx+1))
	} else {
		f.WriteString(fmt.Sprintf("+++\n"+
			"weight = \"%d\"\n"+
			"+++\n", idx+1))
	}

	for _, questionInfo := range questionInfoSlice {
		var answerInfo AnswerInfo
		answerInfo.QuestionDesc = questionInfo.desc
		answerInfo.Question = fmt.Sprintf(questionInfo.template, keyWord)

		thisPathCache := thisCacheFolder + pathMK + keyWord + "_" + questionInfo.cacheKey + ".md"
		//fmt.Println("thisPathCache=" + thisPathCache)

		if io.FileExists(thisPathCache) {
			answerInfo.Answer = io.FileToString(thisPathCache)
		} else {
			answerInfo.Answer = strings.Trim(askChatGpt(answerInfo.Question), " ")

			if answerInfo.Answer != "" {
				errorNum = 0 //get answer reset errorNum
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
		client := openai.NewClient(openAIKey)
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

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
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
