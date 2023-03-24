package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"

	"github.com/learninfun/learn-with-ai/internal/io"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"

	"gopkg.in/yaml.v3"

	openai "github.com/sashabaranov/go-openai"
)

/*
var (
	command              string
	workzone             string
	outputFolder         string
	cacheFolder          string
	pathMK               string
	lang                 string
	outputFolderLen      int
	errorNum             int
	questionInfoSlice    []QuestionInfo
	answerOutputTemplate *template.Template
	translateMap         map[string]interface{}
	openAIKey            string
)
*/

var (
	parm      Parm
	globalVar GlobalVar
	config    Config
	pathMK    string
)

type GlobalVar struct {
	workzone             string
	errorNum             int
	inputFilePath        string
	configFilePath       string
	translateFilePath    string
	outputFolderPath     string
	outputFolderLen      int
	cacheFolderPath      string
	answerOutputTemplate *template.Template
	translateMap         map[string]interface{}
	questionVarMap       map[string]string
}

type Parm struct {
	fmFormat          string
	toFormat          string
	inputFilePath     string
	configFilePath    string
	translateFilePath string
	outputFolderPath  string
	cacheFolderPath   string
	apiKeyOpenAI      string
}

type Config struct {
	InitRole      string         `yaml:"initRole"`
	QuestionInfos []QuestionInfo `yaml:"questions"`
}

type QuestionInfo struct {
	Key         string `yaml:"key"`
	Desc        string `yaml:"desc"`
	TemplateStr string `yaml:"template"`
	TemplateObj *template.Template
	Include     []string `yaml:"include,omitempty"`
	Exclude     []string `yaml:"exclude,omitempty"`
	IncludeRe   []*regexp.Regexp
	ExcludeRe   []*regexp.Regexp
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
	initParm()
	initGlobalVar()
	initConfig()

	initTemplate()
	initTranslation()
	// initRole()

	err := os.RemoveAll(globalVar.outputFolderPath) //clear ori result
	checkErr(err)

	rootTreeNode := mdListToTreeNode()
	for childIdx, child := range rootTreeNode.Children {
		traverseTreeNode(child, globalVar.outputFolderPath, childIdx)
	}
}

func initParm() {
	flag.StringVar(&parm.fmFormat, "f", "", "From format")
	flag.StringVar(&parm.toFormat, "t", "markdownFolder", "To format")
	flag.StringVar(&parm.configFilePath, "c", "config.yaml", "Config file name")
	flag.StringVar(&parm.translateFilePath, "trans", "translation.json", "Translation file name")
	flag.StringVar(&parm.outputFolderPath, "of", "output", "Output folder path")
	flag.StringVar(&parm.cacheFolderPath, "cf", "cache", "Cache folder path")
	flag.StringVar(&parm.apiKeyOpenAI, "apikey-openai", "", "API key for OpenAI")

	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Input file name is required")
		return
	}
	parm.inputFilePath = flag.Arg(0)

	fmt.Printf("parm: %%+v: %+v\n", parm) // %+v: {name:Yuto age:35}
}

func initGlobalVar() {
	var err error

	pathMK = string(os.PathSeparator)

	currentFolder, err := os.Getwd()
	checkErr(err)
	fmt.Println(currentFolder)

	globalVar.workzone = currentFolder

	globalVar.inputFilePath = getAbsoluteFilePath(parm.inputFilePath)
	globalVar.configFilePath = getAbsoluteFilePath(parm.configFilePath)
	globalVar.translateFilePath = getAbsoluteFilePath(parm.translateFilePath)
	globalVar.outputFolderPath = getAbsoluteFilePath(parm.outputFolderPath)
	globalVar.cacheFolderPath = getAbsoluteFilePath(parm.cacheFolderPath)

	globalVar.outputFolderLen = len(globalVar.outputFolderPath)
	globalVar.questionVarMap = make(map[string]string)
}

func getAbsoluteFilePath(inputFile string) string {
	if path.IsAbs(inputFile) {
		return inputFile
	}
	return globalVar.workzone + pathMK + inputFile
}

