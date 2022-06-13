package candy_store

type (
	Querier interface {
		AnalyzeIncomingData(
			row []string,
			totalConsumptionMapping *map[string]*map[string]*PersonalConsumption,
		) error

		PopulateOverviewList(
			overviewMapping *map[string]*ConsumeOverview,
			overviewList *[]*ConsumeOverview,
		)

		SnackOverview(
			totalConsumptionMapping *map[string]*map[string]*PersonalConsumption,
			overview *map[string]*ConsumeOverview,
		)

		ShowOffResults(
			totalConsumptionMapping interface{},
			overview interface{},
		)
	}
)
