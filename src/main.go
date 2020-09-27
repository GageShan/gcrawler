package main

import (
	"fmt"
	cnvd "gcrawler/src/extractor"
)

/**
 * Author: gageshan
 * Date: 2020/9/26
 * Time: 10:56
 */

func main() {
	urls := cnvd.IndexFromUrl(33, 20, 20)
	for _, i := range urls {
		Json := cnvd.DataFromIndexUrl(i)
		fmt.Println(Json)
	}
}
