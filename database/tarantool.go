package database

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
	"log"
	"os"
	"tarantoolTest/models"
	"time"
)

var TrConns []TarantoolConnect
type TarantoolConnect struct {
	Host string
	Port string
	User string
	Pass string
	client *tarantool.Connection
}

func TarantoolInit(){
	index := 1
	for {
		host := os.Getenv(fmt.Sprintf("TARANTOOL_HOST_%d", index))
		if len(host) == 0 {
			break
		}
		tr := TarantoolConnect{
			Host: host,
			Port: os.Getenv(fmt.Sprintf("TARANTOOL_PORT_%d", index)),
			User: os.Getenv(fmt.Sprintf("TARANTOOL_USER_%d", index)),
			Pass: os.Getenv(fmt.Sprintf("TARANTOOL_PASS_%d", index)),
		}
		TrConns = append(TrConns, tr)
		index++
	}
}

func (conn TarantoolConnect) Connect() *tarantool.Connection{
	if conn.client == nil {
		var err error
		conn.client, err = tarantool.Connect(
			fmt.Sprintf(`%s:%s`,conn.Host,conn.Port),
			tarantool.Opts{
				Timeout:       time.Second,
				Reconnect:     time.Second,
				MaxReconnects: 5,
				User:          conn.User,
				Pass:          conn.Pass,
			})

		if err != nil {
			log.Fatal(err.Error())
		}
		return conn.client
	}
	return conn.client
}

func (conn TarantoolConnect)Disconnect(){
	conn.client.Close()
}

func InsertOrUpdate(conn *tarantool.Connection,table string,item interface{}) {
	switch table {
	case "fulltraffic":
		item = item.(models.FullTraffic)
	case "clicks":
		item = item.(models.Click)
	case "breaks":
		item = item.(models.Breaking)
	}

	_, err := conn.Call17("InsertOrUpdate", []interface{}{item, table})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func DropTraffList(conn *tarantool.Connection, list []uint, table string) {
	_, err := conn.Call17("DropTraffList", []interface{}{list, table})
	if err != nil {
		fmt.Println(err.Error())
	}
}


func GetTrafficIVCodes(conn *tarantool.Connection) ([]models.FullTraffic,[]uint) {
	var tr []models.FullTraffic
	var ivcodes []uint
	err := conn.SelectTyped("fulltraffic", "secondary", 0, 4294967295, tarantool.IterLe, []interface{}{time.Now().Add(-time.Minute).Unix()}, &tr)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _,v := range tr{
		ivcodes = append(ivcodes, v.IVCode)
	}
	return tr, ivcodes
}

func GetClicksIVCodes(conn *tarantool.Connection) ([]models.Click, []uint) {
	var tr []models.Click
	var ivcodes []uint
	err := conn.SelectTyped("clicks", "secondary", 0, 4294967295, tarantool.IterLe, []interface{}{1573116625}, &tr)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _,v := range tr{
		ivcodes = append(ivcodes, v.IVCode)
	}
	return tr, ivcodes
}

func GetBreaksIVCodes(conn *tarantool.Connection) ([]models.Breaking,[]uint){
	var tr []models.Breaking
	var ivcodes []uint
	err := conn.SelectTyped("breaks", "secondary", 0, 4294967295, tarantool.IterLe, []interface{}{time.Now().Add(-time.Minute).Unix()}, &tr)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _,v := range tr{
		ivcodes = append(ivcodes, v.IVCode)
	}
	return tr, ivcodes
}

