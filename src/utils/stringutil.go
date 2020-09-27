package utils

import (
	"regexp"
	"strings"
)

/**
 * Author: gageshan
 * Date: 2020/9/27
 * Time: 16:32
 */

/**
 * 去除字符串中空格
 */
func Trim(text string) string {
	r := regexp.MustCompile("\\s+")
	replace := r.ReplaceAllString(text, " ")
	replace = strings.ReplaceAll(replace, " ", "")
	return replace
}
