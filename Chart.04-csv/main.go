package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
)

func main() {

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

