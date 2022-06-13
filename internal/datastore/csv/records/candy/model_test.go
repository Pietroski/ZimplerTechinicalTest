package candy_store

import "testing"

func TestByCandyAmount_sortByCandyAmount(t *testing.T) {
	tests := []struct {
		name string
		ca   ByCandyAmount
	}{
		{
			name: "happy sort",
			ca: ByCandyAmount([]*ConsumeOverview{
				{
					Name:           "Augusto",
					FavouriteSnack: "Hagen-Das",
					TotalSnacks:    5,
				},
				{
					Name:           "Augusto",
					FavouriteSnack: "Kopenhagen",
					TotalSnacks:    7,
				},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ca.sortByCandyAmount()
		})
	}
}
