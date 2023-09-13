package main

import (
	"fmt"
	//"encoding/csv"
	"github.com/gocolly/colly"
	//"log"
	//"os"
)

func main() {
	collector := colly.NewCollector(
		colly.AllowedDomains("ru.wikipedia.org"),
	)

	collector.OnHTML("a[href]", func(element *colly.HTMLElement) {
		links := element.Attr("href")
		fmt.Println("Link found: %q -> %s\n", element.Text, links)

		collector.Visit(element.Request.AbsoluteURL(links))
	})

	collector.OnError(func(request *colly.Response, err error) {
		fmt.Println("Request URL:", request.Request.URL, "failed with response:", request, "\nError:", err)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})
	collector.Visit("https://ru.wikipedia.org/wiki/%D0%9D%D0%BE%D0%B2%D0%B0%D1%8F_%D0%BD%D0%B0%D1%86%D0%B8%D0%BE%D0%BD%D0%B0%D0%BB%D1%8C%D0%BD%D0%B0%D1%8F_%D0%B3%D0%B0%D0%BB%D0%B5%D1%80%D0%B5%D1%8F")
}
