package main

import (
	"errors"
	"fmt"
	"net/http"
)

//에러 변수 생성
var errRequestFailed = errors.New("Request failed")
 
type result struct{
	url string
	status string
}

func main() {
	resultMap := make(map[string]string)
	// 채널 선언
	c := make(chan result)
	// 슬라이스
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}
	//슬라이스의 range 형태, 슬라이스의 첫번째 값은 인덱스
	//에러 컨트롤
	for _,url := range urls {
		go hitURL(url, c)
	}
	for i:=0;i<len(urls);i++{
		results := <-c
		resultMap[results.url] = results.status
	}

	for url, status := range resultMap{
		fmt.Println(url, status)
	}
}
// c chan<-result 는 전송전용
func hitURL(url string, c chan<- result){
	//url 체크 메서드
	resp, err := http.Get(url)
	status := "OK"
	//statusCode 방법
	if err != nil || resp.StatusCode >= 400{
		status = "Failed"
	} 
	//채널에 데이터 전송
	c <- result{url:url, status:status}
}