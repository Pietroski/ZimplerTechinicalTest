package main

import (
	candy_store "ZimplerTechinicalTest/internal/datastore/csv/records/candy"
	error_handler "ZimplerTechinicalTest/internal/tools/handlers/errors"
	"encoding/csv"
	"io"
	"log"
	"os"
	"sort"
)

const csvRecordsPath = "internal/datastore/csv/records/candy/records.csv"

func main() {
	file, err := os.Open(csvRecordsPath)
	error_handler.Handle(err)
	defer func() {
		_ = file.Close()
	}()

	totalAnalysis := make(map[string]*map[string]*candy_store.PersonalConsumption)
	overview := make(map[string]*candy_store.ConsumeOverview)
	csvReader := csv.NewReader(file)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		error_handler.Handle(err)

		if err := candy_store.AnalyzeIncomingData(rec, &totalAnalysis); err != nil {
			log.Printf("failed to run row analysis: %v", err)
		}
	}
	candy_store.SnackOverview(&totalAnalysis, &overview)

	overviewList := make([]*candy_store.ConsumeOverview, len(overview))
	candy_store.PopulateOverviewList(&overview, &overviewList)
	sort.Sort(candy_store.ByCandyAmount(overviewList))

	candy_store.ShowOffResults(&totalAnalysis, &overviewList)
}
