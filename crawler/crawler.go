package crawler

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// GetProducts returns products url
func GetProducts(url string) (productsURL []string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		panic(err)
	}

	productsURL = doc.Find("#mp-collection-grid > div").Map(func(i int, product *goquery.Selection) string {
		href, _ := product.Find("a").First().Attr("href")
		return href
	})
	return
}
