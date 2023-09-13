package main

import (
	"fmt"
	"time"

	//"encoding/csv"
	"github.com/gocolly/colly"
	//"log"
	//"os"
)

type vacancy struct {
	Company string
	Salary  string
	URL     string
	Title   string
}

func main() {
	vacancies := []vacancy{}
	collector := colly.NewCollector(
		colly.AllowedDomains("hh.ru"),
		colly.Async(true),
	)

	collector.OnHTML("div[vacancy-serp-item-body__main-info]", func(element *colly.HTMLElement) {
		temp := vacancy{}
		temp.Title = element.ChildText("a[serp-item__title]")
		temp.URL = element.ChildAttr("a[serp-item__title]", "href")
		temp.Company = element.ChildText("a[bloko-link bloko-link_kind-tertiary]")
		temp.Salary = element.ChildText("span[bloko-header-section-2]")
		fmt.Println(temp)
		vacancies = append(vacancies, temp)

		// links := element.Attr("href")
		// fmt.Println("Link found: %q -> %s\n", element.Text, links)

		// collector.Visit(element.Request.AbsoluteURL(links))
	})

	collector.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	collector.OnError(func(request *colly.Response, err error) {
		fmt.Println("Request URL:", request.Request.URL, "failed with response:", request, "\nError:", err)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})
	collector.Visit("https://hh.ru/search/vacancy?area=1&experience=between1And3&search_field=name&search_field=company_name&search_field=description&text=DevOps&enable_snippets=false&L_save_area=true")

	collector.Wait()
	fmt.Println(vacancies)
}
