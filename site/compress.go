package main

import (
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/svg"
	// "log"
	"io/ioutil"
	"os"
	"flag"
	"fmt"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var path = flag.String("path", "deploy", "path compress files too")

func main() {

	// Create deploy folder first
	flag.Parse()
	if _, err := os.Stat(*path); os.IsNotExist(err) {
		os.MkdirAll(*path, 0755)
	}

	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/javascript", js.Minify)
	m.AddFunc("image/svg+xml", svg.Minify)

	// non minified files
	data, err := ioutil.ReadFile("searchicon.png")
	checkErr(err)
	err = ioutil.WriteFile(fmt.Sprintf("%s/searchicon.png", *path), data, 0644)
	checkErr(err)
	data, err = ioutil.ReadFile("enums.js")
	checkErr(err)
	err = ioutil.WriteFile(fmt.Sprintf("%s/enums.js", *path), data, 0644)
	checkErr(err)

	// style.css
	b, err := ioutil.ReadFile("style.css")
	checkErr(err)
	mb, err := m.Bytes("text/css", b)
	checkErr(err)
	f, err := os.Create(fmt.Sprintf("%s/style.css", *path))
	checkErr(err)
	defer f.Close()
	_, err = f.WriteString(string(mb))
	checkErr(err)

	// main.js
	b, err = ioutil.ReadFile("main.js")
	checkErr(err)
	mb, err = m.Bytes("text/javascript", b)
	checkErr(err)
	f, err = os.Create(fmt.Sprintf("%s/main.js", *path))
	checkErr(err)
	defer f.Close()
	_, err = f.WriteString(string(mb))
	checkErr(err)

	// index.html
	b, err = ioutil.ReadFile("index.html")
	checkErr(err)
	mb, err = m.Bytes("text/html", b)
	checkErr(err)
	f, err = os.Create(fmt.Sprintf("%s/index.html", *path))
	checkErr(err)
	defer f.Close()
	_, err = f.WriteString(string(mb))
	checkErr(err)

	// gltf_logo.svg
	b, err = ioutil.ReadFile("gltf_logo.svg")
	checkErr(err)
	mb, err = m.Bytes("image/svg+xml", b)
	checkErr(err)
	f, err = os.Create(fmt.Sprintf("%s/gltf_logo.svg", *path))
	checkErr(err)
	defer f.Close()
	_, err = f.WriteString(string(mb))
	checkErr(err)
}
