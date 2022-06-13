package candy_store

import (
	error_handler "ZimplerTechinicalTest/internal/tools/handlers/errors"
	"encoding/json"
	"fmt"
	"strconv"
)

func rowExtractor(row []string) (name, candy string, amount int, err error) {
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
	error_handler.Handle(err)
	fmt.Println(string(bs))

	fmt.Println()
	fmt.Println("#############################################")
	fmt.Println()

	bs, err = json.MarshalIndent(overview, "", "  ")
	error_handler.Handle(err)
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

func SnackOverview(
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

		(*overview)[name] = newConsumeOverview(name, candyName, candyCount)
	}
}

func AnalyzeIncomingData(
	row []string,
	totalConsumptionMapping *map[string]*map[string]*PersonalConsumption,
) error {
	name, candy, amnt, err := rowExtractor(row)
	if err != nil {
		return fmt.Errorf("failed to extract row values: %v", err)
	}

	personAnalysis, ok := (*totalConsumptionMapping)[name]
	if !ok {
		(*totalConsumptionMapping)[name] = &map[string]*PersonalConsumption{
			candy: newPersonalConsumption(name, candy, amnt),
		}

		return nil
	}

	candyAnalysis, ok := (*personAnalysis)[candy]
	if !ok {
		(*personAnalysis)[candy] = newPersonalConsumption(name, candy, amnt)

		return nil
	}

	candyAnalysis.Amount += amnt

	return nil
}
