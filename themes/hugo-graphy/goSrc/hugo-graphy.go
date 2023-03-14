package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/learninfun/hugo-graphy/parser/metadecoders"
	"github.com/learninfun/hugo-graphy/tools/tools_check"
	"github.com/learninfun/hugo-graphy/tools/tools_convert"
	"github.com/learninfun/hugo-graphy/tools/tools_other"
)

var (
	siteFolder       string
	sourseFolder     string
	langSourceFolder string
	langNow          string //currently handle lang
	resultFolder     string
	sourseFolderLen  int
	pathMK           string
	InLinkRegex      *regexp.Regexp
	cardMap          map[string]Card
	dirCardMap       map[string]CardDir
)

type Card struct {
	Name   string   `json:"name"`
	Url    string   `json:"url"`
	Md     string   `json:"md"`
	Link   []string `json:"to"`
	Linked []string `json:"fm"`
}

type CardDir struct {
	Name    string   `json:"name"`
	Url     string   `json:"url"`
	Sub     []string `json:"sub"`     //Children node
	SubType []string `json:"subType"` //children type: D, P
}

func main() {
	fmt.Println("Hello World")
	//fmt.Println(os.Args)
	initGlobalVar()
	confMap := getConfMap()

	err := os.RemoveAll(resultFolder) //clear ori result
	check(err)

	graphyJsFolder := siteFolder + pathMK + "static" + pathMK + "hugo-graphy" + pathMK + "js"
	if !tools_check.FileExists(graphyJsFolder) {
		err := os.MkdirAll(graphyJsFolder, 0755)
		check(err)
	}

	makeFolderIfNotExist(resultFolder)

	languages, languagesFound := confMap["Languages"]
	if !languagesFound { //without language, just walk through souseFolder
		cardMap = make(map[string]Card)
		dirCardMap = make(map[string]CardDir)
		sourceFileInfo, err := os.Lstat(sourseFolder)
		check(err)
		rootDirCard := convertDirToCard(sourseFolder, sourceFileInfo)
		check(err)
		myFileWalk(&rootDirCard, sourseFolder, sourceFileInfo)

		cardMapStr := tools_convert.JsonToString(confMap)
		tools_convert.StringToFile(graphyJsFolder+pathMK+"hugo-graphy-data.js", "hugoGraphyData="+cardMapStr)
	} else {
		languagesMap := languages.(map[string]any)
		for langName, langConf := range languagesMap { //loop all lang
			langNow = langName
			langConfMap := langConf.(map[string]any)
			fmt.Println("langName:", langName, "=>", "langConf:", langConfMap)
			langContentDir, langContentDirFound := langConfMap["contentDir"]
			if !langContentDirFound {
				fmt.Println(langContentDir)
				panic("Not support language without contentDir")
			}

			langContentDirStr := langContentDir.(string)
			langSourceFolder = siteFolder + pathMK + "contentOri" + pathMK + langContentDirStr[8:]

			cardMap = make(map[string]Card)
			dirCardMap = make(map[string]CardDir)
			sourceFileInfo, err := os.Lstat(langSourceFolder)
			check(err)
			rootDirCard := convertDirToCard(langSourceFolder, sourceFileInfo)
			fileInfo, err := os.Lstat(langSourceFolder)
			check(err)
			myFileWalk(&rootDirCard, langSourceFolder, fileInfo)

			/*
				cardMapStrBeauty := tools_convert.JsonToStringBeauty(cardMap)
				fmt.Println(cardMapStrBeauty)

				cardMapStrBeauty = tools_convert.JsonToStringBeauty(dirCardMap)
				fmt.Println(cardMapStrBeauty)

				cardMapStr := tools_convert.JsonToString(cardMap)
				dirCardMapStr := tools_convert.JsonToString(dirCardMap)
				tools_convert.StringToFile(graphyJsFolder+pathMK+"hugo-graphy-data_"+langName+".js", "var hugoGraphyData="+cardMapStr+";var hugoGraphyDirData="+dirCardMapStr)
			*/
		}
	}
}

