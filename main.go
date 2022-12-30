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
	id 			string
	title		string
	location 	string
	salary		string
	stack		string
}

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?searchword=%ED%8C%8C%EC%9D%B4%EC%8D%AC&go=&flag=n&searchMode=1&searchType=search&search_done=y&search_optional_item=n&recruitSort=relation&recruitPageCount=40&inner_com_type=&company_cd=0%2C1%2C2%2C3%2C4%2C5%2C6%2C7%2C9%2C10&show_applied=&quick_apply=&except_read=&ai_head_hunting="
var detailURL string = "https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&recommend_ids=eJxNkMkRA0EIA6PxH8QheDsQ55%2BFZ7xrZp9dalGAA8gIfgr64ts9Fd35Kc2Nkg31hegrtXLWP3Wz1tx4dQkmOTLNV3kwrIMPWYqFByKrz%2BTOUg5GLDkGfa0lGEzz%2BKV6y2Lco%2BTeKoNHNpg5j1wUObIL102TZqBif0O%2Bi%2FpAcw%3D%3D&view_type=search&searchword=%ED%8C%8C%EC%9D%B4%EC%8D%AC&searchType=search&gz=1&t_ref_content=generic&t_ref=search&paid_fl=n&search_uuid=9390d19c-8703-43f7-8ba6-47b1d36862e3"
// &rec_idx=44621811

func main() {
	var jobs []extractedJob

	totalPages := getPages()
	fmt.Println(totalPages)

	for i:=1;i<=totalPages;i++ {
		// [ [] + [] + [] + [] ]
		extractedJobs := getPage(i)
		// 이거 되는 거였어?
		jobs = append(jobs, extractedJobs...)
	}

	// fmt.Println(jobs)
	writingJob(jobs)
}

func writingJob(jobs []extractedJob) {
	file, err := os.Create("data/jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	// 함수가 끝나는 시점에 CSV 파일에 데이터를 추가하는 함수
	defer w.Flush()

	headers := []string{"ID", "TITLE", "LOCATION", "SALARY", "STACK"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range(jobs) {
		jobSlice := []string{job.id, job.title, job.location, job.salary, job.stack}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}

	fmt.Println("Done, extracted", len(jobs), )
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob

	pageUrl := baseURL + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageUrl)

	resp, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(resp)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	searchCards  := doc.Find("div.item_recruit")
	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})

	return jobs
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("value")
	title := cleanString(card.Find("h2.job_tit span").Text())
	location := cleanString(card.Find("div.job_condition span:nth-child(1)").Text())
	salary := cleanString(card.Find("div.job_condition span:nth-child(5)").Text())
	stack := cleanString(card.Find("div.job_sector").Text())
	return extractedJob{
		id: detailURL + "&rec_idx=" + id + "#seq=0", 
		title: title, 
		location: location, 
		salary: salary, 
		stack: stack,
	}
}

func cleanString(str string) string {
	// Space를 없애고, Fields로 분리
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
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

