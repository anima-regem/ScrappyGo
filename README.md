
# 🕸️ Web Crawler 3000 — The Link Gobbler

![crawler gif](https://media.giphy.com/media/3oz8xK7FR9kQjwhmxi/giphy.gif)

## 🚀 About This Project
**Web Crawler 3000** is your very own link-hungry web-surfing bot, powered by Go and Colly. It'll keep gobbling up links until there are none left on the internet… or until it realizes it can't quite visit *everything*.

Simply point it to a webpage, and watch it go! It'll scrape, crawl, queue, dequeue, and consume links like a toddler with spaghetti. But don’t worry — it won't be ringing anyone’s doorbell (no mailto or tel links!) or trying to get itself tangled in any weird infinite loops. 

## 💼 Features
- **Queue-based BFS**: None of that pesky recursion here! Our bot's got its priorities straight and only gobbles up links it hasn't tasted before.
- **Absolute URL Master**: Ever tried figuring out whether `../about` or `#contact` will lead to hidden treasure? Let **Web Crawler 3000** handle the math.
- **Selective Tastes**: No "mailto" or "tel" links in this diet! We’re *web* crawlers, not phone crawlers.

## 🎬 Quickstart Guide
Fire up your terminal and run:
```bash
go run main.go
```

Then, just give it a starting URL when it politely asks:
```plaintext
Enter the root link: https://example.com
```

Grab a coffee while **Web Crawler 3000** works its way through every corner of the internet… okay, *not* every corner, but close enough. Links it finds will show up in your terminal, one after the other, so you know it’s not slacking off!

## 🤖 Behind the Scenes

Here's how it works:
1. **Enqueue and Dequeue**: We’ve got a queue, a map of visited links, and a dream.
2. **No Double-Dipping**: Every link is stored in our `links` map, so we don’t get déjà vu.
3. **Filtered Crawling**: Only valid HTTP/HTTPS links are allowed. No funny business.

## 😆 Fun Facts
- **Warning**: While this bot loves the internet, it might break up with you if you try pointing it at infinite loops. It’s BFS, not “Find Forever and Survive.”
- **Ideal Use Case**: Impress your friends by pretending you wrote a bot that could totally crawl *all of Google* if it wanted to.
- **Real Use Case**: Discovering that *many* websites out there are just slightly different versions of the same few links.

## ⚠️ Disclaimer
This bot isn’t responsible for any existential crises caused by repetitive URLs or missing `robots.txt` files. Run responsibly!

---

