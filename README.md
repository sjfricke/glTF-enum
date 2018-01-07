# glTF-enum

This site was developed for ~three~ four reason:

1. Wanted to force myself to learn the glTF spec better
2. **Hate** `ctrl+f`ing through spec sheets
3. Wanted a productive project to practice my Golang with
4. Saw this domain was not taken...

> Side note: This should NOT have been written in Golang as the JSON scarping support is beyond frusterating and could have finished this with NodeJS in fraction of the time

## How it works

The HTML is designed to be populated via the [enums.js](site/enums.js) file. The [glTF-enum.go](glTF-enum.go) Golang script kicks off every day and cross references the [glTF Spec](https://github.com/KhronosGroup/glTF/blob/master/specification/2.0/) where it generates a JSON object and writes it out to [enums.js](site/enums.js).

## Deploying

I also have a [compress.go](compress.go) script that runs and it just minifies all the HTML files for deployment.