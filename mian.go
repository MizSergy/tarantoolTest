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
	"queue_click",
	"queue_breaking",
	"checkTime",
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.RedisInit()

	//write()

	read()
}

func write() {
	for {
		trConn := database.TarantoolConnect()
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
							database.InsertOrUpdate(trConn, "clicks", item)
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
							database.InsertOrUpdate(trConn, "breaks", item)
						}
						database.InsertOrUpdate(trConn, "fulltraffic", traff)
					}
				}
			}
		}
		database.TarantoolClose(trConn)
		time.Sleep(2 * time.Second)
	}
}

func read() {
	for {
		var keys []uint
		time.Sleep(2 * time.Minute)
		trConn := database.TarantoolConnect()
		keys = database.GetTrafficIVCodes(trConn)
		if len(keys) > 0 {
			fmt.Println("Удалили", len(keys))
			database.DropTraffList(trConn, keys, "fulltraffic")
		}

		keys = database.GetClicksIVCodes(trConn)
		if len(keys) > 0 {
			fmt.Println("Удалили", len(keys))
			database.DropTraffList(trConn, keys, "clicks")
		}

		keys = database.GetBreaksIVCodes(trConn)
		if len(keys) > 0 {
			fmt.Println("Удалили", len(keys))
			database.DropTraffList(trConn, keys, "breaks")
		}

		database.TarantoolClose(trConn)

	}
}
