package html

import (
	"fmt"
	"regexp"
	"strings"
)

func RemoveStyle(content string,tagArray []string) string {
	reg, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	for _, style := range tagArray {
		expr := "\\<" + style + "[\\S\\s]+?\\</" + style + "\\>"
		reg, _ = regexp.Compile(expr)
		content = reg.ReplaceAllString(content, "")
	}
	return content
}
func FetchText(html string, list []string) string {
	for _, v := range list {
		style := "<" + v + "[^>]+>"
		reg := regexp.MustCompile(style)
		html = reg.ReplaceAllString(html, "")

		style = "</" + v + ">"
		reg = regexp.MustCompile(style)
		html = reg.ReplaceAllString(html, "")
	}

	return html
}
func GetUrlMap(content string) (map[string]string){
	regx,err := regexp.Compile("<img.+?src=[\"'](.+?)?\"")
	if err != nil {
		fmt.Printf("regexp Compile error:" + err.Error())
		return nil
	}

	urlMap := make(map[string]string)
	urls := regx.FindAllString(content, -1)
	for _, url := range urls {
		newurl := strings.Replace(url, "src=\"//", "src=\"http://", -1)
		urlMap[url] = newurl
	}
	return urlMap
}