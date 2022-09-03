package dhttp

import (
	"bytes"
	"context"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/reusee/e5"
)

func TestReq(t *testing.T) {
	defer he(nil, e5.TestingFatal(t))

	scope := Req(context.Background(), "GET", "https://qq.com")

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
