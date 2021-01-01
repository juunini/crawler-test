package crawler_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCrawlerTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CrawlerTest Suite")
}
