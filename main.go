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

type SitemapPage struct {
	NewStories []NewsStory `xml:"url"`
}

type NewsStory struct {
	Location string `xml:"loc"`
	Date string `xml:"news>publication_date"`
	Publication string `xml:"news>publication>name"`
	Title    string `xml:"news>title"`
}

func getNewsStories() SitemapPage {
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
	var siteMapPage SitemapPage
	xml.Unmarshal(bytes, &siteMapPage)

	return siteMapPage
}

func newsHandler(responseWriter http.ResponseWriter, request *http.Request) {

	siteMapPage := getNewsStories()
	template, _ := template.ParseFiles("news.html")
	err := template.Execute(responseWriter, siteMapPage)

	if err != nil {
		fmt.Print("Error parsing template")
	}
}

func main() {
	http.HandleFunc("/news", newsHandler)
	http.ListenAndServe(":8080", nil)
}
