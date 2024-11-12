package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
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
func writeFeedFile(rssFeed []byte, outputFileName, outputFilePath string) error {
	filepath := outputFileName
	if outputFilePath != "" {
		filepath = fmt.Sprintf("%s/%s", outputFilePath, outputFileName)
	}
	file, err := os.Create(filepath)
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
	xmlData, err := xml.MarshalIndent(rssFeed, "", "    ")
	if err != nil {
		return nil, err
	}

	xmlFeed := []byte(xml.Header + string(xmlData))

	return xmlFeed, nil
}

func validateInput(inputFilePath string, outputFilePath string) error {
	//inputFilePath
	if inputFilePath == "" {
		fmt.Println("Please provide a filepath for input file. Example: somedirectory/data.json")
		return errors.New("error")
	}
	_, err := os.Stat(inputFilePath)
	if os.IsNotExist(err) {
		fmt.Println("path does not exist")
		return fmt.Errorf("path does not exist")
	} else if err != nil {
		fmt.Printf("unable to access the path: %v\n", inputFilePath)
		return err
	}

	//outputFilePath
	if outputFilePath == "" {
		return nil
	}
	info, err := os.Stat(outputFilePath)
	if err != nil {
		fmt.Printf("unable to access the path: %v", err)
		return err

	}

	if !info.IsDir() {
		fmt.Println("Path is a not a valid directory.")
		return errors.New("error")
	}

	return nil
}

func app() error {
	inputFilePath := flag.String("input-filepath", "", "example: somedirectory/data.json")
	outputFilePath := flag.String("output-filepath", "", "current working directory")
	outputFileName := flag.String("output-filename", "feed.xml", "feed.xml")

	flag.Parse()
	err := validateInput(*inputFilePath, *outputFilePath)
	if err != nil {
		os.Exit(0)
	}

	fileContent, err := readInputFile(*inputFilePath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	xmlData, err := generateXMLFromJSON(fileContent)
	if err != nil {
		return err
	}

	err = writeFeedFile(xmlData, *outputFileName, *outputFilePath)
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
