package errors

import "errors"

var (
	ErrNoKey                   = errors.New("empty key cannot be set")
	ErrNoData                  = errors.New("empty data cannot be set")
	ErrKeyNotFound             = errors.New("key not found")
	ErrEmptyKey                = errors.New("key with no data")
	ErrEmptyUnmarshallFunction = errors.New("please provide a valid unmarshall function")
	ErrUnmarshall              = errors.New("failed to unmarshall")
	ErrDbConnection            = errors.New("failed to connect mysql Db with given credential")
)
