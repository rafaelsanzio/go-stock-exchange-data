package data

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/rafaelsanzio/go-stock-exchange-data/model"
)

func ReadJSONData() []model.Stock {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting directory: %v", err)
	}

	err = os.Chdir(fmt.Sprintf("%s/data", rootDir))
	if err != nil {
		log.Fatalf("Error changing directory: %v", err)
	}

	file, err := os.Open("data.json")
	if err != nil {
		log.Fatalf("Error open JSON file: %v", err)
	}
	defer file.Close()

	var data []model.Stock
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		log.Fatalf("Error decoding file: %v", err)
	}

	return data
}
