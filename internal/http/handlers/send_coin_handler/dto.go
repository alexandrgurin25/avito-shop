package send_coin_handler

type SendCoin struct {
	Username string `json:"toUser"`
	Amount   int    `json:"amount"`
}
