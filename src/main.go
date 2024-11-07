package main

import (
  "fmt"
  "github.com/gocolly/colly"
  "net/url"
)

var links = make(map[string]bool)
var linkQueue = make([]string, 0)

func enqueueLink(link string){
  links[link] = true
  linkQueue = append(linkQueue, link)
}

func dequeueLink() string {
  link := linkQueue[0]
  linkQueue = linkQueue[1:]
  return link
}


func BFS(c *colly.Collector, root string) {
  enqueueLink(root)
  for len(linkQueue) > 0 {
    link := dequeueLink()
    fmt.Println(link)
    c.Visit(link)
  }
}

func main() {
  c := colly.NewCollector()
  c.OnHTML("a[href]", func (e *colly.HTMLElement){
    link := e.Attr("href")
    link = e.Request.AbsoluteURL(link)
    if link == "" || link == "#" {
      return
    }
    u, err := url.Parse(link)
    if err != nil {
      return
    }
    if u.Scheme == "mailto" || u.Scheme == "tel" || u.Host == "" {
      return
    }

    link = fmt.Sprintf("%s://%s", u.Scheme, u.Host)
    if !links[link] {
      enqueueLink(link)
    }
  })
  
  var rootLink string
  fmt.Print("Enter the root link : ")
  fmt.Scanln(&rootLink)
  BFS(c, rootLink)
}
