package tools

import (
	"time"
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails {
	"alex": {
		AuthToken: "123ABC",
		Username: "Alex",
	},
	"piolo": {
		AuthToken: "456DEF",
		Username: "piolows",
	},
	"pierce": {
		AuthToken: "789GHI",
		Username: "pdpmiki",
	},
}

var mockCoinDetails = map[string]CoinDetails {
	"alex": {
		Coins: 500,
		Username: "Alex",
	},
	"piolo": {
		Coins: 300,
		Username: "piolows",
	},
	"pierce": {
		Coins: 100,
		Username: "pdpmiki",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	var clientData LoginDetails
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	var clientCoins CoinDetails
	clientCoins, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientCoins
}

func (d *mockDB) SetupDatabase() error {
	return nil
}