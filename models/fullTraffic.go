package models

type FullTraffic struct {
	IVCode uint   `json:"ivcode" db:"ivcode,omitempty" msgpack:"ivcode"`
	VCode  string `json:"vcode" db:"vcode,omitempty" msgpack:"vcode"`

	CreateAt   int `json:"create_at" db:"create_at,omitempty" msgpack:"create_at"`
	CreateDate int `json:"create_date" db:"create_date,omitempty" msgpack:"create_date"`

	SourceID    int `json:"source_id" db:"source_id,omitempty" msgpack:"source_id"`
	Campaign    int `json:"campaign" db:"campaign,omitempty" msgpack:"campaign"`
	StreamID    int `json:"stream_id" db:"stream_id,omitempty" msgpack:"stream_id"`
	AffiliateID int `json:"affiliate_id" db:"affiliate_id,omitempty" msgpack:"affiliate_id"`
	PrelandID   int `json:"preland_id" db:"preland_id,omitempty" msgpack:"preland_id"`

	IsBreaked int `json:"is_breaked" db:"is_breaked" msgpack:"is_breaked"`
	IsRefused int `json:"is_refused" db:"is_refused" msgpack:"is_refused"`
	IsUnique  int `json:"is_unique" db:"is_unique" msgpack:"is_unique"`
	IsTest    int `json:"is_test" db:"is_test" msgpack:"is_test"`
	IsClick   int `json:"is_click" db:"is_click" msgpack:"is_click"`

	ProcessInterval float64 `json:"process_interval" db:"process_interval" msgpack:"process_interval"`
	ScreenWidth     int     `json:"screen_width" db:"screen_width" msgpack:"screen_width"`
	ScreenHeight    int     `json:"screen_height" db:"screen_height" msgpack:"screen_height"`

	Language    string  `json:"language" db:"language" msgpack:"language"`
	ClickPrice  float64 `json:"click_price" db:"click_price" msgpack:"click_price"`
	Browser     string  `json:"browser" db:"browser" msgpack:"browser"`
	BrowserV    string  `json:"browserv" db:"browserv" msgpack:"browserv"`
	Os          string  `json:"os" db:"os" msgpack:"os"`
	OsV         string  `json:"osv" db:"osv" msgpack:"osv"`
	Country     string  `json:"country" db:"country" msgpack:"country"`
	CountryCode string  `json:"country_code" db:"country_code" msgpack:"country_code"`
	Region      string  `json:"region" db:"region" msgpack:"region"`
	City        string  `json:"city" db:"city" msgpack:"city"`
	Ip          uint    `json:"ip" db:"ip" msgpack:"ip"`
	Device      int     `json:"device" db:"device" msgpack:"device"`
	IsMobil     int     `json:"is_mobil" db:"is_mobil" msgpack:"is_mobil"`
	Ad          string  `json:"ad" db:"ad" msgpack:"ad"`
	Site        string  `json:"site" db:"site" msgpack:"site"`

	Sid1       string `json:"sid1" db:"sid1" msgpack:"sid1"`
	Sid2       string `json:"sid2" db:"sid2" msgpack:"sid2"`
	Sid3       string `json:"sid3" db:"sid3" msgpack:"sid3"`
	Sid4       string `json:"sid4" db:"sid4" msgpack:"sid4"`
	Sid5       string `json:"sid5" db:"sid5" msgpack:"sid5"`
	Sid6       string `json:"sid6" db:"sid6" msgpack:"sid6"`
	Sid7       string `json:"sid7" db:"sid7" msgpack:"sid7"`
	Sid8       string `json:"sid8" db:"sid8" msgpack:"sid8"`
	Sid9       string `json:"sid9" db:"sid9" msgpack:"sid9"`
	Sid10      string `json:"sid10" db:"sid10" msgpack:"sid10"`
	PrelandUrl string `json:"preland_url" db:"preland_url" msgpack:"preland_url"`
	Session    string `json:"session" db:"session_id" msgpack:"session_id"`
	UID        int    `json:"uid" db:"uid" msgpack:"uid"`
}
