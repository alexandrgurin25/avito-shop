package common

import (
	"errors"
	"time"
)

const ExpirationTime = 2 * time.Hour // 2 ч

var ErrIncorrectPassword = errors.New("incorrect password")
var ErrLowBalance = errors.New("insufficient coins")
var ErrUserNotFound = errors.New("user not found")

const StartBalance = 1000
