package main

import (
	"bytes"
	// "github.com/json-iterator/go"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	// "reflect"
	"net/http"
	"strings"
	"time"
)

type schemaList []struct {
	Name        string `json:"name"`
	DownloadURL string `json:"download_url"`
}

type schema struct {
	Description string              `json:"description"`
	Required    []string            `json:"required"`
	Title       string              `json:"title"`
	Type        string              `json:"type"`
	Properties  map[string]property `json:"properties"`
}

type property struct {
	AnyOf []struct {
		Description string            `json:"description"`
		Enum        []json.RawMessage `json:"enum"`
		Type        string            `json:"type"`
	} `json:"anyOf"`
	Items struct {
		Type string `json:"type"`
	} `json:"items"`
	// Default                 int64  `json:"default"` // not consistant
	Description             string `json:"description"`
	Format                  string `json:"format"`
	GltfDetailedDescription string `json:"gltf_detailedDescription"`
	GltfURIType             string `json:"gltf_uriType"`
	GltfWebgl               string `json:"gltf_webgl"`
	Maximum                 int64  `json:"maximum"`
	MaxItems                int64  `json:"maxItems"`
	Minimum                 int64  `json:"minimum"`
	MinItems                int64  `json:"minItems"`
	MultipleOf              int64  `json:"multipleOf"`
	Type                    string `json:"type"`
}

type eType struct {
	Link     string `json:"link"`
	Name     string `json:"name"`
	Required string `json:"required"`
	Type     string `json:"type"`
}

// enums is the output json going to enums.js
type enums struct {
	Name  string          `json:"name"`
	Types []eType         `json:"types"`
	Value json.RawMessage `json:"value"`
}

const SchemasUrl string = "https://api.github.com/repositories/7921466/contents/specification/2.0/schema"
const SchemaTest string = "https://raw.githubusercontent.com/KhronosGroup/glTF/master/specification/2.0/schema/accessor.schema.json"
const ReadmeUrl string = "https://raw.githubusercontent.com/KhronosGroup/glTF/master/specification/2.0/README.md"

var tags = make([]string, 0, 256) // TODO by [][]byte

// Readme is used to setup the tags from the readme for links
func Readme() {
	// res, err := http.Get(ReadmeUrl)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// readme, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	readme, err := ioutil.ReadFile("READMEE.md")
	if err != nil {
		panic(err)
	}

	start := bytes.Index(readme, []byte("# Properties Reference"))
	end := bytes.Index(readme, []byte("# Acknowledgments"))
	props := readme[start:end]

	n, x, y := 0, 0, 0
	for {
		x = bytes.Index(props[n:], []byte("####"))
		if x < 0 {
			break
		}

		y = bytes.Index(props[n+x:], []byte("\n"))
		s := string(props[n+x+5 : n+x+y]) // 5 added from ####_
		s = strings.ToLower(s)
		s = strings.TrimSpace(s)
		s = strings.Replace(s, ":white_check_mark:", "white_check_mark", -1)
		s = strings.Replace(s, ".", "", -1)
		s = strings.Replace(s, " ", "-", -1)
		tags = append(tags, s)

		n = n + x + 1
	}
}

// Requires is to check for a string in a string slice
// returns if string to be seen on site
func Requires(a []string, b string) string {
	for _, c := range a {
		if c == b {
			return "Required"
		}
	}
	return "Not Required"
}

func Link(file, title, prop string) string {

	p := []byte(prop)
	p = bytes.ToLower(p)

	// gets type from file
	f := []byte(file)
	f = f[0:bytes.Index(f, []byte(".schema.json"))]
	f = f[bytes.LastIndex(f, []byte("."))+1:]
	f = bytes.ToLower(f)
	f = append(f, p...)

	// cases like "Animation Sampler" force a check of title first
	t := []byte(title)
	t = bytes.ToLower(t)
	t = bytes.Replace(t, []byte(" "), []byte("-"), -1)
	t = append(t, p...)

	for _, s := range tags {
		if bytes.Index([]byte(s), t) == 0 {
			return s
		}

		if bytes.Index([]byte(s), f) == 0 {
			return s
		}
	}

	log.Println(file, ",", title, ",", prop)
	panic("Cannot find a valid link")
}

func main() {
	// var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var enumJson []enums

	// res, err := http.Get(SchemasUrl)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// schemas, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// set tags ups
	Readme()
	log.Println("Tags: ", len(tags))

	var sUrls schemaList
	err := json.Unmarshal(schemas, &sUrls)
	if err != nil {
		log.Fatal(err)
	}

	for _, sUrl := range sUrls[0:2] {

		res, err := http.Get(sUrl.DownloadURL)
		if err != nil {
			log.Fatal(err)
		}

		schemaData, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		var s schema
		err = json.Unmarshal(schemaData, &s)
		if err != nil {
			log.Fatal(err)
		}

		// each property key is the name of it
		// each AnyOf is array of enums
		for prop, v := range s.Properties {

			// need to scan and get type first
			// TODO see if can assume last ele in array is type
			var valueType string
			for _, vv := range v.AnyOf {
				if vv.Type != "" {
					valueType = vv.Type
					break
				}
			}

			for _, vv := range v.AnyOf {
				if len(vv.Enum) > 0 {
					var e enums

					// TODO better way of handling string enums
					if vv.Description == "" {
						raw := json.RawMessage(vv.Enum[0])
						j, err := json.Marshal(&raw)
						if err != nil {
							panic(err)
						}
						e.Name = string(j)
					} else {
						e.Name = vv.Description
						e.Value = vv.Enum[0] // TODO, assume possible filled array
					}

					req := Requires(s.Required, prop)

					// creates name by adding title and prop name with periods
					n := append([]byte(s.Title), []byte(" ")...)
					n = append(n, []byte(prop)...)
					n = bytes.Replace(n, []byte(" "), []byte("."), -1)

					e.Types = append(e.Types, eType{
						Link:     Link(sUrl.Name, s.Title, prop),
						Name:     string(n),
						Required: req,
						Type:     valueType})

					enumJson = append(enumJson, e)
				}
			}
		}
	}

	// Time to write everything out to file now

	outJson, _ := json.Marshal(enumJson)

	f, err := os.Create("site/enums.js")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	t := time.Now()
	var dateStr bytes.Buffer
	dateStr.WriteString(`const LAST_UPDATE = "`)
	dateStr.WriteString(t.UTC().Format(time.UnixDate))
	dateStr.WriteString(`"; const ENUMS = `)
	if _, err = f.WriteString(dateStr.String()); err != nil {
		panic(err)
	}

	if _, err = f.WriteString(string(outJson[:])); err != nil {
		panic(err)
	}

	// TODO find a better way to add semicolon
	if _, err = f.WriteString(";"); err != nil {
		panic(err)
	}
}
