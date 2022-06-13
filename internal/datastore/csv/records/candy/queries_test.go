package candy_store

import "testing"

func TestAnalyzeIncomingData(t *testing.T) {
	type args struct {
		row                     []string
		totalConsumptionMapping *map[string]*map[string]*PersonalConsumption
	}
	tests := []struct {
		name    string
		args    func() args
		wantErr bool
	}{
		{
			name: "happy path",
			args: func() args {
				name, candy, amnt, amount := "Augusto", "Kopenhagen", "7", 7
				a := args{
					row: []string{
						name, candy, amnt,
					},
					totalConsumptionMapping: &map[string]*map[string]*PersonalConsumption{
						name: {
							candy: {
								Name:   name,
								Candy:  candy,
								Amount: amount,
							},
						},
					},
				}

				return a
			},
			wantErr: false,
		},
		{
			name: "err not enough arguments",
			args: func() args {
				name, candy := "Augusto", "Kopenhagen"
				a := args{
					row: []string{
						name, candy,
					},
					totalConsumptionMapping: &map[string]*map[string]*PersonalConsumption{},
				}

				return a
			},
			wantErr: true,
		},
		{
			name: "failed to parse int amount",
			args: func() args {
				name, candy, amnt := "Augusto", "Kopenhagen", "----"
				a := args{
					row: []string{
						name, candy, amnt,
					},
					totalConsumptionMapping: &map[string]*map[string]*PersonalConsumption{},
				}

				return a
			},
			wantErr: true,
		},
		{
			name: "incomplete map - empty map",
			args: func() args {
				name, candy, amnt := "Augusto", "Kopenhagen", "7"
				a := args{
					row: []string{
						name, candy, amnt,
					},
					totalConsumptionMapping: &map[string]*map[string]*PersonalConsumption{},
				}

				return a
			},
			wantErr: false,
		},
		{
			name: "incomplete map - full until name",
			args: func() args {
				name, candy, amnt := "Augusto", "Kopenhagen", "7"
				a := args{
					row: []string{
						name, candy, amnt,
					},
					totalConsumptionMapping: &map[string]*map[string]*PersonalConsumption{
						name: {},
					},
				}

				return a
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AnalyzeIncomingData(tt.args().row, tt.args().totalConsumptionMapping); (err != nil) != tt.wantErr {
				t.Errorf("AnalyzeIncomingData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSnackOverview(t *testing.T) {
	type args struct {
		totalConsumptionMapping *map[string]*map[string]*PersonalConsumption
		overview                *map[string]*ConsumeOverview
	}
	tests := []struct {
		name string
		args func() args
	}{
		{
			name: "happy path",
			args: func() args {
				name, candy, amount := "Augusto", "Kopenhagen", 7
				a := args{
					totalConsumptionMapping: &map[string]*map[string]*PersonalConsumption{
						name: {
							candy: {
								Name:   name,
								Candy:  candy,
								Amount: amount,
							},
						},
					},
					overview: &map[string]*ConsumeOverview{
						name: {
							Name:           name,
							FavouriteSnack: candy,
							TotalSnacks:    amount,
						},
					},
				}

				return a
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SnackOverview(tt.args().totalConsumptionMapping, tt.args().overview)
		})
	}
}

func TestPopulateOverviewList(t *testing.T) {
	type args struct {
		overviewMapping *map[string]*ConsumeOverview
		overviewList    *[]*ConsumeOverview
	}
	tests := []struct {
		name string
		args func() args
	}{
		{
			name: "happy path",
			args: func() args {
				name, candy, amount := "Augusto", "Kopenhagen", 7
				a := args{
					overviewMapping: &map[string]*ConsumeOverview{
						name: {
							Name:           name,
							FavouriteSnack: candy,
							TotalSnacks:    amount,
						},
					},
					overviewList: &[]*ConsumeOverview{
						{
							Name:           name,
							FavouriteSnack: candy,
							TotalSnacks:    amount,
						},
					},
				}

				return a
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PopulateOverviewList(tt.args().overviewMapping, tt.args().overviewList)
		})
	}
}

func TestShowOffResults(t *testing.T) {
	type args struct {
		totalConsumptionMapping interface{}
		overview                interface{}
	}
	tests := []struct {
		name string
		args func() args
	}{
		{
			name: "happy path",
			args: func() args {
				name, candy, amount := "Augusto", "Kopenhagen", 7
				a := args{
					totalConsumptionMapping: &map[string]*map[string]*PersonalConsumption{
						name: {
							candy: {
								Name:   name,
								Candy:  candy,
								Amount: amount,
							},
						},
					},
					overview: &[]*ConsumeOverview{
						{
							Name:           name,
							FavouriteSnack: candy,
							TotalSnacks:    amount,
						},
					},
				}

				return a
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShowOffResults(tt.args().totalConsumptionMapping, tt.args().overview)
		})
	}
}
