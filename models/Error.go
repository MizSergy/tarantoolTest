package models

import "time"

type Errors struct {
	VCode           string		`json:"vcode" db:"vcode"`
	Description     string		`json:"description" db:"description"`
	CreateAt		time.Time	`json:"create_at" db:"create_at"`
}