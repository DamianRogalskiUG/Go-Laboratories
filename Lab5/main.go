package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

// Structure to store data from the table
type TableData struct {
	Party    string
	Term2002 string
	Term2006 string
	Term2010 string
	Term2014 string
	Term2018 string
	Term2024 string
}

// Function to write data to a CSV file
func writeToCSV(data []TableData, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Party", "Term2002", "Term2006", "Term2010", "Term2014", "Term2018", "Term2024"} // Dodaj więcej nagłówków, jeśli tabela ma więcej kolumn
	writer.Write(headers)

	for _, record := range data {
		row := []string{record.Party, record.Term2002, record.Term2006, record.Term2010, record.Term2014, record.Term2018, record.Term2024} // Dodaj więcej pól, jeśli tabela ma więcej kolumn
		writer.Write(row)
	}

	return nil
}

// Function to scrape data from a table on a Wikipedia page
func scrapeTable(url string, tableIndex int) ([]TableData, error) {
	var tableData []TableData

	// Tworzenie nowego kolektora
	c := colly.NewCollector()
	currentTableIndex := 0

	// Definition of the HTML element to be scraped
	c.OnHTML(".wikitable > tbody", func(h *colly.HTMLElement) {
		currentTableIndex++
		// check if the current table index is the same as the one we are looking for
		if currentTableIndex == tableIndex {
			rowIndex := 0
			h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				// Skip the first row as it contains the headers
				if rowIndex == 0 {
					rowIndex++
					return
				}
				data := TableData{
					Party:    el.ChildText("td:nth-child(1)"),
					Term2002: el.ChildText("td:nth-child(2)"),
					Term2006: el.ChildText("td:nth-child(3)"),
					Term2010: el.ChildText("td:nth-child(4)"),
					Term2014: el.ChildText("td:nth-child(5)"),
					Term2018: el.ChildText("td:nth-child(6)"),
					Term2024: el.ChildText("td:nth-child(7)"),
				}
				tableData = append(tableData, data)
			})
		}
	})

	// Visit the URL
	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return tableData, nil
}

func main() {
	url := "https://pl.wikipedia.org/wiki/Gdynia"
	tableIndex := 2
	outputFile := "table_data.csv"

	// Scrape the table data
	tableData, err := scrapeTable(url, tableIndex)
	if err != nil {
		fmt.Println("Error scraping table:", err)
		return
	}

	// Write the data to a CSV file
	err = writeToCSV(tableData, outputFile)
	if err != nil {
		fmt.Println("Error writing to CSV:", err)
		return
	}

	fmt.Println("Data has been written to", outputFile)
}
