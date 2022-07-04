package request

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/anaskhan96/soup"
	"go.uber.org/ratelimit"
	"io"
	"net/http"
	"time"
)

var DefaultCustomHTTPClient = NewCustomHTTPClient(WithLimiter(ratelimit.NewUnlimited()))

type CustomHTTPClient struct {
	Limiter ratelimit.Limiter
}

func WithLimiter(limiter ratelimit.Limiter) CustomRequestOption {
	return func(options *CustomHTTPClient) {
		options.Limiter = limiter
	}
}

type CustomRequestOption func(options *CustomHTTPClient)

func NewCustomHTTPClient(opts ...CustomRequestOption) *CustomHTTPClient {
	var request CustomHTTPClient
	for _, opt := range opts {
		opt(&request)
	}

	if request.Limiter == nil {
		request.Limiter = ratelimit.NewUnlimited()
	}
	return &request
}

func (r *CustomHTTPClient) GetDocument(url string) (*goquery.Document, error) {
	r.Limiter.Take()
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	return goquery.NewDocumentFromReader(resp.Body)
}

func (r *CustomHTTPClient) GetSoupDocument(url string) (soup.Root, error) {
	r.Limiter.Take()
	resp, err := soup.GetWithClient(url, &http.Client{Timeout: time.Second * 10})
	if err != nil {
		return soup.Root{}, err
	}
	return soup.HTMLParse(resp), nil
}

func (r *CustomHTTPClient) Get(url string) (*http.Response, error) {
	r.Limiter.Take()
	return http.Get(url)
}

func (r *CustomHTTPClient) SoupGet(url string) (string, error) {
	r.Limiter.Take()
	return soup.Get(url)
}
