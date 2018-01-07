package main

import (
	// "os"
	"github.com/json-iterator/go"
	// "io/ioutil"
	"log"
	// "net/http"
)

type schemaList []struct {
	DownloadURL string `json:"download_url"`
}

type schema struct {
	Description string                 `json:"description"`
	Required    []string               `json:"required"`
	Title       string                 `json:"title"`
	Type        string                 `json:"type"`
	Properties  map[string]interface{} `json:"properties"`
}

type properties struct {
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
}
