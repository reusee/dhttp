package dhttp

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
)

func (Def) GoqueryDocument(
	content []byte,
) (
	doc *goquery.Document,
) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
	ce(err)
	return doc
}
