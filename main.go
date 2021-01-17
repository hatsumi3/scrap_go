// go言語　スクレイピングサンプル
// 参考記事
// agouti goquery
// https://qiita.com/tenten0213/items/1f897ff8a64bd8b5270c
// https://qiita.com/libra_lt/items/b1e3ba2c9fc33a23951d
package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
)

var driver *agouti.WebDriver

func configureDriver() {
	driver = agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",
			"--windows-size=1280,800",
		}),
		agouti.Debug,
	)
	// driver = agouti.ChromeDriver()
}

func startDriver(targetURL string) *agouti.Page {
	if err := driver.Start(); err != nil {
		log.Fatal(err)
	}
	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatal(err)
	}

	// URL Setting
	page.Navigate(targetURL)
	return page
}

func scrapyPage() {
	//var targetURL string = "https://b.hatena.ne.jp/hotentry/all"
	var targetURL string = "https://qiita.com/"
	configureDriver()
	page := startDriver(targetURL)
	defer driver.Stop()
	source, err := page.HTML()
	if err != nil {
		log.Fatal("source", err)
	}
	r := strings.NewReader(source)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal("doc", err)
	}
	fmt.Println(source)
	fmt.Println(doc)
	time.Sleep(3 * time.Second)
}

func main() {
	fmt.Println("hello, world!")
	log.Println("start")
	scrapyPage()
	log.Println("end")
}
