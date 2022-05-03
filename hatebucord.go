package hatebucord

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type ScrapeResult struct {
	Href  string
	Title string
}

func PostHatebu() error {
	result, err := scrapeHatebu()
	if err != nil {
		return err
	}

	endpoint := os.Getenv("DISCORD_WEBHOOK")
	param := fmt.Sprintf(`{
		"username": "今日のはてブ1位",
		"avatar_url": "https://b.st-hatena.com/54e3e2fdaf3b549836dedda8cb7409f9336287d9/images/v4/public/gh-logo@2x.png",
		"content": "%s\r%s"
	}`, result.Title, result.Href)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(param)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	_, err = client.Do(req)
	return err
}

func scrapeHatebu() (*ScrapeResult, error) {
	res, err := http.Get("https://b.hatena.ne.jp/hotentry/it")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, err
	}

	result := document.Find("div.entrylist-contents-main").First().Find("h3.entrylist-contents-title").Find("a")
	href, _ := result.Attr("href")
	title, _ := result.Attr("title")
	return &ScrapeResult{Href: href, Title: title}, nil
}
