package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"raiden": {
		AuthToken: "123RAI",
		Username:  "raiden",
	},
	"goro": {
		AuthToken: "123GOR",
		Username:  "goro",
	},
	"cassie": {
		AuthToken: "123CAS",
		Username:  "cassie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"raiden": {
		Coins:    200,
		Username: "raiden",
	},
	"goro": {
		Coins:    250,
		Username: "goro",
	},
	"cassie": {
		Coins:    400,
		Username: "cassie",
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
