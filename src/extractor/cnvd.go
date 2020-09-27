package extractor

import (
	"encoding/json"
	stringutil "gcrawler/src/utils"
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

type Message struct {
	Title   string `json:"title"`
	Time    string `json:"time"`
	Url     string `json:"url"`
	Content string `json:"content"`
}

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

func DataFromIndexUrl(url string) string {
	resp, err := http.Get(url)
	var mp map[string]string
	mp = make(map[string]string)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer resp.Body.Close()
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	document.Find(".Main .blkContainer " +
		".blkContainerPblk .blkContainerSblk " +
		".tableDiv .gg_detail tr").Each(func(i int, tr *goquery.Selection) {
		var key string
		var value string
		tr.Find("td").Each(func(i int, td *goquery.Selection) {
			text := td.Text()
			text = stringutil.Trim(text)
			if 0 == i {
				key = text
			} else {
				value = text
			}
		})
		if value != "" {
			mp[key] = value
		}
	})

	title := document.Find(".Main .blkContainer .blkContainerPblk .blkContainerSblk h1").Text()
	marshal, _ := json.Marshal(mp)
	message := Message{title, url, mp["公开日期"], string(marshal)}
	mes, _ := json.Marshal(message)
	return string(mes)
}
