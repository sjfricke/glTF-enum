# glTF-enum

This site was developed for ~three~ four reason:

1. Wanted to force myself to learn the glTF spec better
2. Wanted a productive project to learn Golang better
3. **Hate** `ctrl+f`ing through spec sheets
4. Saw this domain was not taken...

> Side note: Golang probably was not the best tool for JSON scraping IMO after working on this

## Suggestions or improvements

Open to expanding project to add more useful items, leave PR or issue.

## How it works

The HTML is designed to be populated via the [enums.js](site/enums.js) file. The [glTF-enum.go](glTF-enum.go) Golang script kicks off every day and cross references the [glTF Spec](https://github.com/KhronosGroup/glTF/blob/master/specification/2.0/) where it generates a JSON object and writes it out to [enums.js](site/enums.js).

## Run it yourself

- Get the repo: `go get github.com/sjfricke/glTF-enum`
- `cd` to the `$GOPATH/src/github.com/sjfricke/glTF-enum`
- Grab json library: `go get github.com/json-iterator/go`
- Run `go run glTF-enum.go`
- Open `site/index.html` in a browser


## Deploying

I also have a [compress.go](site/compress.go) script that runs and it just minifies all the HTML files for deployment.

Run the same commands but after run `go run compress.go` to create a `deploy` folder with a compressed version of the website. You will also need to to run `go get github.com/tdewolff/minify`.
