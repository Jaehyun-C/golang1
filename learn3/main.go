package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

var baseURL string = "https://kr.indeed.com/jobs?q=%EB%B8%94%EB%A1%9D%EC%B2%B4%EC%9D%B8&l&from=searchOnHP&vjk=8f59e8696a325187"

func main() {
	totalPage:= getPages()
	for i := 0; i < totalPage; i++ {
		getPage(i)
	}
}

func getPage(page int){
	pageURL := baseURL + "&start=" + strconv.Itoa(page * 50)
	fmt.Println("Requesting", pageURL)
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection){
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil{
		log.Fatalln(err)
	}
	
}

func checkCode(res *http.Response){
	if res.StatusCode != 200{
		log.Fatalln("Request failed with Status : ",res.StatusCode)
	}
}