func initGlobalVar() {
	var err error

	currentFolder, err := os.Getwd()
	check(err)
	fmt.Println(len(os.Args))
	if len(os.Args) >= 2 {
		siteFolder = os.Args[1]
		fmt.Println(siteFolder)
	} else if tools_check.FileExists(currentFolder + pathMK + "contentOri") {
		siteFolder = currentFolder
	} else {
		siteFolder = filepath.Dir(currentFolder) //如果執行目錄是在hugo-graphy那層，要先回到前一層
		siteFolder = filepath.Dir(siteFolder)
		siteFolder = filepath.Dir(siteFolder)
	}

	pathMK = string(os.PathSeparator)

	sourseFolder = siteFolder + pathMK + "contentOri"
	sourseFolderLen = len(sourseFolder)
	resultFolder = siteFolder + pathMK + "content"

	regexInLink := `\[\[[^\/\\\:\*\?\"\<\>\|\]]{1,254}\]\]`
	InLinkRegex, err = regexp.Compile(regexInLink)
	check(err)
}

func myFileWalk(parentCard *CardDir, pathSourse string, fileInfo os.FileInfo) {
	var pathRelative = pathSourse[sourseFolderLen:]
	var pathResult = resultFolder + pathRelative

	if fileInfo.IsDir() {
		makeFolderIfNotExist(pathResult)
		convertDirToCard(pathSourse, fileInfo)
		if langSourceFolder != pathSourse &&
			!tools_check.FileExists(pathSourse+pathMK+"_index.md") &&
			!tools_check.FileExists(pathSourse+pathMK+"_index.html") &&
			!tools_check.FileExists(pathSourse+pathMK+fileInfo.Name()+".md") {
			fmt.Println("Create _index.md:" + pathSourse + pathMK + "_index.md")
			create_indexMd(fileInfo.Name(), pathResult+pathMK+"_index.md")
		}

		files, err := ioutil.ReadDir(pathSourse)
		check(err)

		fmt.Println("Folder:" + fileInfo.Name())
		dirCard := convertDirToCard(pathSourse, fileInfo)
		for _, subFile := range files {
			myFileWalk(&dirCard, filepath.Join(pathSourse, subFile.Name()), subFile)
		}

		//dirArr = append(dirArr, dirCard)
		dirCardMap[dirCard.Url] = dirCard

		parentCard.Sub = append(parentCard.Sub, dirCard.Url)
		parentCard.SubType = append(parentCard.SubType, "D")
	} else {
		fmt.Println("File:" + fileInfo.Name())
		folderName := getFolderName(pathSourse)
		fmt.Println("FolderName:" + folderName)

		isFolderMd := fileInfo.Name() == folderName+".md"
		if isFolderMd {
			pathResult = pathResult[:len(pathResult)-len(fileInfo.Name())] + "_index.md" //folder\oo.md ==> folder\_index.md
		}
		handledMarkdown := handleMarkdown(fileInfo.Name(), pathSourse, pathResult, isFolderMd)
		if !handledMarkdown {
			tools_other.FileCopy(pathSourse, pathResult)
		} else {
			parentCard.Sub = append(parentCard.Sub, tools_convert.FileNameNoExt(fileInfo.Name()))
			parentCard.SubType = append(parentCard.SubType, "C")
		}
	}
}

/**
 * iterate all file to create new markdown with internal link
 */
/*
func walkAllFile(pathSourse string, fileInfo os.FileInfo, err error) error {
	check(err)

	var pathRelative = pathSourse[sourseFolderLen:]
	var pathResult = resultFolder + pathRelative
	if fileInfo.IsDir() {
		//fmt.Printf("Folder Name: %s\n", pathRelative)
		makeFolderIfNotExist(pathResult)
		convertDirToCard(pathSourse, fileInfo)
		if langSourceFolder != pathSourse &&
			!tools_check.FileExists(pathSourse+pathMK+"_index.md") &&
			!tools_check.FileExists(pathSourse+pathMK+"_index.html") {
			fmt.Println("Create _index.md:" + pathSourse + pathMK + "_index.md")
			create_indexMd(fileInfo.Name(), pathResult+pathMK+"_index.md")
		}
	} else {
		//fmt.Printf("file Name: %s\n", pathRelative)
		handledMarkdown := handleMarkdown(fileInfo.Name(), pathSourse, pathResult)
		if !handledMarkdown {
			tools_other.FileCopy(pathSourse, pathResult)
		}
	}

	return nil
}
*/

