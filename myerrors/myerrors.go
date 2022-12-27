package myerrors

import "errors"

var NoMoney = errors.New("Can't withdraw you are poor")