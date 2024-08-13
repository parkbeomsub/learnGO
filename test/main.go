package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Request
	url := "http://ncov.mohw.go.kr/"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// HTML 읽기
	html, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 현황판 파싱
	wrapper := html.Find("ul.liveNum")
	items := wrapper.Find("li")

	// items 순회하면서 출력
	items.Each(func(idx int, sel *goquery.Selection) {
		title := sel.Find("strong.tit").Text()
		value := sel.Find("span.num").Text()
		before := sel.Find("span.before").Text()

		fmt.Println(title, value, before)
	})
}
