# KIʻI

A [golang][go] command line utility and library to extract distinct, representative icons from various structured data formats and other forms embedded in [HTML][html] -- typically associated with a particular [website][website] or [web page][webpage].

## Supported Formats

In order of preference (default sort order):

- [JSON-LD][json-ld]
- [Open Graph][og]
- [HTML External Resource Link][link]
- [Apple Touch Icon][apple]
- [Twitter Cards Markup][twitter]
- [Microsoft Tile Image][ms]
- [Favicon][favicon]

## Installation

```sh
go get github.com/daetal-us/kii/cmd/kii
```

## Usage
_Note: results are always returned as an [array][array]._

### Command Line

Retrieve icons for a url:
```sh
kii http://apple.com
```

### Golang

```go
package main

import (
	"log"
	"encoding/json"

	"github.com/daetal-us/kii"
)

func main() {
	extractFromHTML()
	extractFromURL()
}

func extractFromHTML() {
	html := `
	<html>
		<head>
			<meta property="og:image" content="a.png">
			<meta name="twitter:image" content="a.png" />
			<link rel="shortcut icon" href="b.ico" />
		</head>
		<body>
			<script type="application/ld+json">
			{
				"@type": "Organization",
				"name": "Example",
				"image": "c.png"
			}
			</script>
		</body>
	</html>
	`
	results, err := kii.FromHTML(html)
	if err != nil {
		log.Fatal(err)
	}
	encoded, err := json.Marshal(results)
	if err != nil {
		log.fatal(err)
	}
	log.Println(string(encoded)) // ["c.png","a.png","b.ico"]
}

func extractFromURL() {
	results, err := kii.FromURL("http://apple.com")
	if err != nil {
		log.Fatal(err)
	}
	encoded, err := json.Marshal(results)
	if err != nil {
		log.fatal(err)
	}
	log.Println(string(encoded)) // [...,"favicon.ico"]
}
```

[go]:https://golang.org
[favicon]:https://en.wikipedia.org/wiki/Favicon
[website]:https://en.wikipedia.org/wiki/Website
[webpage]:https://en.wikipedia.org/wiki/Web_page
[json-ld]:https://json-ld.org
[og]:http://ogp.me
[apple]:https://developer.apple.com/library/archive/documentation/AppleApplications/Reference/SafariWebContent/ConfiguringWebApplications/ConfiguringWebApplications.html
[twitter]:https://developer.twitter.com/en/docs/tweets/optimize-with-cards/overview/markup
[ms]:https://docs.microsoft.com/en-us/previous-versions/windows/internet-explorer/ie-developer/platform-apis/dn255024(v=vs.85)#msapplication-TileImage
[link]:https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link
[html]:https://www.w3.org/html
[array]:https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array
