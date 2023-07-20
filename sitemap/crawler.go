package sitemap

import (
	"log"
	"main/utils"

	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Crawler(sitemapURL string, concurrency int) []string {
	semaphore := make(chan struct{}, concurrency)

	worklist := make(chan []string)
	go func() { worklist <- []string{sitemapURL} }()

	toCrawl := []string{}
	for n := 1; n > 0; n-- {
		links := <-worklist
		for _, link := range links {
			n++
			go func(link string) {
				semaphore <- struct{}{}
				resp, err := utils.MakeRequest(link)
				<-semaphore
				if err != nil {
					log.Fatal("Failed to get a response:", err)
				}

				links, err := extractURLs(resp)
				if err != nil {
					log.Fatal("Failed to parse response body:", err)
				}

				sitemapLinks, pageLinks := utils.IsSitemap(links)
				if sitemapLinks != nil {
					worklist <- sitemapLinks
				}

				toCrawl = append(toCrawl, pageLinks...)
			}(link)
		}
	}

	return toCrawl
}

func extractURLs(response *http.Response) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	foundURLs := []string{}
	doc.Find("loc").Each(func(i int, s *goquery.Selection) {
		URL := s.Text()
		if URL != "" {
			foundURLs = append(foundURLs, URL)
		}
	})

	return foundURLs, nil
}
