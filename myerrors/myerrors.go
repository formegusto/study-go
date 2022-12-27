package myerrors

import "errors"

var NoMoney = errors.New("Can't withdraw you are poor")

var NotFound = errors.New("Not Found")
var WordExists = errors.New("That word already exists")
var CantUpdate = errors.New("Cant update non-existing word")
var CantDelete = errors.New("Cant delete non-existing word")