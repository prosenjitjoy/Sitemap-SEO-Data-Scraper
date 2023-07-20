package utils

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"main/model"
	"os"
	"strconv"
)

func OutputJSON(results []model.SEOdata) {
	jsonBytes, err := json.MarshalIndent(results, "", "    ")
	if err != nil {
		log.Fatal("Failed to marshal into json serial")
	}

	err = os.WriteFile("./out/seo_data.json", jsonBytes, 0644)
	if err != nil {
		log.Fatal("Failed to create json file:", err)
	}
}

func OutputCSV(results []model.SEOdata) {
	file, err := os.Create("./out/seo_data.csv")
	if err != nil {
		log.Fatal("Failed to create csv file:", err)
	}
	defer file.Close()

	csvFile := csv.NewWriter(file)
	csvFile.Comma = ';'
	csvFile.Write([]string{
		"Status Code",
		"URL",
		"Title",
		"Heading",
		"Description",
	})
	defer csvFile.Flush()

	for _, res := range results {
		csvFile.Write([]string{
			strconv.Itoa(res.StatusCode),
			res.URL,
			res.Title,
			res.Heading,
			res.Description,
		})
	}
}
