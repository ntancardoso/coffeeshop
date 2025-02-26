package coffee

import (
	"testing"
)

func init() {
	Coffees = CoffeeList{
		List: []CoffeeDetails{
			CoffeeDetails{"Latte", 2.5},
			CoffeeDetails{"Flat white", 2},
			CoffeeDetails{"Capuccino", 2.25},
		},
	}
}

func TestIsCoffeeAvailable(t *testing.T) {
	type testCase struct {
		coffeeType string
		want       bool
	}

	cases := []testCase{
		{"lat", false},
		{"Latte", true},
		{"", false},
		{"cappacunio", false},
	}

	for _, tc := range cases {
		got := IsCoffeeAvailable(tc.coffeeType)
		if tc.want != got {
			t.Errorf("Expected '%v', but got '%v'", tc.want, got)
		}
	}
}
