package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"time"
)

type Post struct {
	Title       string `xml:"title" json:"title"`
	Slug        string `xml:"-" json:"slug"`
	Description string `xml:"description" json:"description"`
	Date        string `xml:"-" json:"date"`
	RSSLink     string `xml:"link" json:"link"`
	RSSPubDate  string `xml:"pubDate" json:"pubDate"`
}

type RssChannel struct {
	Title         string    `xml:"title" json:"title"`
	Description   string    `xml:"description" json:"description"`
	Link          string    `xml:"link" json:"link"`
	LastBuildDate time.Time `xml:"lastBuildDate" json:"-"`
	Language      string    `xml:"language" json:"language"`
	Items         []Post    `xml:"items" json:"items"`
}

func readInputFile(filePath string) ([]byte, error) {
	body, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return body, nil
}
func writeFeedFile(rssFeed []byte) error {
	file, err := os.Create("feed.xml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	_, err = file.Write(rssFeed)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return nil
}

func generateXMLFromJSON(jsonData []byte) ([]byte, error) {
	rssFeed := RssChannel{}
	err := json.Unmarshal(jsonData, &rssFeed)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v\n", rssFeed.Items[0].Title)
	xmlData, err := xml.MarshalIndent(rssFeed, "", "    ")
	if err != nil {
		return nil, err
	}

	xmlFeed := []byte(xml.Header + string(xmlData))

	return xmlFeed, nil
}

func app() error {

	filePath := "internal/data/test-data.json"
	fileContent, err := readInputFile(filePath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	xmlData, err := generateXMLFromJSON(fileContent)
	if err != nil {
		return err
	}

	err = writeFeedFile(xmlData)
	if err != nil {
		return err
	}
	fmt.Println("RSS feed generated and saved to feed.xml.")
	return nil
}

func main() {
	err := app()
	if err != nil {
		log.Fatal("Something went wrong", err)
	}
}
