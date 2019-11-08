package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"log"
	"strings"
	"tarantoolTest/database"
	"tarantoolTest/models"
	"time"
)

var bkrType = []string{
	//"queue_click",
	//"queue_breaking",
	//"checkTime",
	"queue_post_back",
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading example.env file")
	}

	database.RedisInit()
	database.TarantoolInit()
	//database.ClickhouseInit()
	write()
	//read()
}

func write() {
	for {
		for _, client := range database.ListRedisConnect {
			if client.Connect() == nil {
				fmt.Printf("Error connect %s\n", client.Host)
				continue
			}

			for _, t := range bkrType {
				items := client.GetAllCollections(t)
				itCount := len(items)
				if itCount > 0 {
					fmt.Println("Получили ", itCount, t)
					for _, v := range database.TrConns {
						conn := v.Connect()
						for _, v := range items {
							var traff models.FullTraffic
							switch t {
							case "queue_click":
								item := models.Click{}
								err := json.Unmarshal([]byte(v), &item)
								if err != nil {
									fmt.Println(err.Error())
								}
								if len(item.VCode) > 36 {
									item.VCode = strings.Replace(item.VCode, " ", "", -1)[:36]
								}

								uid, err := uuid.Parse(item.VCode)
								if err != nil {
									fmt.Println(err.Error())
								}
								item.IVCode = uint(uid.ID())

								item.CreatedAt = int(item.CreateAt.Unix())
								traff = item.ClickToTraffic()
								database.InsertOrUpdate(conn, "clicks", item)
							case "queue_post_back":
								item := models.PostBack{}
								err := json.Unmarshal([]byte(v), &item)
								if err != nil {
									fmt.Println(err.Error())
								}
								if len(item.VCode) > 36 {
									item.VCode = strings.Replace(item.VCode, " ", "", -1)[:36]
								}

								uid, err := uuid.Parse(item.VCode)
								if err != nil {
									fmt.Println(err.Error())
								}
								item.IVCode = uint(uid.ID())

								item.CreatedAt = int(item.CreateAt.Unix())
								if item.CreateDate.IsZero() && item.CreatedDate == 0{
									item.CreatedDate = int(item.CreateAt.Unix())
								}
								database.InsertOrUpdate(conn, "postbacks", item)
							default:
								item := models.Breaking{}
								err := json.Unmarshal([]byte(v), &item)
								if err != nil {
									fmt.Println(err.Error())
								}
								if len(item.VCode) > 36 {
									item.VCode = strings.Replace(item.VCode, " ", "", -1)[:36]
								}

								uid, err := uuid.Parse(item.VCode)
								if err != nil {
									fmt.Println(err.Error())
								}
								item.IVCode = uint(uid.ID())
								item.CreatedAt = int(item.CreateAt.Unix())

								traff = item.BreaksToTraffic()
								database.InsertOrUpdate(conn, "breaks", item)
							}
							database.InsertOrUpdate(conn, "fulltraffic", traff)
						}
						conn.Close()
					}
				}
			}
		}
	}
}

func read() {
	for {
		readClicks()
		fmt.Println("Клики")
		readBreaks()
		fmt.Println("Пробивы")
		readTraffic()
		fmt.Println("Трафик")
		time.Sleep(6 * time.Minute )
	}

}

func readClicks() {
	itdc := make(map[int][]uint)
	var clks []models.Click
	for k, v := range database.TrConns {
		conn := v.Connect()
		var cls []models.Click

		cls, itdc[k] = database.GetClicksIVCodes(conn)
		conn.Close()
		if len(cls) > 0 {
			clks = append(clks, cls...)
		}

		fmt.Println("Забрали клики с ", k)
	}
	for k, v := range database.TrConns {
		if len(itdc[k]) > 0 {
			fmt.Println("Удалили ", len(itdc[k]))
			conn := v.Connect()
			database.DropTraffList(conn, itdc[k], "clicks")
			conn.Close()
		}
	}
}

func readBreaks() {
	itdb := make(map[int][]uint)
	var brks []models.Breaking
	for k, v := range database.TrConns {
		conn := v.Connect()
		var brk []models.Breaking
		brk, itdb[k] = database.GetBreaksIVCodes(conn)
		conn.Close()
		if len(brk)>0{
			brks = append(brks, brk...)
		}
		fmt.Println("Забрали пробивы с ", k)
	}
	for k, v := range database.TrConns {
		if len(itdb[k]) > 0 {
			fmt.Println("Удалили ", len(itdb[k]))
			conn := v.Connect()
			database.DropTraffList(conn, itdb[k], "breaks")
			conn.Close()
		}
	}
}

func readPostBacks() {
	itdb := make(map[int][]uint)
	var pbs []models.PostBack
	for k, v := range database.TrConns {
		conn := v.Connect()
		var pb []models.PostBack
		pb, itdb[k] = database.GetPostBacksIVCodes(conn)
		conn.Close()
		if len(pb)>0{
			pbs = append(pbs, pb...)
		}
		fmt.Println("Забрали пробивы с ", k)
	}
	for k, v := range database.TrConns {
		if len(itdb[k]) > 0 {
			fmt.Println("Удалили ", len(itdb[k]))
			conn := v.Connect()
			database.DropTraffList(conn, itdb[k], "breaks")
			conn.Close()
		}
	}
}

func readTraffic() {
	itdb := make(map[int][]uint)
	var trfs []models.FullTraffic
	for k, v := range database.TrConns {
		conn := v.Connect()
		var trf []models.FullTraffic
		trf, itdb[k] = database.GetTrafficIVCodes(conn)
		conn.Close()
		if len(trf) > 0{
			trfs = append(trfs, trf...)
			trfs = compareTraff(trfs)
		}
		fmt.Println("Забрали трафик с ", k)
	}
	for k, v := range database.TrConns {
		if len(itdb[k]) > 0 {
			fmt.Println("Удалили ", len(itdb[k]))
			conn := v.Connect()
			database.DropTraffList(conn, itdb[k], "fulltraffic")
			conn.Close()
		}

	}
}

func compareTraff(tr []models.FullTraffic) []models.FullTraffic {
	a := make(map[uint]models.FullTraffic)
	for _, v := range tr {
		item, ok := a[v.IVCode]
		if !ok {
			a[v.IVCode] = v
			continue
		}
		item.Compare(&v)
	}

	var new []models.FullTraffic
	for _, v := range a {
		new = append(new, v)
	}
	return new
}
