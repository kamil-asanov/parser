package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func Parse(site string, link string) {
	//Vacancies := []Vacancy{}
	collector := colly.NewCollector(
		colly.AllowedDomains(site),
		colly.Async(true),
	)

	collector.OnHTML(".vacancy-serp-item-body__main-info", func(element *colly.HTMLElement) {
		temp := Vacancy{}
		temp.Title = element.ChildText(".serp-item__title")
		temp.URL = element.ChildAttr(".serp-item__title", "href")
		temp.Company = element.ChildText(".bloko-link bloko-link_kind-tertiary")
		temp.Salary = element.ChildText(".bloko-header-section-2")
		Vacancies = append(Vacancies, temp)

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
	collector.Visit(link)

	collector.Wait()
}
