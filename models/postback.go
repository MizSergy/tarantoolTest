package models

import "time"

type PostBack struct {
	IVCode  uint   `json:"ivcode" db:"ivcode" msgpack:"ivcode"`
	OrderID string `json:"order_id" db:"order_id" msgpack:"order_id"`

	VCode       string `json:"vcode" db:"vcode" msgpack:"vcode"`
	CreatedAt   int    `json:"created_at" db:"created_at" msgpack:"created_at"`
	CreatedDate int    `json:"created_date" db:"created_date" msgpack:"created_date"`

	CreateAt   time.Time `json:"create_at" db:"create_at" msgpack:"-"`
	CreateDate time.Time `json:"create_date" db:"create_date" msgpack:"-"`

	Url    string `json:"url" db:"url" msgpack:"url"`
	Method string `json:"method" db:"method" msgpack:"method"`
	Params string `json:"params" db:"params" msgpack:"params"`

	StatusConfirmed int `json:"status_confirmed" db:"status_confirmed" msgpack:"status_confirmed"`
	StatusHold      int `json:"status_hold" db:"status_hold" msgpack:"status_hold"`
	StatusDeclined  int `json:"status_declined" db:"status_declined" msgpack:"status_declined"`
	StatusOther     int `json:"status_other" db:"status_other" msgpack:"status_other"`
	StatusPaid      int `json:"status_paid" db:"status_paid" msgpack:"status_paid"`

	Amount        float32 `json:"amount" db:"amount" msgpack:"amount"`
	Profit        float32 `json:"profit" db:"profit" msgpack:"profit"`
	PredictProfit float32 `json:"predict_profit" db:"predict_profit" msgpack:"predict_profit"`
	ResultMessage string  `json:"result_message" db:"result_message" msgpack:"result_message"`
}
