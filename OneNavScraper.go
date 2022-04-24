package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	client := &http.Client{}
	begin := time.Now()
	url := "https://jav.iuo.ink/"
	file, _ := os.OpenFile("README.md", os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36")
	request.Header.Add("Host", "jav.iuo.ink")
	response, _ := client.Do(request)
	opending, _ := goquery.NewDocumentFromReader(response.Body)
	ele := opending.Find(".list")
	ele.Each(func(i int, s *goquery.Selection) {
		dataUrl, ok := s.Attr("data-url")
		if ok {
			name := strings.TrimSpace(s.Find(".name").Text())
			desc := strings.TrimSpace(s.Find(".desc").Text())
			file.WriteString("==========\n\n")
			file.WriteString("链接：" + dataUrl + "\n\n")
			file.WriteString("名称：" + name + "\n\n")
			file.WriteString("描述：" + desc + "\n\n")
			file.WriteString("\n\n")
		}
	})
	end := time.Now()
	spendTime := end.Sub(begin)
	fmt.Println("花费时间为:", spendTime)
}
