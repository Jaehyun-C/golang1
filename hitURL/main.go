package main

import (
	"errors"
	"fmt"
	"net/http"
)

//에러 변수 생성
var errRequestFailed = errors.New("Request failed")


func main() {
	//map 초기화 방법
	results := map[string]string{}
	//슬라이스
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
		result := "OK"
		err := hitURL(url)
		if err != nil{
			result = "Failed"
		}
		results[url] = result
	}
	//map의 range 형태
	for url, result := range(results){
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {

	fmt.Println("Checking url : ",url)
	//url 체크 메서드
	resp, err := http.Get(url)
	//statusCode 방법
	if err != nil || resp.StatusCode >= 400{
		fmt.Println(err)
		return errRequestFailed
	}
	return nil
}