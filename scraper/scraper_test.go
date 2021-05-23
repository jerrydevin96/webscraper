package scraper

import (
	"testing"
)

func TestVersionHandler(t *testing.T) {
	getHTMLVersionHandler("https://medium.com/")
}
