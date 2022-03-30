package dhttp

import (
	"bytes"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/reusee/e4"
)

func TestReq(t *testing.T) {
	defer he(nil, e4.TestingFatal(t))

	scope := Req("GET", "https://qq.com")

	var content Content
	scope.Assign(&content)
	if !bytes.Contains(content, []byte("html")) {
		t.Fatal()
	}

	var doc *goquery.Document
	scope.Assign(&doc)
	if doc == nil {
		t.Fatal()
	}

}
