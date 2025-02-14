package common

import "errors"

var ErrIncorrectPassword = errors.New("incorrect password")
var ErrLowBalance = errors.New("insufficient coins")
var ErrUserNotFound = errors.New("user not found")
var ErrItemNotFound = errors.New("this merch name does not exist")
var ErrUserIdNotFoundContext = errors.New("userId not found in context")
