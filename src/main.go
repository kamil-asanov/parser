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
        colly.AllowedDomains("habr.com"),
    )
	 collector.OnHTML(".mw-parser-output", func(element *colly.HTMLElement) {
        links := element.ChildAttrs("a", "href")
        fmt.Println(links)
    })
    collector.Visit("habr.com")
}

