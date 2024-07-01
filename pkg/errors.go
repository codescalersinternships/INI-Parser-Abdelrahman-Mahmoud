package iniparser

import "errors"

var (
	errorReadingFile     = errors.New("error: trying to read file")
	errorWritingFile     = errors.New("error: trying to write file")
	errorWrongTypeOfFile = errors.New("error: wrong file extension")
	errorFileIsEmpty     = errors.New("error: ini file is empty or does not have section key value pair")
)
