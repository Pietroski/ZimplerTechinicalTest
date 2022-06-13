package candy_store

import "sort"

const (
	NameFieldNum = iota
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

func newConsumeOverview(name, candy string, count int) *ConsumeOverview {
	return &ConsumeOverview{
		Name:           name,
		FavouriteSnack: candy,
		TotalSnacks:    count,
	}
}

func newPersonalConsumption(name, candy string, amount int) *PersonalConsumption {
	return &PersonalConsumption{
		Name:   name,
		Candy:  candy,
		Amount: amount,
	}
}

func (ca ByCandyAmount) sortByCandyAmount() {
	sort.Sort(ca)
}
