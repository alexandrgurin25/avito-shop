package entity

type UserInfo struct {
	Coins       int             `json:"coins"`
	Inventory   []InventoryItem `json:"inventory"`
	CoinHistory CoinHistory     `json:"coinHistory"`
}

type InventoryItem struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type ReceivedCoinTransaction struct {
	FromUser string `json:"fromUser"`
	Amount   int    `json:"amount"`
}

type SentCoinTransaction struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}

type CoinHistory struct {
	Received []SentCoinTransaction     `json:"received"`
	Sent     []ReceivedCoinTransaction `json:"sent"`
}
