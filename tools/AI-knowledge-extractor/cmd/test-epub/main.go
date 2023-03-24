package main

import (
	"log"

	"golang.org/x/net/html"

	"github.com/bmaupin/go-epub"
)

const (
	effectiveGoCoverImg      = "assets/covers/effective-go.png"
	effectiveGoFilename      = "Effective Go.epub"
	effectiveGoSectionTag    = "h2"
	effectiveGoTitle         = "Effective Go"
	effectiveGoTitleFilename = "title.xhtml"
	effectiveGoUrl           = "https://golang.org/doc/effective_go.html"
	epubCSSFile              = "assets/styles/epub.css"
	preFontFile              = "assets/fonts/SourceCodePro-Regular.ttf"
)

type epubSection struct {
	title    string
	filename string
	nodes    []html.Node
}

func main() {
	var err error
	// Create a new EPUB
	e := epub.NewEpub("My title")

	// Set the author
	e.SetAuthor("Hingle McCringleberry")

	// Add a section
	section1Body := `<h1>Section 1</h1>
	<p>This is a paragraph.</p>`
	e.AddSection(section1Body, "Section 1", "", "")

	// Write the EPUB
	err = e.Write("My EPUB.epub")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
