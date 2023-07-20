package model

type ConfigData struct {
	SitemapURL         string `yaml:"sitemap_url"`
	CrawlerConcurrency int    `yaml:"crawler_concurrency"`
	ScraperConcurrency int    `yaml:"scraper_concurrency"`
	RequestTimeout     int    `yaml:"request_timeout"`
}

type SEOdata struct {
	StatusCode  int
	URL         string
	Title       string
	Heading     string
	Description string
}
