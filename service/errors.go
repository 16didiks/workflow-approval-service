package service

import "errors"

var (
	ErrRequestNotFound         = errors.New("request not found")
	ErrRequestAlreadyProcessed = errors.New("request already approved or rejected")
)
