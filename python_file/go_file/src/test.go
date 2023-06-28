package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	_ "strings"
	"sync"
	"time"
)

type Comment struct {
	Email string `json:"email"`
}

func main() {
	start := time.Now()

	urls := generateURLs(100)
	emails := make(chan string, 500)
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchEmails(url, &wg, emails)
	}

	wg.Wait()
	close(emails)

	saveEmails(emails)

	elapsed := time.Since(start)
	fmt.Printf("爬虫完成！总共耗时：%s\n", elapsed)
}

func generateURLs(count int) []string {
	urls := make([]string, count)
	baseURL := "https://jsonplaceholder.typicode.com/posts/%d/comments"
	for i := 0; i < count; i++ {
		urls[i] = fmt.Sprintf(baseURL, i+1)
	}
	return urls
}

func fetchEmails(url string, wg *sync.WaitGroup, emails chan<- string) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("无法获取URL：%s, 错误：%v\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应体失败：%v\n", err)
		return
	}

	var comments []Comment
	if err := json.Unmarshal(body, &comments); err != nil {
		fmt.Printf("解析JSON失败：%v\n", err)
		return
	}

	for _, comment := range comments {
		emails <- comment.Email
	}
}

func saveEmails(emails <-chan string) {
	file, err := os.Create("emails.txt")
	if err != nil {
		fmt.Printf("创建文件失败：%v\n", err)
		return
	}
	defer file.Close()

	for email := range emails {
		file.WriteString(email + "\n")
	}
}
