package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/kshvakov/clickhouse"
	"os"
	"strconv"
	"tarantoolTest/models"
	"time"
)

var ListClickhouseConnect []ClickhouseClient

//--------------------------------------------------Новое подключение кликхауса-----------------------------------------

type ClickhouseClient struct {
	sqlxclient *sqlx.DB
	Host       string
	Port       string
	Username   string
	Password   string
	DB         string
	Debug      string
}

func ClickhouseInit() {
	index := 1
	for {
		host := os.Getenv(fmt.Sprintf("CLICKHOUSE_HOST_%d", index))
		if len(host) == 0 {
			break
		}
		port := os.Getenv(fmt.Sprintf("CLICKHOUSE_PORT_%d", index))
		username := os.Getenv(fmt.Sprintf("CLICKHOUSE_USERNAME_%d", index))
		pass := os.Getenv(fmt.Sprintf("CLICKHOUSE_PASS_%d", index))
		db := os.Getenv(fmt.Sprintf("CLICKHOUSE_DB_%d", index))

		ListClickhouseConnect = append(ListClickhouseConnect, ClickhouseClient{
			Host:     host,
			Port:     port,
			Username: username,
			Password: pass,
			DB:       db,
			Debug:    os.Getenv("CLICKHOUSE_DEBUG"),
		})
		fmt.Printf("Connect clickhouse %s:%s\n", host, port)
		index ++
	}
	fmt.Printf("Clickhouse count connect %d\n", len(ListClickhouseConnect))

}

func (cl *ClickhouseClient) Disconnect() {
	cl.sqlxclient.Close()
}

func SqlxConnect() *sqlx.DB {
	var connection *sqlx.DB
	for _,v := range ListClickhouseConnect{
		if v.sqlxclient == nil{
			params := fmt.Sprintf(`%s:%s?username=%s&password=%s&database=%s&debug=%s`, v.Host, v.Port, v.Username, v.Password, v.DB, v.Debug)
			var err error
			v.sqlxclient, err = sqlx.Open("clickhouse", params)
			if err != nil {
				fmt.Printf("%s:%s - no connect\n", v.Host, v.Port)
				return nil
			}
			connection = v.sqlxclient
		}
		connection = v.sqlxclient
	}
	return connection
}


func writeBreak(items []models.Breaking) {
	time.Sleep(time.Second)
	query := `
					INSERT INTO tracker_db.breaking
						(vcode,
						ivcode,
						create_at,
					 	is_breaked,
						stream_id,						
						process_interval,
						affiliate_id,
						screen_height, 
						screen_width,
						language,
						is_refused)
					VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)`

	clickhouse := SqlxConnect()

	tx, err := clickhouse.Begin()
	if err != nil {
		panic(err.Error())
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		panic(err.Error())
	}

	for _, item := range items {
		refTime, err := strconv.Atoi(os.Getenv("MIN_REFUSE_TIME"))
		if err != nil || refTime == 0 {
			refTime = 14
		}
		item.IsRefused = 0
		if  item.ProcessInterval != 0  &&  item.ProcessInterval < float64(refTime) {
			item.IsRefused = 1
		}
		if _, err := stmt.Exec(
			item.VCode,
			item.IVCode,
			item.CreateAt,
			item.IsBreaked,
			item.StreamID,
			item.ProcessInterval,
			item.AffiliateID,
			item.ScreenHeight,
			item.ScreenWidth,
			item.Language,
			item.IsRefused,
		); err != nil {
			panic(err.Error())
		}

	}
	if err := tx.Commit(); err != nil {
		panic(err.Error())
	}

	stmt.Close()
	clickhouse.Close()
}

