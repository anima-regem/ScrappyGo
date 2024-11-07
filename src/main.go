package main

import (
  "fmt"
  "github.com/gocolly/colly"
  "net/url"
  
)

var links = make(map[string]bool)
var linksArray = make([]string, 0)

func pushLink(link string){
  linksArray = append(linksArray, link)
  links[link] = true
}

func popLink() string {
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
  if len(linksArray) > 0 {
    link := popLink()
    if !containsLink(link){
      BFS(c, link)
    }
  }
}

func main(){
  c := colly.NewCollector()

  c.OnHTML("a[href]", func(e *colly.HTMLElement){
    link := e.Attr("href")
    url, _ := url.Parse(link)
    link = url.Host
    if link == "" {
      return
    }
    if containsLink(link){
      return
    }
    fmt.Println(link)
    pushLink(link)
    fmt.Println(linksArray)
  })

  c.OnRequest(func(r *colly.Request){
    fmt.Println("Visiting", r.URL)
  })

  BFS(c, "https://go.dev")
}
