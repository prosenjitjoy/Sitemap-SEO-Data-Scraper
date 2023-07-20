package main

import (
	"fmt"
	"main/config"
	"main/sitemap"
	"main/utils"
)

var conf = config.LoadData()

func main() {
	pagesURL := sitemap.Crawler(conf.SitemapURL, conf.CrawlerConcurrency)
	results := sitemap.Scraper(pagesURL, conf.ScraperConcurrency)

	utils.OutputCSV(results)
	utils.OutputJSON(results)

	fmt.Println("🎉 Done!")
}
