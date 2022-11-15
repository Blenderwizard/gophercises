package scrapper

import (
	"net/http"
	"strings"
)

func finder(url string, arr []Link) bool {
	for _, x := range arr {
		if x.HREF == url {
			return true
		}
	}
	return false
}

func builder(url string, newPath string) string {
	var pos int
	var char rune
	for pos, char = range url {
		if char == '/' && pos > 7 {
			break
		}
	}
	url = url[:pos]
	return url + newPath
}

func getPathFromUrl(url string) string {
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "https://")
	var pos int
	var char rune
	for pos, char = range url {
		if char == '/' {
			break
		}
	}
	url = url[pos:]
	return url
}

func Scrapper(url string, depth int) []Link {
	if depth == 0 {
		return nil
	}
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var ret []Link
	data := HtmlParser(resp.Body)
	for _, d := range data {
		if strings.HasPrefix(d.HREF, "/") {
			d.HREF = builder(url, d.HREF)
			if d.HREF == url {
				if !(finder(d.HREF, ret)) {
					ret = append(ret, d)
				}
			} else {
				if !(finder(d.HREF, ret)) {
					ret = append(ret, d)
					dam := Scrapper(d.HREF, depth-1)
					for _, e := range dam {
						if !(finder(e.HREF, ret)) {
							ret = append(ret, e)
						}
					}
				}
			}
		} else if strings.HasPrefix(d.HREF, "http://") || strings.HasPrefix(d.HREF, "https://") {
			if !(finder(d.HREF, ret)) {
				ret = append(ret, d)
			}
		}
	}
	return ret
}
