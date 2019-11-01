package database

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
	"github.com/tarantool/go-tarantool/queue"
	"log"
	"tarantoolTest/models"
	"time"
)

var opts = tarantool.Opts{
	Timeout:       time.Second,
	Reconnect:     time.Second,
	MaxReconnects: 5,
	User:          "tracker",
	Pass:          "dskjfhgkjdsfhg",
}

var cfg = queue.Cfg{
	Temporary:   true,
	IfNotExists: true,
	Kind:        queue.FIFO,
	Opts: queue.Opts{
		Ttl:   10 * time.Second,
		Ttr:   5 * time.Second,
		Delay: 3 * time.Second,
		Pri:   1,
	},
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

func TarantoolConnect() *tarantool.Connection {
	conn, err := tarantool.Connect("89.208.35.145:3301", opts)
	if err != nil {
		log.Fatalf("connection: %s", err)
	}
	return conn
}



func GetTrafficIVCodes(conn *tarantool.Connection) []uint {
	var tr []models.FullTraffic
	var ivcodes []uint
	err := conn.SelectTyped("fulltraffic", "secondary", 0, 4294967295, tarantool.IterLe, []interface{}{time.Now().Add(-time.Minute).Unix()}, &tr)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _,v := range tr{
		ivcodes = append(ivcodes, v.IVCode)
	}
	return ivcodes
}

func GetClicksIVCodes(conn *tarantool.Connection) []uint {
	var tr []models.Click
	var ivcodes []uint
	err := conn.SelectTyped("clicks", "secondary", 0, 4294967295, tarantool.IterLe, []interface{}{time.Now().Add(-time.Minute).Unix()}, &tr)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _,v := range tr{
		ivcodes = append(ivcodes, v.IVCode)
	}
	return ivcodes
}

func GetBreaksIVCodes(conn *tarantool.Connection) []uint {
	var tr []models.Breaking
	var ivcodes []uint
	err := conn.SelectTyped("breaks", "secondary", 0, 4294967295, tarantool.IterLe, []interface{}{time.Now().Add(-time.Minute).Unix()}, &tr)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _,v := range tr{
		ivcodes = append(ivcodes, v.IVCode)
	}
	return ivcodes
}


func TarantoolClose(conn *tarantool.Connection) {
	conn.Close()
}

func CreateOrPushToQueue(conn *tarantool.Connection, key string, item interface{}) {
	que := queue.New(conn, key)
	if err := que.Create(cfg); err != nil {
		log.Fatalf("queue create: %s", err)
		return
	}

	task, err := que.Put(&item)
	if err != nil {
		log.Fatalf("put typed task: %s", err)
	}
	fmt.Println("Task id is ", task.Id())
}

func TakeFromQueue(conn *tarantool.Connection, key string) {
	var item interface{}
	switch key {
	case "queue_click":
		item = item.(models.Click)
	case "queue_breaking", "checkTime":
		item = item.(models.Breaking)
	}

	que := queue.New(conn, key)
	if err := que.Create(cfg); err != nil {
		log.Fatalf("queue create: %s", err)
		return
	}

	task, err := que.TakeTyped(&item)
	if err != nil {
		log.Fatalf("put typed task: %s", err)
	}
	fmt.Println("Task id is ", task.Id())

}