func initConfig() {
	// read YAML file
	data, err := ioutil.ReadFile(globalVar.configFilePath)
	if err != nil {
		panic(err)
	}

	// parse YAML
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	//for i := range config.QuestionInfos {
	//	config.QuestionInfos[i].TemplateObj = template.New("questionTemplate")
	//}

	// print parsed config
	fmt.Printf("InitRole: %s\n", config.InitRole)
	for i, q := range config.QuestionInfos {
		fmt.Printf("Question %d:\n", i+1)
		fmt.Printf("  Key: %s\n", q.Key)
		fmt.Printf("  Desc: %s\n", q.Desc)
		fmt.Printf("  Template: %s\n", q.TemplateStr)

		var err error
		//q.TemplateObj, err = q.TemplateObj.Parse(q.TemplateStr)
		config.QuestionInfos[i].TemplateObj, err = template.New("questionTemplate").Parse(q.TemplateStr)
		checkErr(err)

		for j, includeRegex := range q.Include {
			if pathMK == "\\" {
				includeRegex = strings.ReplaceAll(includeRegex, "\\", "\\\\")
				includeRegex = strings.ReplaceAll(includeRegex, "/", "\\\\")
			}

			includeRegex := strings.ReplaceAll(includeRegex, ".", "\\.")
			includeRegex = strings.ReplaceAll(includeRegex, "**", "(.|\\n)*")
			includeRegex = strings.ReplaceAll(includeRegex, "*", "[^/\\n]*")
			includeRegex = "^" + includeRegex + "$"

			config.QuestionInfos[i].Include[j] = includeRegex
			re := regexp.MustCompile(includeRegex)
			config.QuestionInfos[i].IncludeRe = append(config.QuestionInfos[i].IncludeRe, re)
		}

		//if len(q.Include) == 0 {
		//	config.QuestionInfos[i].Include = append(q.Include, "*")
		//}

		fmt.Printf("  Include: %v\n", q.Include)

		if len(q.Exclude) > 0 {
			fmt.Printf("  Exclude: %v\n", q.Exclude)
		}

		for j, excludeRegex := range q.Exclude {
			if pathMK == "\\" {
				excludeRegex = strings.ReplaceAll(excludeRegex, "\\", "\\\\")
				excludeRegex = strings.ReplaceAll(excludeRegex, "/", "\\\\")
			}

			excludeRegex = strings.ReplaceAll(excludeRegex, ".", "\\.")
			excludeRegex = strings.ReplaceAll(excludeRegex, "**", "(.|\\n)*")
			excludeRegex = strings.ReplaceAll(excludeRegex, "*", "[^/\\n]*")
			excludeRegex = "^" + excludeRegex + "$"

			config.QuestionInfos[i].Exclude[j] = excludeRegex
			re := regexp.MustCompile(excludeRegex)
			config.QuestionInfos[i].ExcludeRe = append(config.QuestionInfos[i].ExcludeRe, re)
		}
	}

	fmt.Printf("config: %%+v: %+v\n", config)
}

type TreeNode struct {
	Name     string
	Parent   *TreeNode
	Children []*TreeNode
}

