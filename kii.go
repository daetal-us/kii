package kii

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func FromURL(u string) (result []string, err error) {
	res, err := http.Get(u)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return result, errors.New(res.Status)
	}
	result, err = FromReader(res.Body)
	if err != nil {
		return result, err
	}
	f, err := url.Parse("/favicon.ico")
	if err != nil {
		return result, err
	}
	base, err := url.Parse(u)
	if err != nil {
		return result, err
	}
	favicon := base.ResolveReference(f).String()
	res, err = http.Get(favicon)
	if err != nil {
		return result, err
	}
	if res.StatusCode == http.StatusOK {
		result = append(result, favicon)
	}
	return result, nil
}

func FromHTML(html string) (result []string, err error) {
	return FromReader(strings.NewReader(html))
}

func FromReader(r io.Reader) (results []string, err error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return results, err
	}

	results = append(results, findJSONLD(doc)...)
	results = append(results, findOpenGraph(doc)...)
	results = append(results, findLink(doc)...)
	results = append(results, findAppleTouchIcon(doc)...)
	results = append(results, findTwitterCard(doc)...)
	results = append(results, findMicrosoftTileImage(doc)...)

	return results, nil
}

func findLink(doc *goquery.Document) (result []string) {
	selector := `link[rel~="icon"]`
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		content, set := s.Attr("href")
		if set && content != "" {
			result = append(result, content)
		}
	})
	return result
}

func findOpenGraph(doc *goquery.Document) (result []string) {
	selector := `meta[property="og:image"]`
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		content, set := s.Attr("content")
		if set && content != "" {
			result = append(result, content)
		}
	})
	return result
}

func findTwitterCard(doc *goquery.Document) (result []string) {
	selector := `meta[name="twitter:image"]`
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		content, set := s.Attr("content")
		if set && content != "" {
			result = append(result, content)
		}
	})
	return result
}

func findMicrosoftTileImage(doc *goquery.Document) (result []string) {
	selector := `meta[name="msapplication-TileImage"]`
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		content, set := s.Attr("content")
		if set && content != "" {
			result = append(result, content)
		}
	})
	return result
}

func findAppleTouchIcon(doc *goquery.Document) (result []string) {
	selector := `link[rel="apple-touch-icon"]`
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		content, set := s.Attr("href")
		if set && content != "" {
			result = append(result, content)
		}
	})
	return result
}

func findJSONLD(doc *goquery.Document) (result []string) {
	selector := `script[type="application/ld+json"]`
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		d := &struct {
			Logo  string `json:"logo"`
			Image string `json:"image"`
		}{}
		err := json.Unmarshal([]byte(s.Text()), d)
		if err == nil {
			if d.Logo != "" {
				result = append(result, d.Logo)
			}
			if d.Image != "" {
				result = append(result, d.Image)
			}
		}
	})
	return result
}
