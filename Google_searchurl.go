package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 输出图案
	fmt.Println(`
BBBB   AAAAA  III  
B   B  A   A   I   
BBBB   AAAAA   I   
B   B  A   A   I   
BBBB   A   A  III  `)
	fmt.Println("by:【白】")
	var searchKeyword string
	var urlCount int
	var proxyAddress string

	// 接受用户输入的搜索内容
	fmt.Print("\n输入你搜索的关键词: ")
	reader := bufio.NewReader(os.Stdin)
	searchKeyword, _ = reader.ReadString('\n')
	searchKeyword = strings.TrimSpace(searchKeyword)

	// 接受用户输入的要爬取的 URL 数量
	fmt.Print("输入你要爬取的url数量: ")
	fmt.Scanln(&urlCount)

	// 接受用户输入的代理地址，设置默认代理地址为 http://127.0.0.1:10809
	fmt.Print("设置你的代理地址 (默认代理地址： http://127.0.0.1:10809): ")
	reader = bufio.NewReader(os.Stdin)
	proxyAddress, _ = reader.ReadString('\n')
	proxyAddress = strings.TrimSpace(proxyAddress)
	if proxyAddress == "" {
		proxyAddress = "http://127.0.0.1:10809"
	}

	// 设置代理
	proxyURL, err := url.Parse(proxyAddress)
	if err != nil {
		fmt.Println("Error parsing proxy URL:", err)
		return
	}
	http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}

	// 构建谷歌搜索的 URL
	searchURL := fmt.Sprintf("https://www.google.com/search?q=%s&num=%d", url.QueryEscape(searchKeyword), urlCount)

	// 发送 HTTP 请求并获取响应
	response, err := http.Get(searchURL)
	if err != nil {
		fmt.Println("Error fetching search results:", err)
		return
	}
	defer response.Body.Close()

	// 使用 goquery 解析 HTML
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	// 提取搜索结果中的 URL
	var urls []string
	urlMap := make(map[string]bool) // 用于存放已经存在的 URL，以去重
	doc.Find("a[href]").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists && strings.HasPrefix(href, "/url?q=") {
			re := regexp.MustCompile(`^/url\?q=(.+?)&`)
			match := re.FindStringSubmatch(href)
			if len(match) == 2 {
				decodedURL, err := url.QueryUnescape(match[1])
				if err != nil {
					decodedURL = match[1]
				}
				if !urlMap[decodedURL] {
					urlMap[decodedURL] = true
					urls = append(urls, decodedURL)
				}
			}
		}
	})

	// 限制爬取的 URL 数量
	if len(urls) > urlCount {
		urls = urls[:urlCount]
	}

	// 将 URL 写入文本文件
	file, err := os.Create("urls.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, u := range urls {
		_, err := file.WriteString(u + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("\nURLs have been saved to urls.txt")
}
