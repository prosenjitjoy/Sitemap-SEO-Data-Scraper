package utils

import (
	"fmt"
	"main/config"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var UserAgent = []string{
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:42.0) Gecko/20100101 Firefox/42.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.59",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safari/537.36 OPR/38.0.2220.41",
}

var conf = config.LoadData()

func MakeRequest(targetURL string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * time.Duration(conf.RequestTimeout),
	}

	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", RandomUserAgent())
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func RandomUserAgent() string {
	seed := time.Now().Unix()
	r := rand.New(rand.NewSource(seed))
	randomNum := r.Int() % len(UserAgent)
	return UserAgent[randomNum]
}

func IsSitemap(links []string) ([]string, []string) {
	sitemapLinks := []string{}
	pageLinks := []string{}

	for _, link := range links {
		if strings.Contains(link, "xml") {
			fmt.Println("Found sitemap:", link)
			sitemapLinks = append(sitemapLinks, link)
		} else {
			pageLinks = append(pageLinks, link)
		}
	}

	return sitemapLinks, pageLinks
}
