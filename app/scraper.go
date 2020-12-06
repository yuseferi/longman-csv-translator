package app

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
)

// Scrape a word from translator website
func (app *Application) ScrapeWord(word string) (string, error) {
	var wordTranslate = ""
	var WordUrl = fmt.Sprintf("%s/%s", app.Config.BaseUrl, word)
	res, err := http.Get(WordUrl)
	if err != nil || res.StatusCode != 200 {
		app.Logger.Error("err to get the word", zap.Error(err), zap.Any("word", word), zap.Any("status_code", res.StatusCode))
		return "", err
	}

	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		app.Logger.Error("err on load html", zap.Error(err), zap.Any("word", word))
		return "", err
	}

	doc.Find(".entry_content").Each(func(i int, s *goquery.Selection) {
		wordTranslate = s.Text()

	})
	return wordTranslate, nil
}

func (app *Application) ScrapeAll() error {
	inputFile, err := os.Open(app.Config.CSVWordInputFile)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	reader := csv.NewReader(inputFile)
	reader.Comma = '|'

	// File writer
	outputFile, err := os.Create(app.Config.CSVWordOutputFile)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()
	// Process CSV file line by line
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			app.Logger.Error("err on load a word from csv", zap.Error(err), zap.Any("word", record[0]))
			continue
		}
		app.Logger.Info("start to translate", zap.Any("word", record[0]))
		translatedWorld, err := app.ScrapeWord(record[0])
		if err != nil {
			app.Logger.Error("error on get word from translator website", zap.Error(err), zap.Any("word", record[0]))
			continue
		}
		err = writer.Write([]string{record[0], translatedWorld})
		if err != nil {
			return err
		}
	}
	return nil
}
