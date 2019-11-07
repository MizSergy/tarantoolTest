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
)

var bkrType = []string{
	"queue_click",
	"queue_breaking",
	"checkTime",
	"post_backs",
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading example.env file")
	}

	database.RedisInit()
	database.TarantoolInit()
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

								item.CreateAt = int(item.CreatedAt.Unix())
								traff = item.ClickToTraffic()
								database.InsertOrUpdate(conn, "clicks", item)
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
								item.CreateAt = int(item.CreatedAt.Unix())

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
	readClicks()
	readBreaks()
	readTraffic()
}

func readClicks() {
	for {
		itdc := make(map[int][]uint)
		var clks []models.Click
		for k, v := range database.TrConns {
			conn := v.Connect()
			var cls []models.Click
			cls, itdc[k] = database.GetClicksIVCodes(conn)
			clks = append(clks, cls...)
			v.Disconnect()
		}

		for k, v := range database.TrConns {
			conn := v.Connect()
			database.DropTraffList(conn, itdc[k], "clicks")
			conn.Close()
		}
	}
}

func readBreaks() {
	for {
		itdb := make(map[int][]uint)
		var brks []models.Breaking
		for k, v := range database.TrConns {
			conn := v.Connect()
			var brk []models.Breaking
			brk, itdb[k] = database.GetBreaksIVCodes(conn)
			brks = append(brks, brk...)
			conn.Close()
		}
		for k, v := range database.TrConns {
			conn := v.Connect()
			database.DropTraffList(conn, itdb[k], "breaks")
			conn.Close()
		}
	}
}

func readTraffic() {
	for {
		itdb := make(map[int][]uint)
		var trfs []models.FullTraffic
		for k, v := range database.TrConns {
			conn := v.Connect()
			var trf []models.FullTraffic
			trf, itdb[k] = database.GetTrafficIVCodes(conn)
			trfs = append(trfs, trf...)
			trfs = compareTraff(trfs)
			conn.Close()
		}
		for k, v := range database.TrConns {
			conn := v.Connect()
			database.DropTraffList(conn, itdb[k], "breaks")
			conn.Close()
		}
	}
}

func compareTraff(tr []models.FullTraffic) []models.FullTraffic{
	a := make(map[uint]models.FullTraffic)
	for _,v := range tr{
		item, ok := a[v.IVCode]
		if !ok {
			a[v.IVCode] = v
			continue
		}
		item.Compare(v)
	}

	var new []models.FullTraffic
	for _,v := range a{
		new = append(new,v)
	}
	return new
}