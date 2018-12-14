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
	base, err := url.Parse(u)
	if err != nil {
		return result, err
	}
	res, err := http.Get(u)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return result, errors.New(res.Status)
	}
	if (res.Request.Response != nil) {
		base, err = res.Request.Response.Location()
		if err != nil {
			return result, err
		}
	}
	result, err = FromReader(res.Body)
	if err != nil {
		return result, err
	}
	result = append(result, "/favicon.ico")
	icons := []string{}
	for _, i := range result {
		u, err := url.Parse(i)
		if err != nil {
			continue
		}
		if !u.IsAbs() {
			if u.Host != "" {
				u.Scheme = base.Scheme
				i = u.String()
			} else {
				i = base.ResolveReference(u).String()
			}
		}
		icons = append(icons, i)
	}
	return icons, nil
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
