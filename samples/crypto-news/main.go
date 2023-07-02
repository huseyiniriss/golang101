package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

type Feed struct {
	Title string
	URL   string
}

var feedList = []Feed{
	{"Cointelegraph", "https://cointelegraph.com/rss"},
	{"CoinDesk", "https://www.coindesk.com/feed"},
	{"Bitcoin Magazine", "https://bitcoinmagazine.com/feed"},
	{"Bitcoin News", "https://news.bitcoin.com/feed"},
}

func main() {
	// sample rss feed reader app
	fmt.Println("--- CRYPTO NEWS ---")
	for {
		fmt.Println("1. View Feed List")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice: ")
		condition := 0
		fmt.Scanf("%d", &condition)
		switch condition {
		case 1:
			viewFeedList()
		case 2:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func viewFeedList() {
	for i, feed := range feedList {
		fmt.Printf("%d. %s\n", i+1, feed.Title)
	}
	fmt.Print("Enter your choice and read the news: ")
	condition := 0
	fmt.Scanf("%d", &condition)
	if condition > 0 && condition <= len(feedList) {
		showNews(feedList[condition-1])
	} else {
		fmt.Println("Invalid choice")
	}
}

func showNews(f Feed) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(f.URL)
	for _, item := range feed.Items {
		fmt.Println("Title:", item.Title)
		fmt.Println(item.Description)
		fmt.Println()
	}
}
