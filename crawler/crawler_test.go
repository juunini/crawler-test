package crawler_test

import (
	"crawler_sample/crawler"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetProducts", func() {
	It("should returns url list", func() {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Set("Content-Type", "text/html")
			rw.Write([]byte(productsPageDocument))
		}))
		defer server.Close()

		productsURL := crawler.GetProducts(server.URL)

		Expect(productsURL[0]).To(Equal("/sample/1"))
		Expect(productsURL[1]).To(Equal("/sample/2"))
	})
})
