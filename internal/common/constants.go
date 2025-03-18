package common

import (
	"time"
)

const ExpirationTime = 2 * time.Hour // 2 ч живет jwt
const StartBalance = 1000            // Баланс пользователя

// Мерч - цена
var ItemMap = map[string]int{
	"t-shirt":    80,
	"cup":        20,
	"book":       50,
	"pen":        10,
	"powerbank":  200,
	"hoody":      300,
	"umbrella":   200,
	"socks":      10,
	"wallet":     50,
	"pink-hoody": 500,
}
