package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?searchword=%ED%8C%8C%EC%9D%B4%EC%8D%AC&go=&flag=n&searchMode=1&searchType=search&search_done=y&search_optional_item=n&recruitPage=1&recruitSort=relation&recruitPageCount=40&inner_com_type=&company_cd=0%2C1%2C2%2C3%2C4%2C5%2C6%2C7%2C9%2C10&show_applied=&quick_apply=&except_read=&ai_head_hunting="
var originalUrl string = "https://www.saramin.co.kr/zf_user/search/recruit?searchword=%ED%8C%8C%EC%9D%B4%EC%8D%AC&go=&flag=n&searchMode=1&searchType=search&search_done=y&search_optional_item=n&recruitSort=relation&recruitPageCount=40&inner_com_type=&company_cd=0%2C1%2C2%2C3%2C4%2C5%2C6%2C7%2C9%2C10&show_applied=&quick_apply=&except_read=&ai_head_hunting="

func main() {
	totalPages := getPages()
	fmt.Println(totalPages)

	for i:=1;i<=totalPages;i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageUrl := originalUrl + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageUrl)
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	// 함수 종료 후 body stream close
	defer res.Body.Close()

	doc,err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// doc.Find("span.page, a.page > span").Each(func(i int, s *goquery.Selection) {
	// 	fmt.Println(s.Html())
	// 	// 중첩 selecting 가능
	// 	// fmt.Println(s.Find("a"))
	// })
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		// fmt.Println(s.Find("span.page, a.page").Length())
		pages = s.Find("span.page, a.page").Length()
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
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}