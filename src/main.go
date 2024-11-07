package main

import (
  "fmt"
  "github.com/gocolly/colly"
  "net/url"
  
)

var links = make(map[string]bool)
var linksArray = make([]string, 0)

func pushLink(link string){
  fmt.Println("Pushed ", link)
  if containsLink(link){
    return
  }
  linksArray = append(linksArray, link)
  links[link] = true
}

func popLink() string {
  fmt.Println("Popped ", linksArray[len(linksArray)-1])
  length := len(linksArray)
  if length == 0 {
    return ""
  } else {
    link := linksArray[length -1]
    linksArray = linksArray[:length -1]
    return link
  }
}

func containsLink(link string) bool {
  _, ok := links[link]
  return ok
}

func BFS(c *colly.Collector, root string){
  c.Visit(root)
  for len(linksArray) > 0 {
    link := popLink()
    fmt.Println(link)
    BFS(c, link)
  }
}

func main(){
  c := colly.NewCollector()

  c.OnHTML("a[href]", func(e *colly.HTMLElement){
    link := e.Attr("href")
    		parsedURL, err := url.Parse(link)
		if err != nil || parsedURL.Host == "" {
			return
		}

		link = e.Request.AbsoluteURL(link)

    parsedURL, _ = url.Parse(link)
    link = parsedURL.Host
    pushLink(link)

  })

  c.OnRequest(func(r *colly.Request){
    fmt.Println("Visiting", r.URL)
  })

  BFS(c, "https://go.dev")
}
