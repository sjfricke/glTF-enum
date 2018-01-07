package main

import (
	// "os"
	"github.com/json-iterator/go"
	"reflect"
	// "io/ioutil"
	"log"
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

// enums is the output json going to enums.js
type enums struct {
	Name  string `json:"name"`
	Types []struct {
		Link     string `json:"link"`
		Name     string `json:"name"`
		Required string `json:"required"`
		Type     string `json:"type"`
	} `json:"types"`
	Value int64 `json:"value"`
}

const SchemasUrl string = "https://api.github.com/repositories/7921466/contents/specification/2.0/schema"

const SchemaTest string = "https://raw.githubusercontent.com/KhronosGroup/glTF/master/specification/2.0/schema/accessor.schema.json"
const ReadmeUrl string = "https://raw.githubusercontent.com/KhronosGroup/glTF/master/specification/2.0/README.md"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
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

	log.Println(reflect.TypeOf(s.Properties))

	for k, v := range s.Properties {
		log.Println(k)
		log.Println(v.Description)
	}
}
