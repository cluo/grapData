package main

import (
	"bytes"
	"encoding/csv"
	. "fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	// rf "reflect"
)

func main() {
	doc, err := goquery.NewDocument("http://1.guotie.sinaapp.com/?p=501")
	if err != nil {
		Println("error")
		return
	}
	meta, _ := doc.Find("meta[name='generator']").Attr("content")
	div := doc.Find("div#post-501")
	h1 := div.Find("h1").Text()
	/*
	date := doc.Find(".entry-meta").Find("a").Eq(0).Text()
	author := div.Find(".entry-meta").Find("a").Eq(1).Text()
	authorLink, _ := div.Find(".entry-meta").Find("a").Eq(1).Attr("href")
	*/
	var linksTitle []string
	var linksHref []string
	as := div.Find("a")
	for i := 0; i < as.Length(); i++ {
		linksTitle = append(linksTitle, as.Eq(i).Text())
		ss, _ := as.Eq(i).Attr("href")
		linksHref = append(linksHref, ss)
	}
	// test
	/*
	Println("meta: \t", meta)
	Println(h1, "\t", date, "\t", author, "\t", authorLink)
	for _, i := range links {
		Println(i)
	}
	Println("len(links)= ", len(links))
	Println("links[10] = ", links[11])
	*/	
	title := []string{"mataContent", "head"}
	content := []string{meta, h1}
	title = append(title, linksTitle...)
	content = append(content, linksHref...)
	wholeContext := [][]string{title, content}
	// wholeContext:={title,content,}  # error not figure out
	Println(title)
	Println(content)

	filename := "grapDate.csv"
	WriteCsv(filename, wholeContext)
}
func WriteCsv(filename string, line [][]string) {
	buf := new(bytes.Buffer)
	r2 := csv.NewWriter(buf)
	r2.WriteAll(line)
	r2.Flush()
	// Println(buf)		# for test
	fout, err := os.Create(filename)
	defer fout.Close()
	if err != nil {
		Println(filename, err)
		return
	}
	fout.WriteString(buf.String())
	Println("ok")
}

