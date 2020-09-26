package extractor

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
)

/**
 * Author: gageshan
 * Date: 2020/9/26
 * Time: 17:22
 */

var prefix string = "https://www.cnvd.org.cn/"

func IndexFromUrl(typeId int, max int, offset int) []string {
	var url string = mapToUrl(typeId, max, offset)
	log.Println("parser " + url)
	resp, err := http.Post(url, "text/html;charset=UTF-8", nil)
	var x []string
	if err != nil {
		log.Fatal(err)
		return x
	}
	defer resp.Body.Close()
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
		return x
	}

	document.Find(".tlist tr a").Each(func(i int, s *goquery.Selection) {
		key, exists := s.Attr("href")
		if !exists {
			log.Fatal(exists)
		}
		x = append(x, prefix+key)
	})
	log.Println(x)
	return x
}

func mapToUrl(typeId int, max int, offset int) string {

	return prefix + "flaw/typeResult?typeId=" +
		strconv.Itoa(typeId) + "&max=" +
		strconv.Itoa(max) + "&offset=" +
		strconv.Itoa(offset)
}