/**
 * Handle single makdownFile to create new markdown with internal link
 */
func handleMarkdown(fileName, pathSourse, pathResult string, isFolderMd bool) bool {
	if strings.ToLower(filepath.Ext(pathSourse)) != ".md" { //only handle markdown
		return false
	}

	if fileName == "_index.md" {
		return false
	}

	oriContent := tools_convert.FileTostring(pathSourse)
	newContent := oriContent

	//get card for this page
	fileNameNoExt := tools_convert.FileNameNoExt(fileName)
	fmt.Println(fileNameNoExt)
	cardThisPage, fmCardFound := cardMap[fileNameNoExt]
	if !fmCardFound {
		cardThisPage = Card{Name: fileNameNoExt}
	}
	cardThisPage.Url = getPageUrl(pathSourse)
	cardThisPage.Md = tools_convert.FileTostring(pathSourse)

	//get all internal link
	regexMachesArr := InLinkRegex.FindAllStringSubmatchIndex(oriContent, -1)
	offsetIdx := 0
	for i := 0; i < len(regexMachesArr); i++ {
		linkStr := oriContent[regexMachesArr[i][0]+2 : regexMachesArr[i][1]-2]
		anchorStr := fmt.Sprintf("<a id='inlink%d' class='inlink'>%s</a>", i, linkStr)
		newContent = newContent[:regexMachesArr[i][0]+offsetIdx] + anchorStr + newContent[regexMachesArr[i][1]+offsetIdx:]
		offsetIdx += len(anchorStr) - (regexMachesArr[i][1] - regexMachesArr[i][0])

		//get card link from this page
		cardLink, cardLinkFound := cardMap[linkStr]
		if !cardLinkFound {
			cardLink = Card{Name: linkStr}
		}

		//create link in both direction
		cardLink.Linked = append(cardLink.Linked, fileNameNoExt)
		cardThisPage.Link = append(cardThisPage.Link, linkStr)

		cardMap[linkStr] = cardLink //save card to map
	}

	cardMap[fileNameNoExt] = cardThisPage //save card to map
	fmt.Println("isFolderMd:" + fileNameNoExt + ": " + strconv.FormatBool(isFolderMd))
	//add page info
	if isFolderMd {
		newContent = fmt.Sprintf("+++\ntitle = \"%s\"\n+++", fileNameNoExt) + "\n" + newContent
	}

	tools_convert.StringToFile(pathResult, newContent)

	return true
}

func convertDirToCard(pathSourse string, fileInfo os.FileInfo) CardDir {
	cardThisDir := CardDir{Name: fileInfo.Name()}
	cardThisDir.Url = getPageUrl(pathSourse)

	return cardThisDir
}

/**
 * Get a url of page
 */
func getPageUrl(pathSourse string) string {
	result := langNow + pathSourse[len(langSourceFolder):]

	if result[:1] != pathMK {
		result = pathMK + result
	}

	result = result[:len(result)-len(filepath.Ext(result))]

	return filepath.ToSlash(result)
}

func create_indexMd(folderName, pathResult string) {
	content := fmt.Sprintf("+++\ntitle = \"%s\"\n+++", folderName)
	tools_convert.StringToFile(pathResult, content)
}

func getFolderName(pathSourse string) string {
	parentDir := filepath.Dir(pathSourse)
	parentDirName := filepath.Base(parentDir)
	return parentDirName
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func makeFolderIfNotExist(folderPath string) {
	if !tools_check.FileExists(folderPath) {
		err := os.Mkdir(folderPath, 0755)
		check(err)
	}
}

func getConfMap() map[string]any {
	doc := tools_convert.FileTostring(siteFolder + pathMK + "config.toml")
	docBype := []byte(doc)
	opts := metadecoders.Default
	mapConf, err := opts.UnmarshalToMap(docBype, metadecoders.TOML)
	check(err)

	return mapConf
}
