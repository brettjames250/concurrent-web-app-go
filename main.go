package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type SitemapIndex struct {
	Urls []URL `xml:"url"`
}

type URL struct {
	Location string `xml:"loc"`
}

func main() {

	// retrieving sitemap from BBC News UK
	resp, err := http.Get("https://www.bbc.co.uk/sitemaps/https-sitemap-uk-news-2.xml")

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

	// iterating over URL slice + printing each news story location (link)
	for _, newsStory := range siteMap.Urls {
		fmt.Println(newsStory.Location)
	}

}
