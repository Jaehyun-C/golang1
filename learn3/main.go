package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://www.jobkorea.co.kr/recruit/joblist?menucode=local&localorder=1"

func main() {
	getPages()
}

func getPages() int {
	res, err := http.Get(baseURL)
	
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each()
	fmt.Println(doc)
	return 0
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