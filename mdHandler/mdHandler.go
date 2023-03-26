package mdHandler

import (
	"fmt"
	"html/template"
	"os"

	"srv/util"

	bf "github.com/russross/blackfriday/v2"
)

type dataHTML struct {
	PageTitle string
	PageBody  template.HTML
}

func HandleFile(fileName string, templateFileName string) (dataHTML, *template.Template) {
	fmt.Println("Inside mdHandler")
	mdFileData, err := os.ReadFile(fileName)
	util.CheckErr(err)
	mdHTML := bf.Run(mdFileData)
	pageData := dataHTML{
		PageTitle: "test",
		PageBody:  template.HTML(mdHTML),
	}
	tmpl := template.Must(template.ParseFiles(templateFileName))
	fn := fileName[:len(fileName)-3]
	err = os.MkdirAll("site/"+fn, os.ModePerm)
	util.CheckErr(err)
	htmlFileIO, err := os.Create("site/" + fn + "/index.html")
	util.CheckErr(err)
	tmpl.Execute(htmlFileIO, pageData)

	return pageData, tmpl
}

func HandleDir(dirName string, templateFileName string) {
	files, err := os.ReadDir(dirName)
	util.CheckErr(err)
	for _, file := range files {
		fileName := file.Name()
		fileNameDir := dirName + "/" + file.Name()
		mdFileData, err := os.ReadFile(fileNameDir)
		util.CheckErr(err)
		mdHTML := bf.Run(mdFileData)
		pageData := dataHTML{
			PageTitle: "test",
			PageBody:  template.HTML(mdHTML),
		}
		tmpl := template.Must(template.ParseFiles(templateFileName))
		fn := fileName[:len(fileName)-3]
		err = os.MkdirAll("site/"+ fn, os.ModePerm)
		util.CheckErr(err)
		htmlFileIO, err := os.Create("site/" + fn + "/index.html")
		util.CheckErr(err)
		tmpl.Execute(htmlFileIO, pageData)
	}
}
