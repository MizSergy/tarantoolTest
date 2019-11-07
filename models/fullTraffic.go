package models

import (
	"reflect"
	"time"
)

type FullTraffic struct {
	IVCode uint   `json:"ivcode" db:"ivcode,omitempty" msgpack:"ivcode"`
	VCode  string `json:"vcode" db:"vcode,omitempty" msgpack:"vcode"`

	CreateAt   time.Time `json:"create_at" db:"create_at,omitempty" msgpack:"-"`
	CreateDate time.Time `json:"create_date" db:"create_date,omitempty" msgpack:"-"`

	CreatedAt   int `msgpack:"create_at"`
	CreatedDate int `msgpack:"create_date"`


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

//func (t *FullTraffic)Compare(v FullTraffic){
//	if len(t.VCode) == 0 {
//		t.VCode = v.VCode
//	}
//	if !t.CreateAt.IsZero() && t.CreatedAt > 0{
//		t.CreateAt = time.Unix(int64(v.CreatedAt),0)
//	}
//
//	if !t.CreateDate.IsZero() && t.CreatedDate > 0{
//		t.CreateDate = time.Unix(int64(v.CreatedDate),0)
//	}
//
//	if t.SourceID == 0{
//		t.SourceID = v.SourceID
//	}
//	if t.Campaign == 0{
//		t.Campaign = v.Campaign
//	}
//	if t.StreamID == 0{
//		t.StreamID = v.StreamID
//	}
//	if t.AffiliateID == 0{
//		t.AffiliateID = v.AffiliateID
//	}
//	if t.PrelandID == 0{
//		t.PrelandID = v.PrelandID
//	}
//
//	if t.IsBreaked == 0{
//		t.IsBreaked = v.IsBreaked
//	}
//
//	if t.IsRefused == 0{
//		t.IsRefused = v.IsRefused
//	}
//	if t.IsUnique == 0{
//		t.IsUnique = v.IsUnique
//	}
//	if t.IsTest == 0{
//		t.IsTest = v.IsTest
//	}
//
//	if t.IsClick == 0{
//		t.IsClick = v.IsClick
//	}
//
//	if t.ProcessInterval == 0{
//		t.ProcessInterval = v.ProcessInterval
//	}
//
//	if t.ScreenWidth == 0{
//		t.ScreenWidth = v.ScreenWidth
//	}
//	if t.ScreenHeight == 0{
//		t.ScreenHeight = v.ScreenHeight
//	}
//
//	if len(t.Language) == 0{
//		t.Language = v.Language
//	}
//	if t.ClickPrice == 0{
//		t.ClickPrice = v.ClickPrice
//	}
//	if len(t.Browser) == 0{
//		t.Browser = v.Browser
//	}
//	if len(t.BrowserV) == 0{
//		t.BrowserV = v.BrowserV
//	}
//	if len(t.Os) == 0{
//		t.Os = v.Os
//	}
//	if len(t.OsV) == 0{
//		t.OsV = v.OsV
//	}
//	if len(t.Country) == 0{
//		t.Country = v.Country
//	}
//	if len(t.CountryCode) == 0{
//		t.CountryCode = v.CountryCode
//	}
//	if len(t.Region) == 0{
//		t.Region = v.Region
//	}
//	if len(t.City) == 0{
//		t.City = v.City
//	}
//	if t.Ip == 0{
//		t.Ip = v.Ip
//	}
//	if t.Device == 0{
//		t.Device = v.Device
//	}
//	if t.IsMobil == 0{
//		t.IsMobil = v.IsMobil
//	}
//	if len(t.Ad) == 0{
//		t.Ad = v.Ad
//	}
//	if len(t.Site) == 0{
//		t.Site = v.Site
//	}
//	if len(t.Sid1) == 0{
//		t.Sid1 = v.Sid1
//	}
//	if len(t.Sid2) == 0{
//		t.Sid2 = v.Sid2
//	}
//	if len(t.Sid3) == 0{
//		t.Sid3 = v.Sid3
//	}
//	if len(t.Sid4) == 0{
//		t.Sid4 = v.Sid4
//	}
//	if len(t.Sid5) == 0{
//		t.Sid5 = v.Sid5
//	}
//	if len(t.Sid6) == 0{
//		t.Sid6 = v.Sid6
//	}
//	if len(t.Sid7) == 0{
//		t.Sid7 = v.Sid7
//	}
//	if len(t.Sid8) == 0{
//		t.Sid8 = v.Sid8
//	}
//	if len(t.Sid9) == 0{
//		t.Sid9 = v.Sid9
//	}
//
//	if len(t.Sid10) == 0{
//		t.Sid10 = v.Sid10
//	}
//	if len(t.PrelandUrl) == 0{
//		t.PrelandUrl = v.PrelandUrl
//	}
//	if len(t.Session) == 0{
//		t.Session = v.Session
//	}
//	if t.UID == 0{
//		t.UID = v.UID
//	}
//}


func (t *FullTraffic)Compare(v FullTraffic){
	oval := reflect.ValueOf(t)
	nval := reflect.ValueOf(v)
	for i:=0; i < oval.NumField(); i++{
		name := oval.Type().Field(i).Name
		value := oval.Field(i).Interface()
		if value == nil {
			value = nval.FieldByName(name)
		}
	}
}