func writeClicks(items []models.Click) {
	query := `INSERT INTO tracker_db.click_logs
			(create_at,
			vcode,
			ivcode,
			is_unique,
			campaign,
			source_id, 
			click_price, 
			is_mobil,
			device,
			browser,
			os,
			country,
			region,
			city,
			ip,
			ad,
			site,
			sid1,
			sid2,
			sid3,
			sid4,
			sid5,
			sid6,
			sid7,
			sid8,
			sid9,
			sid10,
			preland_url,
			preland_id,
			session_id,
			is_test,
			country_code,
			osv,
			browserv,
			affiliate_id,
			stream_id,
			uid)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	clickhouse_conn := SqlxConnect()

	tx, err := clickhouse_conn.Begin()
	if err != nil{
		panic(err.Error())
	}

	stmt, err := tx.Prepare(query)
	if err != nil{
		panic(err.Error())
	}
	for _, item := range items {
		if _, err := stmt.Exec(
			item.CreateAt,
			item.VCode,
			item.IVCode,
			item.IsUnique,
			item.Campaign,
			item.SourceID,
			item.ClickPrice,
			item.IsMobil,
			item.Device,
			item.Browser,
			item.Os,
			item.Country,
			item.Region,
			item.City,
			item.Ip,
			item.Ad,
			item.Site,
			item.Sid1,
			item.Sid2,
			item.Sid3,
			item.Sid4,
			item.Sid5,
			item.Sid6,
			item.Sid7,
			item.Sid8,
			item.Sid9,
			item.Sid10,
			item.PrelandUrl,
			item.PrelandID,
			item.Session,
			item.IsTest,
			item.CountryCode,
			item.OsV,
			item.BrowserV,
			item.AffiliateId,
			item.StreamId,
			item.UserID,
		); err != nil {
			panic(err.Error())
		}
	}
	if err := tx.Commit(); err != nil {
		panic(err.Error())
	}
	stmt.Close()
	clickhouse_conn.Close()
	fmt.Println("Записали: ", len(items), " кликов")
}


// Записываем постбеки в таблицу постбеков
func writePostback(items []models.PostBack) {
	time.Sleep(time.Second)
	query := `
				INSERT INTO tracker_db.post_backs
					(vcode,
					ivcode,
					create_at,
					create_date,
					url,
					method, 
					params, 
					status_confirmed,
					status_hold,
					status_declined,
					status_other,
					status_paid,
					order_id,
					amount,
					result_message)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	clickhouse_conn := SqlxConnect()
	tx, err := clickhouse_conn.Begin()
	if err != nil {
		panic(err.Error())
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	for _, item := range items {
		if item.CreateDate.IsZero() {
			item.CreateDate = item.CreateAt
		}
		if item.StatusPaid != 0 {
			item.StatusConfirmed = 1
		}
		if _, err := stmt.Exec(
			item.VCode,
			item.IVCode,
			item.CreateAt,
			item.CreateDate,
			item.Url,
			item.Method,
			item.Params,
			item.StatusConfirmed,
			item.StatusHold,
			item.StatusDeclined,
			item.StatusOther,
			item.StatusPaid,
			item.OrderID,
			item.Amount,
			item.ResultMessage,
		); err != nil {
			panic(err.Error())
		}
	}
	if err := tx.Commit(); err != nil {
		panic(err.Error())
	}
	stmt.Close()
	clickhouse_conn.Close()
	fmt.Println("Записали ", len(items), " постбеков")
}

func WriteTraffic(tr []models.FullTraffic){
	query :=
		`INSERT INTO tracker_db.traffic_data
		(	 vcode,
			 ivcode,
 			 create_at,
   			 create_date,
			 is_click,
 	 		 is_test,
	  		 is_unique,
  			 is_mobil,
  			 is_breaked,
   			 is_refused,
			 screen_height,
			 screen_width,
			 language,
			 process_interval,
	  		 campaign,
 	 		 source_id,
 	 		 affiliate_id,
 	 		 stream_id,
 		 	 preland_url,
 	 		 preland_id,
		  	 device,
		 	 click_price,
	  		 browser,
	  		 browserv,
			 os,
			 osv,
		  	 country,
			 country_code,
  			 region,
	  	 	 city,
 	 		 ip,
 		 	 ad,
 		     site,
 			 sid1,
 			 sid2,
 			 sid3,
 		 	 sid4,
 		 	 sid5,
 		 	 sid6,
 		 	 sid7,
 		 	 sid8,
 		 	 sid9,
 		 	 sid10,
		  	 session_id,
			 uid)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	clickhouse_conn := SqlxConnect()

	tx, err := clickhouse_conn.Begin()
	if err!=nil{
		fmt.Println(err.Error())
	}

	stmt, err := tx.Prepare(query)
	if err!=nil{
		fmt.Println(err.Error())
	}

	for _, item := range tr {
		if item.CreateAt.IsZero(){
			item.CreateAt = time.Now()
		}
		if item.CreateDate.IsZero(){
			item.CreateDate = time.Now()
		}
		if _, err := stmt.Exec(
			item.VCode,
			item.IVCode,
			item.CreateAt,
			item.CreateDate,
			item.IsClick,
			item.IsTest,
			item.IsUnique,
			item.IsMobil,
			item.IsBreaked,
			item.IsRefused,
			item.ScreenHeight,
			item.ScreenWidth,
			item.Language,
			item.ProcessInterval,
			item.Campaign,
			item.SourceID,
			item.AffiliateID,
			item.StreamID,
			item.PrelandUrl,
			item.PrelandID,
			item.Device,
			item.ClickPrice,
			item.Browser,
			item.BrowserV,
			item.Os,
			item.OsV,
			item.Country,
			item.CountryCode,
			item.Region,
			item.City,
			item.Ip,
			item.Ad,
			item.Site,
			item.Sid1,
			item.Sid2,
			item.Sid3,
			item.Sid4,
			item.Sid5,
			item.Sid6,
			item.Sid7,
			item.Sid8,
			item.Sid9,
			item.Sid10,
			item.Session,
			item.UID,
		); err != nil {
			fmt.Println(err.Error())
		}
	}
	if err := tx.Commit(); err != nil {
		fmt.Println(err.Error())
	}
	stmt.Close()
	clickhouse_conn.Close()
}
