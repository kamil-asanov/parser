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

	collector.OnHTML(".vacancy-serp-item-body__main-info", func(element *colly.HTMLElement) {
		temp := vacancy{}
		temp.Title = element.ChildText(".serp-item__title")
		temp.URL = element.ChildAttr(".serp-item__title", "href")
		temp.Company = element.ChildText(".bloko-link bloko-link_kind-tertiary")
		temp.Salary = element.ChildText(".bloko-header-section-2")
		vacancies = append(vacancies, temp)

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
	collector.Visit("https://hh.ru/search/vacancy?search_field=name&search_field=company_name&search_field=description&enable_snippets=false&L_save_area=true&experience=between1And3&professional_role=160&schedule=remote&text=DevOps")

	collector.Wait()
	fmt.Println(vacancies)
}
