package models

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"strconv"
	"time"
)

type Breaking struct {
	IVCode   uint   `json:"ivcode" db:"ivcode" msgpack:"ivcode"`
	VCode    string `json:"vcode" db:"vcode" msgpack:"vcode"`
	CreateAt int    `msgpack:"create_at"`

	StreamID        int     `json:"stream_id" db:"stream_id" msgpack:"stream_id"`
	AffiliateID     int     `json:"affiliate_id" db:"affiliate_id" msgpack:"affiliate_id"`
	IsBreaked       int     `json:"is_breaked" db:"is_breaked" msgpack:"is_breaked"`
	IsRefused       int     `json:"is_refused" db:"is_refused" msgpack:"is_refused"`
	ProcessInterval float64 `json:"process_interval" db:"process_interval" msgpack:"process_interval"`

	ScreenHeight int    `json:"screen_height" db:"screen_height" msgpack:"screen_height"`
	ScreenWidth  int    `json:"screen_width" db:"screen_width" msgpack:"screen_width"`
	Language     string `json:"language" db:"language" msgpack:"language"`
	CreatedAt    time.Time `json:"create_at" db:"create_at"  msgpack:"-"`
}

func (b Breaking) BreaksToTraffic() FullTraffic {
	refTime, err := strconv.Atoi(os.Getenv("MIN_REFUSE_TIME"))
	if err != nil || refTime == 0 {
		refTime = 14
	}

	if b.ProcessInterval != 0 && b.ProcessInterval < float64(refTime) {
		b.IsRefused = 1
	}

	uid, err := uuid.Parse(b.VCode)
	if err != nil {
		fmt.Println(err.Error())
	}

	b.IVCode = uint(uid.ID())

	return FullTraffic{
		VCode:           b.VCode,
		IVCode:          b.IVCode,
		CreateDate:      b.CreateAt,
		IsBreaked:       b.IsBreaked,
		IsClick:         0,
		IsRefused:       b.IsRefused,
		StreamID:        b.StreamID,
		AffiliateID:     b.AffiliateID,
		ProcessInterval: b.ProcessInterval,
		ScreenWidth:     b.ScreenWidth,
		ScreenHeight:    b.ScreenHeight,
		Language:        b.Language,
	}
}
