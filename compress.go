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
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// Create deploy folder first
	path := "deploy/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0666)
	}

	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/javascript", js.Minify)
	m.AddFunc("image/svg+xml", svg.Minify)

	// non minified files
	data, err := ioutil.ReadFile("site/searchicon.png")
	checkErr(err)
	err = ioutil.WriteFile("deploy/searchicon.png", data, 0644)
	checkErr(err)
	data, err = ioutil.ReadFile("site/enums.js")
	checkErr(err)
	err = ioutil.WriteFile("deploy/enums.js", data, 0644)
	checkErr(err)

	// style.css
	b, err := ioutil.ReadFile("site/style.css")
	checkErr(err)
	mb, err := m.Bytes("text/css", b)
	checkErr(err)
	f, err := os.Create("deploy/style.css")
	checkErr(err)
	defer f.Close()
	_, err = f.WriteString(string(mb))
	checkErr(err)

	// main.js
	b, err = ioutil.ReadFile("site/main.js")
	checkErr(err)
	mb, err = m.Bytes("text/javascript", b)
	checkErr(err)
	f, err = os.Create("deploy/main.js")
	checkErr(err)
	defer f.Close()
	_, err = f.WriteString(string(mb))
	checkErr(err)

	// index.html
	b, err = ioutil.ReadFile("site/index.html")
	checkErr(err)
	mb, err = m.Bytes("text/html", b)
	checkErr(err)
	f, err = os.Create("deploy/index.html")
	checkErr(err)
	defer f.Close()
	_, err = f.WriteString(string(mb))
	checkErr(err)

	// gltf_logo.svg
	b, err = ioutil.ReadFile("site/gltf_logo.svg")
	checkErr(err)
	mb, err = m.Bytes("image/svg+xml", b)
	checkErr(err)
	f, err = os.Create("deploy/gltf_logo.svg")
	checkErr(err)
	defer f.Close()
	_, err = f.WriteString(string(mb))
	checkErr(err)
}
