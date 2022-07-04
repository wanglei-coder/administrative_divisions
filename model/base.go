package model

import (
	"administrative_divisions/pkg/request"
	"github.com/PuerkitoBio/goquery"
	"github.com/anaskhan96/soup"
)

type Base struct {
	URL                 string
	Name                string
	StatisticalAreaCode string
	Client              *request.CustomHTTPClient
}

func (b *Base) GetDocument() (*goquery.Document, error) {
	return b.Client.GetDocument(b.URL)
}

func (b *Base) GetSoupDocument() (soup.Root, error) {
	return b.Client.GetSoupDocument(b.URL)
}
