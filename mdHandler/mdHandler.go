package mdHandler

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"srv/util"

	"github.com/BurntSushi/toml"
	bf "github.com/russross/blackfriday/v2"
)

type dataHTML struct {
	PageTitle       string
	PageDescription string
	PageBody        template.HTML
}

type postConfig struct {
	Title       string
	Description string
}

var postCfg_ []postConfig

func ParseCfg(fileName string) postConfig {
	tomlData, err := os.ReadFile(fileName)
	util.CheckErr(err)

	var cfg postConfig
	err = toml.Unmarshal(tomlData, &cfg)
	util.CheckErr(err)

	return cfg
}

func HandleMdFile(fileName string) []byte {
	fmt.Println("Inside mdHandler")
	mdFileData, err := os.ReadFile(fileName)
	util.CheckErr(err)

	mdHTML := bf.Run(mdFileData)

	return mdHTML
}

func GetFileNameWithPath(fileName string, dirName string) string {
	return dirName + "/" + fileName
}

func ParsePost(postName string, templateFileName string) {
	/*
		read all files in the directory
		process html into markdown
		process config details (post name, date, author, description)
		then convert .md file into html with config info
	*/
	files, err := os.ReadDir(postName)
	util.CheckErr(err) // program will exit here if error

	mdFile := "NULL"
	cfgFile := "NULL"
	//mdFileRaw := "NULL"

	/* iterate through all the files in the directory */
	for _, file := range files {
		fileName := file.Name()
		fileType := filepath.Ext(fileName)

		fmt.Println("file,", fileName, ",fileType,", fileType)

		if fileType == ".md" {
			// exec md file stuff
			mdFile = GetFileNameWithPath(fileName, postName)
			//mdFileRaw = fileName
		} else if fileType == ".toml" {
			// exec cfg stuff here
			cfgFile = GetFileNameWithPath(fileName, postName)
		}
	}

	/* get details from cfg and add it to html struct then pass struct
	to handle md function and execute template struct in the end */
	log.Println("mdFile,", mdFile)
	log.Println("cfgFile,", cfgFile)
	var pageData dataHTML

	/* parse cfg */
	cfgData := ParseCfg(cfgFile)
	pageData.PageTitle = cfgData.Title
	pageData.PageDescription = cfgData.Description

	postCfg_ = append(postCfg_, cfgData)

	log.Println("desc,", pageData.PageDescription)
	log.Println("title,", pageData.PageTitle)

	/* parse md */
	htmlData := HandleMdFile(mdFile)
	pageData.PageBody = template.HTML(htmlData)

	tmpl := template.Must(template.ParseFiles(templateFileName))

	log.Println("creating directory ", pageData.PageTitle)
	fn := pageData.PageTitle
	err = os.MkdirAll("site/"+fn, os.ModePerm)
	util.CheckErr(err)

	htmlFileIO, err := os.Create("site/" + fn + "/index.html")
	util.CheckErr(err)

	tmpl.Execute(htmlFileIO, pageData)
}

func HandleDir(dirName string, markdownTemplate string, blogTemplate string) {
	files, err := os.ReadDir(dirName)
	util.CheckErr(err)
	for _, file := range files {
		fileName := file.Name()
		fileNameDir := dirName + "/" + fileName

		ParsePost(fileNameDir, markdownTemplate)
	}

	/* create blog index file from html template */
	tmpl := template.Must(template.ParseFiles(blogTemplate))
	blogFileIO, err := os.Create("blog.html")
	util.CheckErr(err)
	tmpl.Execute(blogFileIO, postCfg_)
}
