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

type extractedJob struct {
	id       string
	title    string
	location string
	summary  string
	company  string
}

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python&recruitPage="

func main() {

	totalPages := getPages()
	fmt.Println(totalPages)
	for i := 1; i <= totalPages; i++ {
		getPage(i)
	}

}

func getPage(page int) []extractedJob {
	pageUrl := "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python&recruitPage=" + strconv.Itoa(page)
	fmt.Println("Requesting " + pageUrl)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".area_job")

	var jobs []extractedJob
	doc.Find(".item_recruit").Each(func(i int, card *goquery.Selection) {
		job := extractedJob(card)
		jobs = append(jobs, job)
	})
	return jobs

}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), "")
}

func extractJob(card *goquery.Selection) {
	id, _ := card.Attr("value")
	title := cleanString(card.Find(".job_tit>a").Text())
	location := cleanString(card.Find(".job_condition>span>a").Text())
	summary := cleanString(card.Find(".job_sector").Clone().ChildrenFiltered(".job_day").Remove().End().Text())
	company := cleanString(card.Find(".area_corp>strong>a").Text())
	fmt.Println(id, title, location, summary, company)
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("span").Length()

	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request Faild with Status: ", res.StatusCode)
	}

}
func writeJobs(jobs []extractedJob) {
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
