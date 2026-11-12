package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"sara": {
		AuthToken: "987XYZ",
		Username:  "sara",
	},
	"omar": {
		AuthToken: "654UVW",
		Username:  "omar",
	},
	"leila": {
		AuthToken: "321RST",
		Username:  "leila",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"sara": {
		Coins:    150,
		Username: "sara",
	},
	"omar": {
		Coins:    250,
		Username: "omar",
	},
	"leila": {
		Coins:    350,
		Username: "leila",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
