package models

import (
	"time"
)

type Click struct {
	IVCode      uint    `json:"ivcode" db:"ivcode" msgpack:"ivcode"`
	VCode       string  `json:"vcode" db:"vcode" msgpack:"vcode"`
	CreateAt    int     `msgpack:"create_at"`
	IsUnique    int     `json:"is_unique" db:"is_unique" msgpack:"is_unique"`
	IsMobil     int     `json:"is_mobil" db:"is_mobil" msgpack:"is_mobil"`
	IsTest      int     `json:"is_test" db:"is_test" msgpack:"is_test"`
	SourceID    int     `json:"source_id" db:"source_id" msgpack:"source_id"`
	Campaign    int     `json:"campaign" db:"campaign" msgpack:"campaign"`
	StreamId    int     `json:"stream_id" db:"stream_id" msgpack:"stream_id"`
	AffiliateId int     `json:"affiliate_id" db:"affiliate_id" msgpack:"affiliate_id"`
	PrelandID   int     `json:"preland_id" db:"preland_id" msgpack:"preland_id"`
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
	Ad          string  `json:"ad" db:"ad" msgpack:"ad"`
	Site        string  `json:"site" db:"site" msgpack:"site"`
	Sid1        string  `json:"sid1" db:"sid1" msgpack:"sid1"`
	Sid2        string  `json:"sid2" db:"sid2" msgpack:"sid2"`
	Sid3        string  `json:"sid3" db:"sid3" msgpack:"sid3"`
	Sid4        string  `json:"sid4" db:"sid4" msgpack:"sid4"`
	Sid5        string  `json:"sid5" db:"sid5" msgpack:"sid5"`
	Sid6        string  `json:"sid6" db:"sid6" msgpack:"sid6"`
	Sid7        string  `json:"sid7" db:"sid7" msgpack:"sid7"`
	Sid8        string  `json:"sid8" db:"sid8" msgpack:"sid8"`
	Sid9        string  `json:"sid9" db:"sid9" msgpack:"sid9"`
	Sid10       string  `json:"sid10" db:"sid10" msgpack:"sid10"`
	PrelandUrl  string  `json:"preland_url" db:"preland_url" msgpack:"preland_url"`
	Session     string  `json:"session" db:"session_id" msgpack:"session_id"`
	UserID      int     `json:"uid" db:"uid" msgpack:"uid"`

	CreatedAt time.Time `json:"create_at" db:"create_at" msgpack:"-"`
}

func (c Click) ClickToTraffic() FullTraffic {


	return FullTraffic{
		VCode:    c.VCode,
		IVCode:   c.IVCode,
		CreateAt: c.CreateAt,

		IsTest:   c.IsTest,
		IsUnique: c.IsUnique,
		IsMobil:  c.IsMobil,
		IsClick:  1,

		Campaign:    c.Campaign,
		SourceID:    c.SourceID,
		AffiliateID: c.AffiliateId,
		StreamID:    c.StreamId,
		PrelandUrl:  c.PrelandUrl,
		PrelandID:   c.PrelandID,

		Device:      c.Device,
		ClickPrice:  c.ClickPrice,
		Browser:     c.Browser,
		BrowserV:    c.BrowserV,
		Os:          c.Os,
		OsV:         c.OsV,
		Country:     c.Country,
		CountryCode: c.CountryCode,
		Region:      c.Region,
		City:        c.City,
		Ip:          c.Ip,

		Ad:   c.Ad,
		Site: c.Site,

		Sid1:    c.Sid1,
		Sid2:    c.Sid2,
		Sid3:    c.Sid3,
		Sid4:    c.Sid4,
		Sid5:    c.Sid5,
		Sid6:    c.Sid6,
		Sid7:    c.Sid7,
		Sid8:    c.Sid8,
		Sid9:    c.Sid9,
		Sid10:   c.Sid10,
		Session: c.Session,

		UID: c.UserID,
	}
}
