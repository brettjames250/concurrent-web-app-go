# concurrent-web-app-go

A simple Go application that retrieves and displays the latest news stories from the [BBC News UK sitemap](https://www.bbc.co.uk/sitemaps/https-sitemap-uk-news-2.xml).

## Functionality

- Uses the net/http package in Go to retrieve news stories from the BBC News UK sitemap
- Stories are parsed and stored in a struct
- Data from the struct is passed to a html/template, which is rendered on the webpage

## Dependencies

- encoding/xml for XML encoding and decoding
- html/template for rendering the HTML template
- io/ioutil for reading response body into a byte array
- os for checking for errors and exiting the app if necessary

## Running the app

```
go run main.go
```

Navigate to `http://localhost:8080/news` to view the latest news stories.

## Built With

- [Go](https://golang.org/)
- [BBC News UK Sitemap](https://www.bbc.co.uk/sitemaps/https-sitemap-uk-news-2.xml)



