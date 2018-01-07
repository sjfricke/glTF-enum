package main

import (
	"github.com/json-iterator/go"
	// "io/ioutil"
	"bytes"
	"log"
	"os"
	"reflect"
	"time"
	// "net/http"
)

type schemaList []struct {
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
		Description string  `json:"description"`
		Enum        []int64 `json:"enum"`
		Type        string  `json:"type"`
	} `json:"anyOf"`
	Items struct {
		Type string `json:"type"`
	} `json:"items"`
	Default                 int64  `json:"default"`
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
	Name  string  `json:"name"`
	Types []eType `json:"types"`
	Value int64   `json:"value"`
}

const SchemasUrl string = "https://api.github.com/repositories/7921466/contents/specification/2.0/schema"

const SchemaTest string = "https://raw.githubusercontent.com/KhronosGroup/glTF/master/specification/2.0/schema/accessor.schema.json"
const ReadmeUrl string = "https://raw.githubusercontent.com/KhronosGroup/glTF/master/specification/2.0/README.md"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {

	var enumJson []enums

	//log.Println(os.Args)

	// res, err := http.Get(SchemasUrl)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// schemas, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var sUrls schemaList
	err := json.Unmarshal(schemas, &sUrls)
	if err != nil {
		log.Fatal(err)
	}

	for _, sUrl := range sUrls[0:1] {
		log.Println(sUrl.DownloadURL)
	}

	var s schema
	err = json.Unmarshal(indices, &s)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(reflect.TypeOf(s))

	// each property key is the name of it
	// each AnyOf is array of enums
	for k, v := range s.Properties {
		log.Println(k) // compentType

		// need to scan and get type first
		// TODO see if can assume last ele in array is type
		var valueType string
		for _, vv := range v.AnyOf {
			if vv.Type != "" {
				valueType = vv.Type
				break
			}
		}

		log.Println(valueType)

		for _, vv := range v.AnyOf {
			if len(vv.Enum) > 0 {
				log.Println(vv.Description)
				log.Println(vv.Enum[0])
				var e enums
				e.Name = vv.Description
				e.Value = vv.Enum[0] // TODO, assume possible filled array
				e.Types = append(e.Types, eType{Link: "a", Name: k, Required: "c", Type: valueType})
				enumJson = append(enumJson, e)
			}
		}
	}

	// Time to write everything out to file now

	outJson, _ := json.Marshal(enumJson)

	f, err := os.Create("test.json")
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