func mdListToTreeNode() *TreeNode {
	inputStr := io.FileToString(globalVar.inputFilePath)
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
	//resultPath := outputFolder
	Walk(node, func(node ast.Node, entering bool, idx int) (ast.WalkStatus, error) {
		if listItem, ok := node.(*ast.ListItem); ok {
			listText := string(listItem.FirstChild().Text(source))
			if entering {
				childTreeNode := &TreeNode{Name: listText, Parent: currentTreeNode}
				currentTreeNode.Children = append(currentTreeNode.Children, childTreeNode)
				currentTreeNode = childTreeNode

			} else {
				//resultPath = resultPath[:len(resultPath)-1-len(listText)]
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

	var pathRelative = storeFolder[globalVar.outputFolderLen:]
	var pathCache = globalVar.cacheFolderPath + pathRelative
	extractKnowledge(node.Name, pathRelative, storeFolder, pathCache, idx)
}

func initTemplate() {
	templateStr := `## {{ .QuestionDesc}}
{{"{{<"}} ask_chatgpt {{">}}"}}
{{ .Question}}
{{"{{<"}} /ask_chatgpt {{">}}"}}

{{ .Answer}}   

`

	var err error
	globalVar.answerOutputTemplate, err = template.New("answerOutput").Parse(templateStr)
	checkErr(err)
}

func initRole() {
	askChatGpt(config.InitRole)
}

func initTranslation() {
	if io.FileExists(globalVar.translateFilePath) {
		translationStr := io.FileToString(globalVar.translateFilePath)
		err := json.Unmarshal([]byte(translationStr), &globalVar.translateMap)
		checkErr(err)
	}
}

func extractKnowledge(keyword, pathRelative, thisResultFolder, thisCacheFolder string, idx int) error {
	var err error
	fmt.Println("extractKnowledge: " + keyword)

	pathResult := thisResultFolder + pathMK + keyword + ".md"
	f, err := os.OpenFile(pathResult, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(err)

	var keyWordTranslated string
	if globalVar.translateMap != nil {
		if translated, ok := globalVar.translateMap[keyword].(string); ok {
			keyWordTranslated = translated
		}
	}

	if keyWordTranslated != "" {
		f.WriteString(fmt.Sprintf("+++\n"+
			"title = \"%s\"\n"+
			"weight = \"%d\"\n"+
			"+++\n", keyWordTranslated, idx+1))
	} else {
		f.WriteString(fmt.Sprintf("+++\n"+
			"title = \"%s\"\n"+
			"weight = \"%d\"\n"+
			"+++\n", keyword, idx+1))
	}

	globalVar.questionVarMap["keyword"] = keyword

QuestionLoop:
	for _, questionInfo := range config.QuestionInfos {
		if len(questionInfo.Include) != 0 {
			included := false
			for _, includeRegex := range questionInfo.IncludeRe {
				if matched := includeRegex.MatchString(pathRelative); matched {
					included = true
				}
			}
			if !included {
				continue
			}
		}

		for _, excludeRegex := range questionInfo.ExcludeRe {
			if matched := excludeRegex.MatchString(pathRelative); matched {
				continue QuestionLoop
			}
		}

		var answerInfo AnswerInfo
		answerInfo.QuestionDesc = questionInfo.Desc
		//answerInfo.Question = fmt.Sprintf(questionInfo.Template, keyWord)

		var tpl bytes.Buffer
		err = questionInfo.TemplateObj.Execute(&tpl, globalVar.questionVarMap)
		checkErr(err)

		answerInfo.Question = tpl.String()

		thisPathCache := thisCacheFolder + pathMK + keyword + "_" + questionInfo.Key + ".md"
		//fmt.Println("thisPathCache=" + thisPathCache)

		if io.FileExists(thisPathCache) {
			answerInfo.Answer = io.FileToString(thisPathCache)
		} else {
			answerInfo.Answer = strings.Trim(askChatGpt(answerInfo.Question), " ")

			if answerInfo.Answer != "" {
				globalVar.errorNum = 0 //get answer reset errorNum
				io.StringToFile(thisPathCache, answerInfo.Answer)
			} else {
				fmt.Println("No answer=" + keyword)
				globalVar.errorNum++

				if globalVar.errorNum > 10 {
					err = errors.New("Too many empty answer:" + string(globalVar.errorNum))
					panic(err)
				}
			}
		}

		err = globalVar.answerOutputTemplate.Execute(f, answerInfo)
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
		client := openai.NewClient(parm.apiKeyOpenAI)
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

/*
func initQuestion() {
	if lang == "english" {
		var q1 QuestionInfo
		q1.Desc = "Preview"
		q1.Key = "preview"
		q1.Template = "give me five question about %s"
		questionInfoSlice = append(questionInfoSlice, q1)

		var q2 QuestionInfo
		q2.Desc = "Explain"
		q2.Key = "explain"
		q2.Template = "Explain %s and give an example"
		questionInfoSlice = append(questionInfoSlice, q2)

		var q3 QuestionInfo
		q3.Desc = "Keypoint"
		q3.Key = "keypoint"
		q3.Template = "List the key points of %s"
		questionInfoSlice = append(questionInfoSlice, q3)

		var q4 QuestionInfo
		q4.Desc = "Review"
		q4.Key = "test"
		q4.Template = "Give me 5 medium-difficulty questions with answers about %s"
		questionInfoSlice = append(questionInfoSlice, q4)

		//var q5 QuestionInfo
		//q5.desc = "Related webpage"
		//q5.cacheKey = "ref"
		//q5.template = "List the relevant introduction webpages about %s"
		//questionInfoSlice = append(questionInfoSlice, q5)
	} else if lang == "zh-cn" {
		var q1 QuestionInfo
		q1.Desc = "习题预习"
		q1.Key = "preview"
		q1.Template = "给我5题%s的问题"
		questionInfoSlice = append(questionInfoSlice, q1)

		var q2 QuestionInfo
		q2.Desc = "说明知识"
		q2.Key = "explain"
		q2.Template = "说明%s并举例"
		questionInfoSlice = append(questionInfoSlice, q2)

		var q3 QuestionInfo
		q3.Desc = "汇总重点"
		q3.Key = "keypoint"
		q3.Template = "条列%s的重点"
		questionInfoSlice = append(questionInfoSlice, q3)

		var q4 QuestionInfo
		q4.Desc = "知识测验"
		q4.Key = "test"
		q4.Template = "给我5题%s的中等难度问题，并在后面列出答案"
		questionInfoSlice = append(questionInfoSlice, q4)

		//var q5 QuestionInfo
		//q5.desc = "网络数据"
		//q5.cacheKey = "ref"
		//q5.template = "给我5篇%s的网络数据"
		//questionInfoSlice = append(questionInfoSlice, q5)
	} else if lang == "zh-tw" {
		var q1 QuestionInfo
		q1.Desc = "習題預習"
		q1.Key = "preview"
		q1.Template = "給我5題%s的中文問題"
		questionInfoSlice = append(questionInfoSlice, q1)

		var q2 QuestionInfo
		q2.Desc = "說明知識"
		q2.Key = "explain"
		q2.Template = "以中文說明%s並舉例"
		questionInfoSlice = append(questionInfoSlice, q2)

		var q3 QuestionInfo
		q3.Desc = "彙總重點"
		q3.Key = "keypoint"
		q3.Template = "以中文條列%s的重點"
		questionInfoSlice = append(questionInfoSlice, q3)

		var q4 QuestionInfo
		q4.Desc = "知識測驗"
		q4.Key = "test"
		q4.Template = "以中文給我5題%s的中等難度問題，並在後面列出答案"
		questionInfoSlice = append(questionInfoSlice, q4)

		//var q5 QuestionInfo
		//q5.desc = "網路資料"
		//q5.cacheKey = "ref"
		//q5.template = "給我5篇%s的中文網路資料"
		//questionInfoSlice = append(questionInfoSlice, q5)
	}
}
*/
