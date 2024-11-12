package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

type post struct {
	Title       string `xml:"title"`
	Slug        string `xml:"-"`
	Description string `xml:"description"`
	Date        string `xml:"-"`
	RSSLink     string `xml:"link"`
	RSSPubDate  string `xml:"pubDate"`
}

func newPostsData() map[string]post {
	posts := map[string]post{}
	posts["behind-the-scenes"] = post{
		Title:       "Behind the Scenes: Building This Website",
		Slug:        "behind-the-scenes",
		Description: `I'm peeling back the curtain to give you a peek into the technologies and decisions that brought this site to life.`,
		Date:        "November, 2024",
		RSSPubDate:  "2024-11-12T08:00:00Z",
	}
	return posts
}

type RssChannel struct {
	Title         string    `xml:"title"`
	Description   string    `xml:"description"`
	Link          string    `xml:"link"`
	LastBuildDate time.Time `xml:"lastBuildDate"`
	Language      string    `xml:"language"`
	Items         []post    `xml:"items"`
}

func generateRSSFeed() {

	posts := newPostsData()
	items := []post{}
	for _, item := range posts {
		item.RSSLink = fmt.Sprintf("https://torkelaannestad.com/posts/%s", item.Slug)
		items = append(items, item)
	}

	feed := RssChannel{
		Title:         "TorkelAannestad.com developer blog",
		Description:   "My thought on programming, recent discoveries and learnings.",
		Link:          "https://torkelaannestad.com/posts",
		LastBuildDate: time.Now().UTC(),
		Language:      "en-US",
		Items:         items,
	}

	xmlData, err := xml.MarshalIndent(feed, "", "    ")
	if err != nil {
		return
	}

	// Add XML header to the feed
	rssFeed := []byte(xml.Header + string(xmlData))

	file, err := os.Create("ui/public/feed.xml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(rssFeed)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("RSS feed generated and saved to feed.xml.")
}

func main() {
	fmt.Println("hello worlds")
}
