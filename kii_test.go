package kii

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestExtractFromJSONLD(t *testing.T) {
	html := `
		<html>
			<body>
				<script type="application/ld+json">{"image":"a"}</script>
				<script type="application/ld+json">{"logo":"b"}</script>
			</body>
		</html>
	`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatal(err)
	}

	result := findJSONLD(doc)

	if len(result) != 2 {
		t.Fatal("Expected 2 results.")
	}

	if result[0] != "a" {
		t.Errorf("Unexpected result: %v", result[0])
	}

	if result[1] != "b" {
		t.Errorf("Unexpected result: %v", result[1])
	}
}

func TestExtractFromOpenGraph(t *testing.T) {
	html := `
		<html>
			<head>
				<meta property="og:image" content="a" />
			</head>
		</html>
	`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatal(err)
	}

	result := findOpenGraph(doc)

	if len(result) != 1 {
		t.Fatal("Expected 1 result.")
	}

	if result[0] != "a" {
		t.Errorf("Unexpected result: %v", result[0])
	}
}

func TestExtractFromTwitterCard(t *testing.T) {
	html := `
		<html>
			<head>
				<meta name="twitter:image" content="a" />
			</head>
		</html>
	`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatal(err)
	}

	result := findTwitterCard(doc)

	if len(result) != 1 {
		t.Fatal("Expected 1 result.")
	}

	if result[0] != "a" {
		t.Errorf("Unexpected result: %v", result[0])
	}
}

func TestExtractFromMicrosoftTileImage(t *testing.T) {
	html := `
		<html>
			<head>
				<meta name="msapplication-TileImage" content="a" />
			</head>
		</html>
	`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatal(err)
	}

	result := findMicrosoftTileImage(doc)

	if len(result) != 1 {
		t.Fatal("Expected 1 result.")
	}

	if result[0] != "a" {
		t.Errorf("Unexpected result: %v", result[0])
	}
}

func TestExtractFromAppleTouchIcon(t *testing.T) {
	html := `
		<html>
			<head>
				<link rel="apple-touch-icon" href="a" />
			</head>
		</html>
	`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatal(err)
	}

	result := findAppleTouchIcon(doc)

	if len(result) != 1 {
		t.Fatal("Expected 1 result.")
	}

	if result[0] != "a" {
		t.Errorf("Unexpected result: %v", result[0])
	}
}

func TestExtractFromLink(t *testing.T) {
	html := `
		<html>
			<head>
				<link rel="icon" href="a" />
				<link rel="icon shortcut" href="b" />
			</head>
		</html>
	`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatal(err)
	}

	result := findLink(doc)

	if len(result) != 2 {
		t.Fatal("Expected 2 results.")
	}

	if result[0] != "a" {
		t.Errorf("Unexpected result: %v", result[0])
	}

	if result[1] != "b" {
		t.Errorf("Unexpected result: %v", result[1])
	}
}
