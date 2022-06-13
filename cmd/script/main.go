package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

const (
	csvRecordsPath = "cmd/script/records.csv"

	NameFieldNum = iota - 1
	CandyFieldNum
	EatenFieldNum
	FieldsCount
)

type (
	ConsumeOverview struct {
		Name           string
		FavouriteSnack string
		TotalSnacks    int
	}

	PersonalConsumption struct {
		Name   string
		Candy  string
		Amount int
	}

	ByCandyAmount []*ConsumeOverview
)

func (ca ByCandyAmount) Len() int           { return len(ca) }
func (ca ByCandyAmount) Less(i, j int) bool { return ca[i].TotalSnacks > ca[j].TotalSnacks }
func (ca ByCandyAmount) Swap(i, j int)      { ca[i], ca[j] = ca[j], ca[i] }

func NewConsumeOverview(name, candy string, count int) *ConsumeOverview {
	return &ConsumeOverview{
		Name:           name,
		FavouriteSnack: candy,
		TotalSnacks:    count,
	}
}

func NewPersonalConsumption(name, candy string, amount int) *PersonalConsumption {
	return &PersonalConsumption{
		Name:   name,
		Candy:  candy,
		Amount: amount,
	}
}

func main() {
	file, err := os.Open(csvRecordsPath)
	handleErr(err)
	defer func() {
		_ = file.Close()
	}()

	totalAnalysis := make(map[string]*map[string]*PersonalConsumption)
	overview := make(map[string]*ConsumeOverview)
	csvReader := csv.NewReader(file)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		handleErr(err)

		if err := AnalyzeIncomingData(rec, &totalAnalysis); err != nil {
			log.Printf("failed to run row analysis: %v", err)
		}
	}
	snackOverview(&totalAnalysis, &overview)

	overviewList := make([]*ConsumeOverview, len(overview))
	PopulateOverviewList(&overview, &overviewList)
	sort.Sort(ByCandyAmount(overviewList))

	ShowOffResults(&totalAnalysis, &overviewList)
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func AnalyzeIncomingData(
	row []string,
	totalConsumptionMapping *map[string]*map[string]*PersonalConsumption,
) error {
	name, candy, amnt, err := RowExtractor(row)
	if err != nil {
		return fmt.Errorf("failed to extract row values: %v", err)
	}

	personAnalysis, ok := (*totalConsumptionMapping)[name]
	if !ok {
		(*totalConsumptionMapping)[name] = &map[string]*PersonalConsumption{
			candy: NewPersonalConsumption(name, candy, amnt),
		}

		return nil
	}

	candyAnalysis, ok := (*personAnalysis)[candy]
	if !ok {
		(*personAnalysis)[candy] = NewPersonalConsumption(name, candy, amnt)

		return nil
	}

	candyAnalysis.Amount += amnt

	return nil
}

func snackOverview(
	totalConsumptionMapping *map[string]*map[string]*PersonalConsumption,
	overview *map[string]*ConsumeOverview,
) {
	for name, candyMap := range *totalConsumptionMapping {
		var candyName string
		var candyCount, highestCandyCount int

		for candy, stats := range *candyMap {
			if stats.Amount > highestCandyCount {
				highestCandyCount = stats.Amount
				candyName = candy
			}

			candyCount += stats.Amount
		}

		(*overview)[name] = NewConsumeOverview(name, candyName, candyCount)
	}
}

func RowExtractor(row []string) (name, candy string, amount int, err error) {
	rowLen := len(row)
	if rowLen != FieldsCount {
		return "", "", 0, fmt.Errorf("invalid row length: %v", rowLen)
	}

	name = row[NameFieldNum]
	candy = row[CandyFieldNum]
	strAmnt := row[EatenFieldNum]
	amnt, err := strconv.Atoi(strAmnt)
	if err != nil {
		return "", "", 0, fmt.Errorf("failed to parse int string: %v", err)
	}

	amount = int(amnt)

	return
}

func ShowOffResults(
	totalConsumptionMapping interface{},
	overview interface{},
) {
	bs, err := json.MarshalIndent(totalConsumptionMapping, "", "  ")
	handleErr(err)
	fmt.Println(string(bs))

	fmt.Println()
	fmt.Println("#############################################")
	fmt.Println()

	bs, err = json.MarshalIndent(overview, "", "  ")
	handleErr(err)
	fmt.Println(string(bs))
}

func PopulateOverviewList(
	overviewMapping *map[string]*ConsumeOverview,
	overviewList *[]*ConsumeOverview,
) {
	var idx int
	for _, overview := range *overviewMapping {
		(*overviewList)[idx] = overview
		idx++
	}
}
