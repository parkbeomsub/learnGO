package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"

type extractJobs struct {
	id       string
	title    string
	location string
	summary  string
	company  string
}

func main() {
	totalPage := getPages()
	fmt.Println(totalPage)
	var jobs [][]extractJobs
	for i := 1; i <= totalPage; i++ {
		k := getPage(totalPage)
		jobs = append(jobs, k)
	}
	fmt.Println(jobs)
	

}

func extractJob(card *goquery.Selection) extractJobs {
	id, _ := card.Attr("value")
	title := claanString(card.Find(".job_tit>a").Text())
	location := claanString(card.Find(".job_condition>span>a").Text())
	summary := claanString(card.Find(".job_sector").Clone().ChildrenFiltered(".job_day").Remove().End().Text())
	company := claanString(card.Find(".area_corp>strong>a").Text())
	//fmt.Println(id, title, location, summary, company)

	return extractJobs{
		id:       id,
		title:    title,
		location: location,
		summary:  summary,
		company:  company,
	}
}

func getPage(page int) []extractJobs {
	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println("Requesting :", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkStatusCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	var jobs []extractJobs
	doc.Find(".item_recruit").Each(func(i int, card *goquery.Selection) {
		// 값 추출
		// id, _ := s.Attr("value")
		// title := s.Find(".job_tit>a").Text()
		// condition := s.Find(".job_condition").Text()
		// fmt.Println(id, title, condition)
		job := extractJob(card)
		jobs = append(jobs, job)

	})
	return jobs
}

func claanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), "")
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkStatusCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// doc.Find(".area_job").Each(func(i int, s *goquery.Selection) {
	// 	// For each item found, get the title
	// 	title := s.Find("a").Text()
	// 	fmt.Printf("value is  :  %s\n", title)
	// })
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatusCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode, res.Status)

	}

}

func writeJobs(jobs []extractJobs) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush() // must

	headers := []string{"Id", "Title", "Location", "Summary", "Company"}

	Werr := w.Write(headers)
	checkErr(Werr)

	for _, job := range jobs {
		jobSlice := []string{job.id, job.title, job.location, job.summary, job.company}
		jobErr := w.Write(jobSlice)
		checkErr(jobErr)
	}
}
