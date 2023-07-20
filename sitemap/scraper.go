package sitemap

import (
	"log"
	"main/model"
	"main/utils"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Scraper(pages []string, concurrency int) []model.SEOdata {
	semaphore := make(chan struct{}, concurrency)

	worklist := make(chan []string)
	go func() { worklist <- pages }()

	results := []model.SEOdata{}
	for n := 1; n > 0; n-- {
		links := <-worklist
		for _, link := range links {
			if link != "" {
				n++
				go func(link string) {
					log.Println("Requesting URL:", link)
					semaphore <- struct{}{}
					resp, err := utils.MakeRequest(link)
					<-semaphore
					if err != nil {
						log.Fatal("Failed to get a response:", err)
					}

					data, err := getSEOdata(resp)
					if err != nil {
						log.Fatal("Failed to parse response body:", err)
					}

					results = append(results, data)
					worklist <- []string{}
				}(link)
			}
		}
	}

	return results
}

func getSEOdata(resp *http.Response) (model.SEOdata, error) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return model.SEOdata{}, err
	}

	result := model.SEOdata{}
	result.StatusCode = resp.StatusCode
	result.URL = resp.Request.URL.String()
	result.Title = doc.Find("title").First().Text()
	result.Heading = doc.Find("h1").First().Text()
	result.Description, _ = doc.Find("meta[name^=description]").Attr("content")

	return result, nil
}
