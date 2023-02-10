package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

const bbcSiteMapLink = "https://www.bbc.co.uk/sitemaps/https-sitemap-uk-news-2.xml"

type SitemapIndex struct {
	// <loc> tag under each <url>
	NewStories []NewsStory `xml:"url"`
}

type NewsStory struct {
	Location string `xml:"loc"`
	Date string `xml:"news>publication_date"`
	Publication string `xml:"news>publication>name"`
	Title    string `xml:"news>title"`
}

type NewsPage struct {
	News SitemapIndex
}

func getNewsStories() SitemapIndex {
	// retrieving sitemap from BBC News UK
	resp, err := http.Get(bbcSiteMapLink)

	// checking for error
	if err != nil {
		fmt.Print("Error connecting to BBC sitemaps")
		os.Exit(1)
	}

	defer resp.Body.Close()

	// getting byte arru from response body
	bytes, err := ioutil.ReadAll(resp.Body)

	// checking for error
	if err != nil {
		fmt.Print("Error reading response body")
		os.Exit(1)
	}

	// Un-marshalling byte array into SitemapIndex struct
	var siteMap SitemapIndex
	xml.Unmarshal(bytes, &siteMap)

	return siteMap
}

func newsHandler(responseWriter http.ResponseWriter, request *http.Request) {

	siteMap := getNewsStories()
	newsPage := NewsPage{News: siteMap}
	template, _ := template.ParseFiles("news.html")
	err := template.Execute(responseWriter, newsPage)

	if err != nil {
		fmt.Print("Error parsing template")
	}
}

func main() {
	http.HandleFunc("/news", newsHandler)
	http.ListenAndServe(":8080", nil)
}
