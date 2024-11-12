package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"time"
)

type post struct {
	Title       string `xml:"title" json:"title"`
	Slug        string `xml:"-" json:"slug"`
	Description string `xml:"description" json:"description"`
	Date        string `xml:"-" json:"date"`
	RSSLink     string `xml:"link" json:"link"`
	RSSPubDate  string `xml:"pubDate" json:"pubDate"`
}

type RssChannel struct {
	Title         string    `xml:"title"`
	Description   string    `xml:"description"`
	Link          string    `xml:"link"`
	LastBuildDate time.Time `xml:"lastBuildDate"`
	Language      string    `xml:"language"`
	Items         []post    `xml:"items"`
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

func generateXML(fileContent []byte) ([]byte, error) {
	xmlData, err := xml.MarshalIndent(fileContent, "", "    ")
	if err != nil {
		return nil, err
	}

	rssFeed := []byte(xml.Header + string(xmlData))

	fmt.Println("RSS feed generated and saved to feed.xml.")

	return rssFeed, nil
}

func app() error {
	filePath := "internal/data/test-data.json"
	fileContent, err := readInputFile(filePath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	xmlData, err := generateXML(fileContent)
	if err != nil {
		return err
	}
	err = writeFeedFile(xmlData)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := app()
	if err != nil {
		log.Fatal("Something went wrong", err)
	}
}
