package main

import (
	"github.com/atsushinee/go-markdown-generator/doc"
	"io/ioutil"
	"log"
)

func main() {
	example()
}

func example() {
	code, _ := ioutil.ReadFile("main.go")
	book := doc.NewMarkDown()
	book.WriteTitle("Go-MarkDownDoc-Generator", doc.LevelTitle).
		WriteLines(2)

	book.WriteMultiCode(string(code), "go")

	book.WriteTitle("Author", doc.LevelNormal).
		WriteCodeLine("lichun")

	book.WriteTitle("Website", doc.LevelNormal)
	book.WriteLinkLine("lichunorz", "https://lichunorz.com")

	t := doc.NewTable(4, 4)
	t.SetTitle(0, "Version")
	t.SetTitle(1, "Date")
	t.SetTitle(2, "Creator")
	t.SetTitle(3, "Remarks")
	t.SetContent(0, 0, "v1")
	t.SetContent(0, 1, "2019-11-21")
	t.SetContent(0, 2, "Lee")
	t.SetContent(0, 3, "无")
	book.WriteTable(t)
	err := book.Export("README.md")
	if err != nil {
		log.Fatal(err)
	}
}
